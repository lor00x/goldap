package message

import "fmt"

func readBOOLEAN(bytes *Bytes) (ret BOOLEAN, err error) {
	var value interface{}
	value, err = bytes.ReadPrimitiveSubBytes(classUniversal, tagBoolean, tagBoolean)
	if err != nil {
		err = LdapError{fmt.Sprintf("readBOOLEAN:\n%s", err.Error())}
		return
	}
	ret = BOOLEAN(value.(bool))
	return
}
func (b BOOLEAN) write(bytes *Bytes) int {
	return bytes.WritePrimitiveSubBytes(classUniversal, tagBoolean, b)
}
func (b BOOLEAN) writeTagged(bytes *Bytes, class int, tag int) int {
	return bytes.WritePrimitiveSubBytes(class, tag, b)
}
func readTaggedBOOLEAN(bytes *Bytes, class int, tag int) (ret BOOLEAN, err error) {
	var value interface{}
	value, err = bytes.ReadPrimitiveSubBytes(class, tag, tagBoolean)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedBOOLEAN:\n%s", err.Error())}
		return
	}
	ret = BOOLEAN(value.(bool))
	return
}
