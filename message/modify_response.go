package message

import "fmt"

//
//        ModifyResponse ::= [APPLICATION 7] LDAPResult
func readModifyResponse(bytes *Bytes) (ret ModifyResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagModifyResponse)
	if err != nil {
		err = LdapError{fmt.Sprintf("readModifyResponse:\n%s", err.Error())}
		return
	}
	ret = ModifyResponse(res)
	return
}
