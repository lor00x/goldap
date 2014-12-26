package message

import (
	"fmt"
)

type Bytes struct {
	offset *int
	bytes  []byte
}

func NewBytes(offset int, bytes []byte) (ret Bytes) {
	return Bytes{offset: &offset, bytes: bytes}
}

func (b Bytes) Debug() {
	fmt.Printf("Offset: %d, Bytes: %+v\n", *b.offset, b.bytes)
}

// Return a string with the hex dump of the bytes around the current offset
// The current offset byte is put in brackets
// Example: 0x01, [0x02], 0x03
func (b Bytes) DumpCurrentBytes() (ret string) {
	var strings [3]string
	for i := -1; i <= 1; i++ {
		if *b.offset+i >= 0 && *b.offset+i < len(b.bytes) {
			strings[i+1] = fmt.Sprintf("%#x", b.bytes[*b.offset+i])
		}
	}
	ret = fmt.Sprintf("%s, [%s], %s", strings[0], strings[1], strings[2])
	return
}

func (b Bytes) HasMoreData() bool {
	return *b.offset < len(b.bytes)
}

func (b Bytes) PreviewTagAndLength() (tagAndLength TagAndLength, err error) {
	previousOffset := *b.offset // Save offset
	tagAndLength, err = b.ParseTagAndLength()
	*b.offset = previousOffset // Restore offset
	return
}

func (b Bytes) ParseTagAndLength() (ret TagAndLength, err error) {
	var offset int
	ret, offset, err = ParseTagAndLength(b.bytes, *b.offset)
	if err != nil {
		err = LdapError{fmt.Sprintf("ParseTagAndLength: %s", err.Error())}
		return
	} else {
		*b.offset = offset
	}
	return
}

func (b Bytes) ReadSubBytes(class int, tag int, callback func(bytes Bytes) error) (err error) {
	// Check tag
	tagAndLength, err := b.ParseTagAndLength()
	if err != nil {
		return LdapError{fmt.Sprintf("ReadSubBytes:\n%s", err.Error())}
	}
	err = tagAndLength.Expect(class, tag, isCompound)
	if err != nil {
		return LdapError{fmt.Sprintf("ReadSubBytes:\n%s", err.Error())}
	}

	start := *b.offset
	end := *b.offset + tagAndLength.Length

	// Check we got enough bytes to process
	if end > len(b.bytes) {
		return LdapError{fmt.Sprintf("ReadSubBytes: data truncated: expecting %d bytes at offset %d", tagAndLength.Length, *b.offset)}
	}
	// Process sub-bytes
	zero := 0
	subBytes := Bytes{offset: &zero, bytes: b.bytes[start:end]}
	err = callback(subBytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("ReadSubBytes:\n%s", err.Error())}
		*b.offset += *subBytes.offset
		return
	}
	// Check we got no more bytes to process
	if subBytes.HasMoreData() {
		return LdapError{fmt.Sprintf("ReadSubBytes: data too long: %d more bytes to read at offset %d", end-*b.offset, *b.offset)}
	}
	// Move offset
	*b.offset = end
	return
}

func SizeSubBytes(tag int, callback func() int) (size int) {
	size = callback()
	size += sizeTagAndLength(tag, size)
	return
}

//
// Parse tag, length and read the a primitive value
// Supported types are:
// - boolean
// - integer (parsed as int32)
// - enumerated (parsed as int32)
// - UTF8 string
// - Octet string
//
// Parameters:
// - class: the expected class value(classUniversal, classApplication, classContextSpecific)
// - tag: the expected tag value
// - typeTag: the real primitive type to parse (tagBoolean, tagInteger, tagEnym, tagUTF8String, tagOctetString)
//
func (b Bytes) ReadPrimitiveSubBytes(class int, tag int, typeTag int) (value interface{}, err error) {
	// Check tag
	tagAndLength, err := b.ParseTagAndLength()
	if err != nil {
		err = LdapError{fmt.Sprintf("ReadPrimitiveSubBytes:\n%s", err.Error())}
		return
	}
	err = tagAndLength.Expect(class, tag, isNotCompound)
	if err != nil {
		err = LdapError{fmt.Sprintf("ReadPrimitiveSubBytes:\n%s", err.Error())}
		return
	}

	start := *b.offset
	end := *b.offset + tagAndLength.Length

	// Check we got enough bytes to process
	if end > len(b.bytes) {
		// err = LdapError{fmt.Sprintf("ReadPrimitiveSubBytes: data truncated: expecting %d bytes at offset %d but only %d bytes are remaining (start: %d, length: %d, end: %d, len(b): %d, bytes: %#+v)", tagAndLength.Length, *b.offset, len(b.bytes)-start, start, tagAndLength.Length, end, len(b.bytes), b.bytes)}
		err = LdapError{fmt.Sprintf("ReadPrimitiveSubBytes: data truncated: expecting %d bytes at offset %d but only %d bytes are remaining", tagAndLength.Length, *b.offset, len(b.bytes)-start)}
		return
	}
	// Process sub-bytes
	subBytes := b.bytes[start:end]
	switch typeTag {
	case tagBoolean:
		value, err = parseBool(subBytes)
	case tagInteger:
		value, err = parseInt32(subBytes)
	case tagEnum:
		value, err = parseInt32(subBytes)
	case tagUTF8String:
		value, err = parseUTF8String(subBytes)
	case tagOctetString:
		value, err = parseOctetString(subBytes)
	default:
		err = LdapError{fmt.Sprintf("ReadPrimitiveSubBytes: invalid type tag value %d", typeTag)}
		return
	}
	if err != nil {
		err = LdapError{fmt.Sprintf("ReadPrimitiveSubBytes:\n%s", err.Error())}
		return
	}
	// Move offset
	*b.offset = end
	return
}

func SizePrimitiveSubBytes(tag int, typeTag int, value interface{}) (size int) {
	switch typeTag {
	case tagBoolean:
		size = sizeBool(value.(bool))
	case tagInteger:
		size = sizeInt32(value.(int32))
	case tagEnum:
		size = sizeInt32(value.(int32))
	case tagUTF8String:
		size = sizeUTF8String(value.(string))
	case tagOctetString:
		size = sizeOctetString([]byte(value.(string)))
	default:
		panic(fmt.Sprintf("SizePrimitiveSubBytes: invalid type tag value %d", typeTag))
	}
	size += sizeTagAndLength(tag, size)
	return
}

// func (b Bytes) Writedefault:
// 	// Check tag
// 	tagAndLength, err := b.ParseTagAndLength()
// 	if err != nil {
// 		return errors.New(fmt.Sprintf("ParseSequence: %s", err.Error()))
// 	}
// 	err = tagAndLength.Expect(class, tag, isCompound)
// 	if err != nil {
// 		return errors.New(fmt.Sprintf("ParseSequence: %s", err.Error()))
// 	}

// 	start := *b.offset
// 	end := *b.offset + tagAndLength.Length

// 	// Check we got enough bytes to process
// 	if end > len(b.bytes) {
// 		return StructuralError{fmt.Sprintf("ParseSequence : DATA TRUNCATED: expecting %d bytes at offset %d", tagAndLength.Length, b.offset)}
// 	}
// 	// Process sub-bytes
// 	zero := 0
// 	subBytes := Bytes{offset: &zero, bytes: b.bytes[start:end]}
// 	err = callback(subBytes)
// 	if err != nil {
// 		err = errors.New(fmt.Sprintf("ParseSequence: %s", err.Error()))
// 		*b.offset += *subBytes.offset
// 		return
// 	}
// 	// Check we got no more bytes to process
// 	if subBytes.HasMoreData() {
// 		return StructuralError{fmt.Sprintf("ParseSequence: DATA TOO LONG: %d more bytes to read at offset %d", end-*b.offset, *b.offset)}
// 	}
// 	// Move offset
// 	*b.offset = end
// 	return
// }
