package message

import "fmt"

//
//        CompareRequest ::= [APPLICATION 14] SEQUENCE {
//             entry           LDAPDN,
//             ava             AttributeValueAssertion }
func readCompareRequest(bytes *Bytes) (ret CompareRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagCompareRequest, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readCompareRequest:\n%s", err.Error())}
		return
	}
	return
}
func (req *CompareRequest) readComponents(bytes *Bytes) (err error) {
	req.entry, err = readLDAPDN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	req.ava, err = readAttributeValueAssertion(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}
