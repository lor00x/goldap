package message

import "fmt"

//             extensibleMatch [9] MatchingRuleAssertion,
func readFilterExtensibleMatch(bytes *Bytes) (filterextensiblematch FilterExtensibleMatch, err error) {
	var matchingruleassertion MatchingRuleAssertion
	matchingruleassertion, err = readTaggedMatchingRuleAssertion(bytes, classContextSpecific, TagFilterExtensibleMatch)
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilterExtensibleMatch:\n%s", err.Error())}
		return
	}
	filterextensiblematch = FilterExtensibleMatch(matchingruleassertion)
	return
}
