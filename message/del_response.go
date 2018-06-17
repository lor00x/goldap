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

//
//        DelResponse ::= [APPLICATION 11] LDAPResult
func (d DelResponse) write(bytes *Bytes) int {
	return LDAPResult(d).writeTagged(bytes, classApplication, TagDelResponse)
}

//
//        DelResponse ::= [APPLICATION 11] LDAPResult
func (d DelResponse) size() int {
	return LDAPResult(d).sizeTagged(TagDelResponse)
}
