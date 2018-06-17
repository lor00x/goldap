package message

import "fmt"

//
//        AssertionValue ::= OCTET STRING
func readAssertionValue(bytes *Bytes) (assertionvalue AssertionValue, err error) {
	var octetstring OCTETSTRING
	octetstring, err = readOCTETSTRING(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAssertionValue:\n%s", err.Error())}
		return
	}
	assertionvalue = AssertionValue(octetstring)
	return
}
func readTaggedAssertionValue(bytes *Bytes, class int, tag int) (assertionvalue AssertionValue, err error) {
	var octetstring OCTETSTRING
	octetstring, err = readTaggedOCTETSTRING(bytes, class, tag)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedAssertionValue:\n%s", err.Error())}
		return
	}
	assertionvalue = AssertionValue(octetstring)
	return
}

//
//        AssertionValue ::= OCTET STRING
func (a AssertionValue) write(bytes *Bytes) int {
	return OCTETSTRING(a).write(bytes)
}
func (a AssertionValue) writeTagged(bytes *Bytes, class int, tag int) int {
	return OCTETSTRING(a).writeTagged(bytes, class, tag)
}

//
//        AttributeValue ::= OCTET STRING
func (a AttributeValue) size() int {
	return OCTETSTRING(a).size()
}

//
//        AssertionValue ::= OCTET STRING
func (a AssertionValue) size() int {
	return OCTETSTRING(a).size()
}
func (a AssertionValue) sizeTagged(tag int) int {
	return OCTETSTRING(a).sizeTagged(tag)
}
