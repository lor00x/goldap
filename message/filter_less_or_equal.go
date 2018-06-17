package message

import "fmt"

//             lessOrEqual     [6] AttributeValueAssertion,
func readFilterLessOrEqual(bytes *Bytes) (ret FilterLessOrEqual, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = readTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterLessOrEqual)
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilterLessOrEqual:\n%s", err.Error())}
		return
	}
	ret = FilterLessOrEqual(attributevalueassertion)
	return
}
