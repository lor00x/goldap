package message

import "fmt"

//
//        MatchingRuleId ::= LDAPString
func readTaggedMatchingRuleId(bytes *Bytes, class int, tag int) (matchingruleid MatchingRuleId, err error) {
	var ldapstring LDAPString
	ldapstring, err = readTaggedLDAPString(bytes, class, tag)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedMatchingRuleId:\n%s", err.Error())}
		return
	}
	matchingruleid = MatchingRuleId(ldapstring)
	return
}
func (m MatchingRuleId) Pointer() *MatchingRuleId { return &m }
