package goldap

import (
	"goldap/asn1"
	"fmt"
)

type Bytes struct {
	bytes  []byte
	offset int
}
func (b *Bytes) ParseSequence(class int, tag int, callback func(bytes *Bytes)) (err error) {
	// Check tag
	tagAndLength := b.ParseTagAndLength()
	tagAndLength.ExpectClass(class)
	tagAndLength.ExpectTag(tag)
	tagAndLength.ExpectCompound(asn1.IsCompound)

	start := b.offset
	end := b.offset + tagAndLength.GetLength()

	// Check we got enough bytes to process
	if end > len(b.bytes) {
		return StructuralError{fmt.Sprintf("DATA TRUNCATED: expecting %d bytes at offset %d", tagAndLength.GetLength(), b.offset)}
	}
	// Process sub-bytes
	subBytes := &Bytes{offset: 0, bytes: b.bytes[start:end]}
	callback(subBytes)
	// Check we got no more bytes to process
	if subBytes.HasMoreData() {
		return StructuralError{fmt.Sprintf("DATA TOO LONG: %d more bytes to read at offset %d", end-b.offset, b.offset)}
	}
	return
}

func (b *Bytes) HasMoreData() bool {
	return b.offset < len(b.bytes)
}

func (b *Bytes) PreviewTagAndLength() (tagAndLength asn1.TagAndLength) {
	previousOffset := b.offset // Save offset
	tagAndLength = b.ParseTagAndLength() 
	b.offset = previousOffset // Restore offset 
	return
}

func (b *Bytes) ParseTagAndLength() asn1.TagAndLength {
	tagAndLength, offset := asn1.ParseTagAndLength(b.bytes, b.offset)
	b.offset = offset
	return tagAndLength
}

func (b *Bytes) ParseInt32(length int) (value int32) {
	value = asn1.ParseInt32(b.bytes, b.offset, length)
	b.offset += length
	return
}

func (b *Bytes) ParseUTF8STRING(length int) (utf8string UTF8STRING){
	utf8string = UTF8STRING(asn1.ParseUTF8String(b.bytes, b.offset, length))
	b.offset += length
	return
}

func (b *Bytes) ParseOCTETSTRING(length int) (octetstring OCTETSTRING){
	octetstring = OCTETSTRING(b.bytes[b.offset : b.offset + length])
	b.offset += length
	return
}
