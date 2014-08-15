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

func (t *TagAndLength) Expect(class int, tag int, isCompound bool) (err error) {
	err = t.ExpectClass(class)
	if err != nil {
		return
	}
	err = t.ExpectTag(tag)
	if err != nil {
		return
	}
	err = t.ExpectCompound(isCompound)
	return
}
func (t *TagAndLength) ExpectClass(class int) (err error) {
	if class != t.class {
		err = errors.New(fmt.Sprintf("Wrong tag class %d. Expected %d.", t.class, class))
	}
	return
}
func (t *TagAndLength) ExpectTag(tag int) (err error) {
	if tag != t.tag {
		err = errors.New(fmt.Sprintf("Wrong tag value %d. Expected %d.", t.tag, tag))
	}
	return
}
func (t *TagAndLength) ExpectCompound(isCompound bool) (err error) {
	if isCompound != t.isCompound {
		err = errors.New(fmt.Sprintf("Wrong tag compound %t. Expected %d.", t.isCompound, isCompound))
	}
	return
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
