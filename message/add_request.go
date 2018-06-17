package message

import "fmt"

//
//
//
//
//
//
//Sermersheim                 Standards Track                    [Page 58]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//        AddRequest ::= [APPLICATION 8] SEQUENCE {
//             entry           LDAPDN,
//             attributes      AttributeList }
func readAddRequest(bytes *Bytes) (ret AddRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagAddRequest, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAddRequest:\n%s", err.Error())}
		return
	}
	return
}
func (req *AddRequest) readComponents(bytes *Bytes) (err error) {
	req.entry, err = readLDAPDN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	req.attributes, err = readAttributeList(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}

//
//
//
//
//
//
//Sermersheim                 Standards Track                    [Page 58]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//        AddRequest ::= [APPLICATION 8] SEQUENCE {
//             entry           LDAPDN,
//             attributes      AttributeList }
func (a AddRequest) write(bytes *Bytes) (size int) {
	size += a.attributes.write(bytes)
	size += a.entry.write(bytes)
	size += bytes.WriteTagAndLength(classApplication, isCompound, TagAddRequest, size)
	return
}
