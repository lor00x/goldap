package message

import "fmt"

func readOCTETSTRING(bytes *Bytes) (ret OCTETSTRING, err error) {
	var value interface{}
	value, err = bytes.ReadPrimitiveSubBytes(classUniversal, tagOctetString, tagOctetString)
	if err != nil {
		err = LdapError{fmt.Sprintf("readOCTETSTRING:\n%s", err.Error())}
		return
	}
	ret = OCTETSTRING(value.([]byte))
	return
}

func readTaggedOCTETSTRING(bytes *Bytes, class int, tag int) (ret OCTETSTRING, err error) {
	var value interface{}
	value, err = bytes.ReadPrimitiveSubBytes(class, tag, tagOctetString)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedOCTETSTRING:\n%s", err.Error())}
		return
	}
	ret = OCTETSTRING(value.([]byte))
	return
}
func (o OCTETSTRING) Pointer() *OCTETSTRING { return &o }
