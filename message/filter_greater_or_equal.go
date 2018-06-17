package message

import "fmt"

//             greaterOrEqual  [5] AttributeValueAssertion,
func readFilterGreaterOrEqual(bytes *Bytes) (ret FilterGreaterOrEqual, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = readTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterGreaterOrEqual)
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilterGreaterOrEqual:\n%s", err.Error())}
		return
	}
	ret = FilterGreaterOrEqual(attributevalueassertion)
	return
}
