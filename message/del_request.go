package message

import "fmt"

//
//        DelRequest ::= [APPLICATION 10] LDAPDN
func readDelRequest(bytes *Bytes) (ret DelRequest, err error) {
	var res LDAPDN
	res, err = readTaggedLDAPDN(bytes, classApplication, TagDelRequest)
	if err != nil {
		err = LdapError{fmt.Sprintf("readDelRequest:\n%s", err.Error())}
		return
	}
	ret = DelRequest(res)
	return
}

//
//        DelRequest ::= [APPLICATION 10] LDAPDN
func (d DelRequest) write(bytes *Bytes) int {
	return LDAPDN(d).writeTagged(bytes, classApplication, TagDelRequest)
}
