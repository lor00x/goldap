package message

import "fmt"

//
//        SearchRequest ::= [APPLICATION 3] SEQUENCE {
//             baseObject      LDAPDN,
//             scope           ENUMERATED {
//                  baseObject              (0),
//                  singleLevel             (1),
//                  wholeSubtree            (2),
//                  ...  },
//             derefAliases    ENUMERATED {
//                  neverDerefAliases       (0),
//                  derefInSearching        (1),
//                  derefFindingBaseObj     (2),
//                  derefAlways             (3) },
//             sizeLimit       INTEGER (0 ..  maxInt),
//             timeLimit       INTEGER (0 ..  maxInt),
//             typesOnly       BOOLEAN,
//             filter          Filter,
//             attributes      AttributeSelection }
func readSearchRequest(bytes *Bytes) (searchrequest SearchRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagSearchRequest, searchrequest.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readSearchRequest:\n%s", err.Error())}
		return
	}
	return
}
func (searchrequest *SearchRequest) readComponents(bytes *Bytes) (err error) {
	searchrequest.baseObject, err = readLDAPDN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	searchrequest.scope, err = readENUMERATED(bytes, EnumeratedSearchRequestScope)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	searchrequest.derefAliases, err = readENUMERATED(bytes, EnumeratedSearchRequestDerefAliases)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	searchrequest.sizeLimit, err = readPositiveINTEGER(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	searchrequest.timeLimit, err = readPositiveINTEGER(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	searchrequest.typesOnly, err = readBOOLEAN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	searchrequest.filter, err = readFilter(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	searchrequest.attributes, err = readAttributeSelection(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}

//
//        SearchRequest ::= [APPLICATION 3] SEQUENCE {
//             baseObject      LDAPDN,
//             scope           ENUMERATED {
//                  baseObject              (0),
//                  singleLevel             (1),
//                  wholeSubtree            (2),
//                  ...  },
//             derefAliases    ENUMERATED {
//                  neverDerefAliases       (0),
//                  derefInSearching        (1),
//                  derefFindingBaseObj     (2),
//                  derefAlways             (3) },
//             sizeLimit       INTEGER (0 ..  maxInt),
//             timeLimit       INTEGER (0 ..  maxInt),
//             typesOnly       BOOLEAN,
//             filter          Filter,
//             attributes      AttributeSelection }
func (s SearchRequest) write(bytes *Bytes) (size int) {
	size += s.attributes.write(bytes)
	size += s.filter.write(bytes)
	size += s.typesOnly.write(bytes)
	size += s.timeLimit.write(bytes)
	size += s.sizeLimit.write(bytes)
	size += s.derefAliases.write(bytes)
	size += s.scope.write(bytes)
	size += s.baseObject.write(bytes)
	size += bytes.WriteTagAndLength(classApplication, isCompound, TagSearchRequest, size)
	return
}
