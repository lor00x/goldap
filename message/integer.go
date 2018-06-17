package message

import "fmt"

func readINTEGER(bytes *Bytes) (ret INTEGER, err error) {
	var value interface{}
	value, err = bytes.ReadPrimitiveSubBytes(classUniversal, tagInteger, tagInteger)
	if err != nil {
		err = LdapError{fmt.Sprintf("readINTEGER:\n%s", err.Error())}
		return
	}
	ret = INTEGER(value.(int32))
	return
}
func readTaggedINTEGER(bytes *Bytes, class int, tag int) (ret INTEGER, err error) {
	var value interface{}
	value, err = bytes.ReadPrimitiveSubBytes(class, tag, tagInteger)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedINTEGER:\n%s", err.Error())}
		return
	}
	ret = INTEGER(value.(int32))
	return
}
func readTaggedPositiveINTEGER(bytes *Bytes, class int, tag int) (ret INTEGER, err error) {
	ret, err = readTaggedINTEGER(bytes, class, tag)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedPositiveINTEGER:\n%s", err.Error())}
		return
	}
	if !(ret >= 0 && ret <= maxInt) {
		err = LdapError{fmt.Sprintf("readTaggedPositiveINTEGER: Invalid INTEGER value %d ! Expected value between 0 and %d", ret, maxInt)}
	}
	return
}
func readPositiveINTEGER(bytes *Bytes) (ret INTEGER, err error) {
	return readTaggedPositiveINTEGER(bytes, classUniversal, tagInteger)
}