package message

import "fmt"

//
//        SearchResultDone ::= [APPLICATION 5] LDAPResult
func readSearchResultDone(bytes *Bytes) (ret SearchResultDone, err error) {
	var ldapresult LDAPResult
	ldapresult, err = readTaggedLDAPResult(bytes, classApplication, TagSearchResultDone)
	if err != nil {
		err = LdapError{fmt.Sprintf("readSearchResultDone:\n%s", err.Error())}
		return
	}
	ret = SearchResultDone(ldapresult)
	return
}
