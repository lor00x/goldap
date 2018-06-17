package message

import "fmt"

//
//        PartialAttributeList ::= SEQUENCE OF
//                             partialAttribute PartialAttribute
func readPartialAttributeList(bytes *Bytes) (ret PartialAttributeList, err error) {
	ret = PartialAttributeList(make([]PartialAttribute, 0, 10))
	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readPartialAttributeList:\n%s", err.Error())}
		return
	}
	return
}
func (partialattributelist *PartialAttributeList) readComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var partialattribute PartialAttribute
		partialattribute, err = readPartialAttribute(bytes)
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		*partialattributelist = append(*partialattributelist, partialattribute)
	}
	return
}