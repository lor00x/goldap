package message

import "fmt"

//             present         [7] AttributeDescription,
func readFilterPresent(bytes *Bytes) (ret FilterPresent, err error) {
	var attributedescription AttributeDescription
	attributedescription, err = readTaggedAttributeDescription(bytes, classContextSpecific, TagFilterPresent)
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilterPresent:\n%s", err.Error())}
		return
	}
	ret = FilterPresent(attributedescription)
	return
}
