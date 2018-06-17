package message

import "fmt"

//
//        DelResponse ::= [APPLICATION 11] LDAPResult
func readDelResponse(bytes *Bytes) (ret DelResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagDelResponse)
	if err != nil {
		err = LdapError{fmt.Sprintf("readDelResponse:\n%s", err.Error())}
		return
	}
	ret = DelResponse(res)
	return
}
