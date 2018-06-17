package message

import "fmt"

func readTaggedLDAPOID(bytes *Bytes, class int, tag int) (ret LDAPOID, err error) {
	var octetstring OCTETSTRING
	octetstring, err = readTaggedOCTETSTRING(bytes, class, tag)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedLDAPOID:\n%s", err.Error())}
		return
	}
	// @TODO: check RFC4512 for <numericoid>
	ret = LDAPOID(octetstring)
	return
}

//
//
//
//
//Sermersheim                 Standards Track                    [Page 54]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//        LDAPOID ::= OCTET STRING -- Constrained to <numericoid>
//                                 -- [RFC4512]
func readLDAPOID(bytes *Bytes) (ret LDAPOID, err error) {
	return readTaggedLDAPOID(bytes, classUniversal, tagOctetString)
}
func (l LDAPOID) Pointer() *LDAPOID { return &l }
