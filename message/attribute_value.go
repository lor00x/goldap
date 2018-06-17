package message

import "fmt"

//
//        AttributeValue ::= OCTET STRING
func readAttributeValue(bytes *Bytes) (ret AttributeValue, err error) {
	var octetstring OCTETSTRING
	octetstring, err = readOCTETSTRING(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAttributeValue:\n%s", err.Error())}
		return
	}
	ret = AttributeValue(octetstring)
	return
}

//
//        AttributeValueAssertion ::= SEQUENCE {
//             attributeDesc   AttributeDescription,
//             assertionValue  AssertionValue }
func readAttributeValueAssertion(bytes *Bytes) (ret AttributeValueAssertion, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAttributeValueAssertion:\n%s", err.Error())}
		return
	}
	return

}
func readTaggedAttributeValueAssertion(bytes *Bytes, class int, tag int) (ret AttributeValueAssertion, err error) {
	err = bytes.ReadSubBytes(class, tag, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedAttributeValueAssertion:\n%s", err.Error())}
		return
	}
	return
}

//
//        AttributeValue ::= OCTET STRING
func (a AttributeValue) write(bytes *Bytes) int {
	return OCTETSTRING(a).write(bytes)
}
