package message

import "fmt"

//
//        SearchResultEntry ::= [APPLICATION 4] SEQUENCE {
//             objectName      LDAPDN,
//             attributes      PartialAttributeList }
func readSearchResultEntry(bytes *Bytes) (searchresultentry SearchResultEntry, err error) {
	err = bytes.ReadSubBytes(classApplication, TagSearchResultEntry, searchresultentry.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readSearchResultEntry:\n%s", err.Error())}
		return
	}
	return
}
func (searchresultentry *SearchResultEntry) readComponents(bytes *Bytes) (err error) {
	searchresultentry.objectName, err = readLDAPDN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	searchresultentry.attributes, err = readPartialAttributeList(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}

//
//        SearchResultEntry ::= [APPLICATION 4] SEQUENCE {
//             objectName      LDAPDN,
//             attributes      PartialAttributeList }
func (s SearchResultEntry) write(bytes *Bytes) (size int) {
	size += s.attributes.write(bytes)
	size += s.objectName.write(bytes)
	size += bytes.WriteTagAndLength(classApplication, isCompound, TagSearchResultEntry, size)
	return
}
