package message

import "fmt"

//             equalityMatch   [3] AttributeValueAssertion,
func readFilterEqualityMatch(bytes *Bytes) (ret FilterEqualityMatch, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = readTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterEqualityMatch)
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilterEqualityMatch:\n%s", err.Error())}
		return
	}
	ret = FilterEqualityMatch(attributevalueassertion)
	return
}
