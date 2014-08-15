package goldap

import (
	"goldap/asn1"
	"fmt"
)

//const (
//	TagBoolean         = 1
//	TagInteger         = 2
//	TagBitString       = 3
//	TagOctetString     = 4
//	TagOID             = 6
//	TagEnum            = 10
//	TagUTF8String      = 12
//	TagSequence        = 16
//	TagSet             = 17
//	TagPrintableString = 19
//	TagT61String       = 20
//	TagIA5String       = 22
//	TagUTCTime         = 23
//	TagGeneralizedTime = 24
//	TagGeneralString   = 27
//)
//
//const (
//	ClassUniversal       = 0
//	ClassApplication     = 1
//	ClassContextSpecific = 2
//	ClassPrivate         = 3
//)
//
//type TagAndLength struct {
//	class, tag, length int
//	isCompound         bool
//}


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
	err = tagAndLength.Expect(class, tag, asn1.IsCompound)
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

func (b *Bytes) PreviewTagAndLength() (tagAndLength asn1.TagAndLength, err error) {
	previousOffset := b.offset // Save offset
	tagAndLength, err = b.ParseTagAndLength() 
	b.offset = previousOffset // Restore offset 
	return
}

func (b *Bytes) ParseTagAndLength() (ret asn1.TagAndLength, err error) {
	ret, offset := asn1.ParseTagAndLength(b.bytes, b.offset)
	b.offset = offset
	return
}

func (b *Bytes) ParseInt32(length int) (value int32, err error) {
	value = asn1.ParseInt32(b.bytes, b.offset, length)
	b.offset += length
	return
}

func (b *Bytes) ParseUTF8STRING(length int) (utf8string UTF8STRING, err error){
	utf8string = UTF8STRING(asn1.ParseUTF8String(b.bytes, b.offset, length))
	b.offset += length
	return
}

func (b *Bytes) ParseOCTETSTRING(length int) (octetstring OCTETSTRING, err error){
	if b.offset + length > len(b.bytes) {
		err = StructuralError{"Data truncated"}
	} else {
		octetstring = OCTETSTRING(b.bytes[b.offset : b.offset + length])
		b.offset += length
	}
	return
}

func (b *Bytes) ParseBoolean(length int) (ret bool, err error){
	return
}