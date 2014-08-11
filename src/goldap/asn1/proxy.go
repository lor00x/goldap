package asn1

import (
	"errors"
	"fmt"
)

const (
	TagBoolean         = tagBoolean
	TagInteger         = tagInteger
	TagBitString       = tagBitString
	TagOctetString     = tagOctetString
	TagOID             = tagOID
	TagEnum            = tagEnum
	TagUTF8String      = tagUTF8String
	TagSequence        = tagSequence
	TagSet             = tagSet
	TagPrintableString = tagPrintableString
	TagT61String       = tagT61String
	TagIA5String       = tagIA5String
	TagUTCTime         = tagUTCTime
	TagGeneralizedTime = tagGeneralizedTime
	TagGeneralString   = tagGeneralString
)

const (
	ClassUniversal       = classUniversal
	ClassApplication     = classApplication
	ClassContextSpecific = classContextSpecific
	ClassPrivate         = classPrivate
)

const (
	IsCompound    = true
	IsNotCompound = false
)

type TagAndLength tagAndLength

func (t *TagAndLength) GetClass() int {
	return t.class
}
func (t *TagAndLength) GetTag() int {
	return t.tag
}
func (t *TagAndLength) GetLength() int {
	return t.length
}

func (t *TagAndLength) Expect(class int, tag int, isCompound bool) {
	t.ExpectClass(class)
	t.ExpectTag(tag)
	t.ExpectCompound(isCompound)
}
func (t *TagAndLength) ExpectClass(class int) {
	if class != t.class {
		errorMsg := fmt.Sprintf("UNEXPECTED TAG CLASS ! WANT class %d. GOT class %d.", class, t.class)
		panic(errors.New(errorMsg))
	}
}
func (t *TagAndLength) ExpectTag(tag int) {
	if tag != t.tag {
		errorMsg := fmt.Sprintf("UNEXPECTED TAG VALUE! WANT tag %d. GOT tag %d.", tag, t.tag)
		panic(errors.New(errorMsg))
	}
}
func (t *TagAndLength) ExpectCompound(isCompound bool) {
	if isCompound != t.isCompound {
		errorMsg := fmt.Sprintf("UNEXPECTED TAG COMPOUND ! WANT compound %d. GOT compound %d.", isCompound, t.isCompound)
		panic(errors.New(errorMsg))
	}
}

func ParseTagAndLength(bytes []byte, initOffset int) (ret TagAndLength, offset int) {
	tag, offset, err := parseTagAndLength(bytes, initOffset)
	if err != nil {
		panic(err)
	}
	ret = TagAndLength(tag)
	return
}

func ParseInt32(bytes []byte, offset int, length int) int32 {
	ret, err := parseInt32(bytes[offset : offset+length])
	if err != nil {
		panic(err)
	}
	return ret
}

func ParseUTF8String(bytes []byte, offset int, length int) string {
	ret, err := parseUTF8String(bytes[offset : offset+length])
	if err != nil {
		panic(err)
	}
	return ret
}
