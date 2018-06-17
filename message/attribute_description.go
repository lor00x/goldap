package message

import "fmt"

//
//        AttributeDescription ::= LDAPString
//                                -- Constrained to <attributedescription>
//                                -- [RFC4512]
func readAttributeDescription(bytes *Bytes) (ret AttributeDescription, err error) {
	var ldapstring LDAPString
	ldapstring, err = readLDAPString(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAttributeDescription:\n%s", err.Error())}
		return
	}
	// @TODO: check RFC4512
	ret = AttributeDescription(ldapstring)
	return
}
func readTaggedAttributeDescription(bytes *Bytes, class int, tag int) (ret AttributeDescription, err error) {
	var ldapstring LDAPString
	ldapstring, err = readTaggedLDAPString(bytes, class, tag)
	// @TODO: check RFC4512
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedAttributeDescription:\n%s", err.Error())}
		return
	}
	ret = AttributeDescription(ldapstring)
	return
}
func (a AttributeDescription) Pointer() *AttributeDescription { return &a }
