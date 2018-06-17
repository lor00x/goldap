package message

import (
	"errors"
	"fmt"
)

//
//        Control ::= SEQUENCE {
//             controlType             LDAPOID,
//             criticality             BOOLEAN DEFAULT FALSE,
//             controlValue            OCTET STRING OPTIONAL }
func readControl(bytes *Bytes) (control Control, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, control.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readControl:\n%s", err.Error())}
		return
	}
	return
}
func (control *Control) readComponents(bytes *Bytes) (err error) {
	control.controlType, err = readLDAPOID(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	if bytes.HasMoreData() {
		var tag TagAndLength
		tag, err = bytes.PreviewTagAndLength()
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		if tag.Tag == tagBoolean {
			control.criticality, err = readBOOLEAN(bytes)
			if err != nil {
				err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
				return
			}
			if control.criticality == false {
				err = errors.New(fmt.Sprintf("readComponents: criticality default value FALSE should not be specified"))
				return
			}
		}
	}
	if bytes.HasMoreData() {
		var octetstring OCTETSTRING
		octetstring, err = readOCTETSTRING(bytes)
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		control.controlValue = octetstring.Pointer()
	}
	return
}

//
//        Control ::= SEQUENCE {
//             controlType             LDAPOID,
//             criticality             BOOLEAN DEFAULT FALSE,
//             controlValue            OCTET STRING OPTIONAL }
func (c Control) write(bytes *Bytes) (size int) {
	if c.controlValue != nil {
		size += c.controlValue.write(bytes)
	}
	if c.criticality != BOOLEAN(false) {
		size += c.criticality.write(bytes)
	}
	size += c.controlType.write(bytes)
	size += bytes.WriteTagAndLength(classUniversal, isCompound, tagSequence, size)
	return
}

//
//        Control ::= SEQUENCE {
//             controlType             LDAPOID,
//             criticality             BOOLEAN DEFAULT FALSE,
//             controlValue            OCTET STRING OPTIONAL }
func (c Control) size() (size int) {
	if c.controlValue != nil {
		size += c.controlValue.size()
	}
	if c.criticality != BOOLEAN(false) {
		size += c.criticality.size()
	}
	size += c.controlType.size()
	size += sizeTagAndLength(tagSequence, size)
	return
}
