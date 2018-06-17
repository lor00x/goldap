package message

import "fmt"

//
//
//
//
//Sermersheim                 Standards Track                    [Page 56]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//        BindRequest ::= [APPLICATION 0] SEQUENCE {
//             version                 INTEGER (1 ..  127),
//             name                    LDAPDN,
//             authentication          AuthenticationChoice }
func readBindRequest(bytes *Bytes) (bindrequest BindRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagBindRequest, bindrequest.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readBindRequest:\n%s", err.Error())}
		return
	}
	return
}
func (bindrequest *BindRequest) readComponents(bytes *Bytes) (err error) {
	bindrequest.version, err = readINTEGER(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	if !(bindrequest.version >= BindRequestVersionMin && bindrequest.version <= BindRequestVersionMax) {
		err = LdapError{fmt.Sprintf("readComponents: invalid version %d, must be between %d and %d", bindrequest.version, BindRequestVersionMin, BindRequestVersionMax)}
		return
	}
	bindrequest.name, err = readLDAPDN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	bindrequest.authentication, err = readAuthenticationChoice(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}
