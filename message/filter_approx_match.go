package message

import "fmt"

//             approxMatch     [8] AttributeValueAssertion,
func readFilterApproxMatch(bytes *Bytes) (ret FilterApproxMatch, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = readTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterApproxMatch)
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilterApproxMatch:\n%s", err.Error())}
		return
	}
	ret = FilterApproxMatch(attributevalueassertion)
	return
}
