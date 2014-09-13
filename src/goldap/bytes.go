package goldap

import (
	"errors"
	"fmt"
)

type Bytes struct {
	offset *int
	bytes  []byte
}

func NewBytes(offset int, bytes []byte) (ret Bytes){
	return Bytes{offset: &offset, bytes: bytes}
}

func (b Bytes) Debug() {
	fmt.Println(fmt.Sprintf("Offset: %d, Bytes: %+v", *b.offset, b.bytes))
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

func (b Bytes) ParseSubBytes(class int, tag int, callback func(bytes Bytes) error) (err error) {
	// Check tag
	tagAndLength, err := b.ParseTagAndLength()
	if err != nil {
		return errors.New(fmt.Sprintf("ParseSequence: %s", err.Error()))
	}
	err = tagAndLength.Expect(class, tag, isCompound)
	if err != nil {
		return errors.New(fmt.Sprintf("ParseSequence: %s", err.Error()))
	}

	start := *b.offset
	end := *b.offset + tagAndLength.GetLength()

	// Check we got enough bytes to process
	if end > len(b.bytes) {
		return StructuralError{fmt.Sprintf("ParseSequence : DATA TRUNCATED: expecting %d bytes at offset %d", tagAndLength.GetLength(), b.offset)}
	}
	// Process sub-bytes
	zero := 0
	subBytes := Bytes{offset: &zero, bytes: b.bytes[start:end]}
	err = callback(subBytes)
	if err != nil {
		err = errors.New(fmt.Sprintf("ParseSequence: %s", err.Error()))
		*b.offset += *subBytes.offset
		return
	}
	// Check we got no more bytes to process
	if subBytes.HasMoreData() {
		return StructuralError{fmt.Sprintf("ParseSequence: DATA TOO LONG: %d more bytes to read at offset %d", end-*b.offset, *b.offset)}
	}
	// Move offset
	*b.offset = end
	return
}

func (b Bytes) HasMoreData() bool {
	return *b.offset < len(b.bytes)
}

func (b Bytes) PreviewTagAndLength() (tagAndLength tagAndLength, err error) {
	previousOffset := *b.offset // Save offset
	tagAndLength, err = b.ParseTagAndLength()
	*b.offset = previousOffset // Restore offset
	return
}

func (b Bytes) ParseTagAndLength() (ret tagAndLength, err error) {
	ret, *b.offset = ParseTagAndLength(b.bytes, *b.offset)
	return
}

// The parse"Type" functions are use the underlaying asn1 functions
// They are ony here to increase the offset of the current Bytes objet
func (b Bytes) ParseBool(length int) (ret bool, err error) {
	ret, err = parseBool(b.bytes[*b.offset : *b.offset+length])
	if err != nil {
		return
	}
	*b.offset += length
	return
}

func (b Bytes) ParseInt32(length int) (ret int32, err error) {
	ret, err = parseInt32(b.bytes[*b.offset : *b.offset+length])
	if err != nil {
		return
	}
	*b.offset += length
	return
}

func (b Bytes) ParseUTF8String(length int) (utf8string string, err error) {
	utf8string, err = parseUTF8String(b.bytes[*b.offset : *b.offset+length])
	if err != nil {
		return
	}
	*b.offset += length
	return
}

func (b Bytes) ParseOCTETSTRING(length int) (ret []byte, err error) {
	if *b.offset+length > len(b.bytes) {
		err = StructuralError{"Data truncated"}
		return
	}
	ret = b.bytes[*b.offset : *b.offset+length]
	*b.offset += length
	return
}
