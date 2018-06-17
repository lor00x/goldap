package message

import "fmt"

func readModifyRequestChange(bytes *Bytes) (ret ModifyRequestChange, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readModifyRequestChange:\n%s", err.Error())}
		return
	}
	return
}
func (m *ModifyRequestChange) readComponents(bytes *Bytes) (err error) {
	m.operation, err = readENUMERATED(bytes, EnumeratedModifyRequestChangeOperation)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	m.modification, err = readPartialAttribute(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}
func (m ModifyRequestChange) write(bytes *Bytes) (size int) {
	size += m.modification.write(bytes)
	size += m.operation.write(bytes)
	size += bytes.WriteTagAndLength(classUniversal, isCompound, tagSequence, size)
	return
}
