package goldap

import (
	"fmt"
)

type Bytes struct {
	offset int
	bytes  []byte
}

func (b *Bytes) ParseSequence(class int, tag int, callback func(bytes *Bytes) error) (err error) {
	// Check tag
	tagAndLength, err := b.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(class, tag, isCompound)
	if err != nil {
		return
	}

	start := b.offset
	end := b.offset + tagAndLength.GetLength()

	// Check we got enough bytes to process
	if end > len(b.bytes) {
		return StructuralError{fmt.Sprintf("DATA TRUNCATED: expecting %d bytes at offset %d", tagAndLength.GetLength(), b.offset)}
	}
	// Process sub-bytes
	subBytes := &Bytes{offset: 0, bytes: b.bytes[start:end]}
	err = callback(subBytes)
	if err != nil {
		return
	}
	// Check we got no more bytes to process
	if subBytes.HasMoreData() {
		return StructuralError{fmt.Sprintf("DATA TOO LONG: %d more bytes to read at offset %d", end-b.offset, b.offset)}
	}
	// Move offset
	b.offset = end
	return
}

func (b *Bytes) HasMoreData() bool {
	return b.offset < len(b.bytes)
}

func (b *Bytes) PreviewTagAndLength() (tagAndLength tagAndLength, err error) {
	previousOffset := b.offset // Save offset
	tagAndLength, err = b.ParseTagAndLength()
	b.offset = previousOffset // Restore offset
	return
}

func (b *Bytes) ParseTagAndLength() (ret tagAndLength, err error) {
	ret, offset := ParseTagAndLength(b.bytes, b.offset)
	b.offset = offset
	return
}

// The parse"Type" functions are use the underlaying asn1 functions
// They are ony here to increase the offset of the current Bytes object
func (b *Bytes) ParseBool(length int) (ret bool, err error) {
	ret, err = parseBool(b.bytes[b.offset : b.offset+length])
	if err != nil {
		return
	}
	b.offset += length
	return
}

func (b *Bytes) ParseInt32(length int) (ret int32, err error) {
	ret, err = parseInt32( b.bytes[b.offset : b.offset+length] )
	if err != nil {
		return
	}
	b.offset += length
	return
}

func (b *Bytes) ParseUTF8String(length int) (utf8string string, err error) {
	utf8string, err = parseUTF8String( b.bytes[b.offset : b.offset+length] )
	if err != nil {
		return
	}
	b.offset += length
	return
}

func (b *Bytes) ParseOCTETSTRING(length int) (ret []byte, err error) {
	if b.offset+length > len(b.bytes) {
		err = StructuralError{"Data truncated"}
		return
	}
	ret  = b.bytes[b.offset : b.offset+length]
	b.offset += length
	return
}

