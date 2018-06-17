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

//
//        AttributeDescription ::= LDAPString
//                                -- Constrained to <attributedescription>
//                                -- [RFC4512]
func (a AttributeDescription) write(bytes *Bytes) int {
	return LDAPString(a).write(bytes)
}
func (a AttributeDescription) writeTagged(bytes *Bytes, class int, tag int) int {
	return LDAPString(a).writeTagged(bytes, class, tag)
}
