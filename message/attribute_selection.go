package message

import "fmt"

//
//        AttributeSelection ::= SEQUENCE OF selector LDAPString
//                       -- The LDAPString is constrained to
//                       -- <attributeSelector> in Section 4.5.1.8
func readAttributeSelection(bytes *Bytes) (attributeSelection AttributeSelection, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, attributeSelection.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAttributeSelection:\n%s", err.Error())}
		return
	}
	return
}
func (attributeSelection *AttributeSelection) readComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var ldapstring LDAPString
		ldapstring, err = readLDAPString(bytes)
		// @TOTO: check <attributeSelector> in Section 4.5.1.8
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		*attributeSelection = append(*attributeSelection, ldapstring)
	}
	return
}

//
//        AttributeSelection ::= SEQUENCE OF selector LDAPString
//                       -- The LDAPString is constrained to
//                       -- <attributeSelector> in Section 4.5.1.8
func (a AttributeSelection) write(bytes *Bytes) (size int) {
	for i := len(a) - 1; i >= 0; i-- {
		size += a[i].write(bytes)
	}
	size += bytes.WriteTagAndLength(classUniversal, isCompound, tagSequence, size)
	return
}
