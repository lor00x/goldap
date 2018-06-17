package message

import "fmt"

//
//        CompareResponse ::= [APPLICATION 15] LDAPResult
func readCompareResponse(bytes *Bytes) (ret CompareResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagCompareResponse)
	if err != nil {
		err = LdapError{fmt.Sprintf("readCompareResponse:\n%s", err.Error())}
		return
	}
	ret = CompareResponse(res)
	return
}

//
//        CompareResponse ::= [APPLICATION 15] LDAPResult
func (c CompareResponse) write(bytes *Bytes) int {
	return LDAPResult(c).writeTagged(bytes, classApplication, TagCompareResponse)
}
