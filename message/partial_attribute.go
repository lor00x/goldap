package message

import "fmt"

//
//        PartialAttribute ::= SEQUENCE {
//             type       AttributeDescription,
//             vals       SET OF value AttributeValue }
func readPartialAttribute(bytes *Bytes) (ret PartialAttribute, err error) {
	ret = PartialAttribute{vals: make([]AttributeValue, 0, 10)}
	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readPartialAttribute:\n%s", err.Error())}
		return
	}
	return
}
func (partialattribute *PartialAttribute) readComponents(bytes *Bytes) (err error) {
	partialattribute.type_, err = readAttributeDescription(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	err = bytes.ReadSubBytes(classUniversal, tagSet, partialattribute.readValsComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}
func (partialattribute *PartialAttribute) readValsComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var attributevalue AttributeValue
		attributevalue, err = readAttributeValue(bytes)
		if err != nil {
			err = LdapError{fmt.Sprintf("readValsComponents:\n%s", err.Error())}
			return
		}
		partialattribute.vals = append(partialattribute.vals, attributevalue)
	}
	return
}
