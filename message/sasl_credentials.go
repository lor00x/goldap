package message

import "fmt"

//
//        SaslCredentials ::= SEQUENCE {
//             mechanism               LDAPString,
//             credentials             OCTET STRING OPTIONAL }
//
func readSaslCredentials(bytes *Bytes) (authentication SaslCredentials, err error) {
	authentication = SaslCredentials{}
	err = bytes.ReadSubBytes(classContextSpecific, TagAuthenticationChoiceSaslCredentials, authentication.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readSaslCredentials:\n%s", err.Error())}
		return
	}
	return
}
func (authentication *SaslCredentials) readComponents(bytes *Bytes) (err error) {
	authentication.mechanism, err = readLDAPString(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	if bytes.HasMoreData() {
		var credentials OCTETSTRING
		credentials, err = readOCTETSTRING(bytes)
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		authentication.credentials = credentials.Pointer()
	}
	return
}
