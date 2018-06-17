package message

import "fmt"

//
//        ExtendedResponse ::= [APPLICATION 24] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             responseName     [10] LDAPOID OPTIONAL,
//             responseValue    [11] OCTET STRING OPTIONAL }
func readExtendedResponse(bytes *Bytes) (ret ExtendedResponse, err error) {
	err = bytes.ReadSubBytes(classApplication, TagExtendedResponse, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readExtendedResponse:\n%s", err.Error())}
		return
	}
	return
}
func (res *ExtendedResponse) readComponents(bytes *Bytes) (err error) {
	res.LDAPResult.readComponents(bytes)
	if bytes.HasMoreData() {
		var tag TagAndLength
		tag, err = bytes.PreviewTagAndLength()
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		if tag.Tag == TagExtendedResponseName {
			var oid LDAPOID
			oid, err = readTaggedLDAPOID(bytes, classContextSpecific, TagExtendedResponseName)
			if err != nil {
				err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
				return
			}
			res.responseName = oid.Pointer()
		}
	}
	if bytes.HasMoreData() {
		var tag TagAndLength
		tag, err = bytes.PreviewTagAndLength()
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		if tag.Tag == TagExtendedResponseValue {
			var responseValue OCTETSTRING
			responseValue, err = readTaggedOCTETSTRING(bytes, classContextSpecific, TagExtendedResponseValue)
			if err != nil {
				err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
				return
			}
			res.responseValue = responseValue.Pointer()
		}
	}
	return
}

//
//        ExtendedResponse ::= [APPLICATION 24] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             responseName     [10] LDAPOID OPTIONAL,
//             responseValue    [11] OCTET STRING OPTIONAL }
func (e ExtendedResponse) write(bytes *Bytes) (size int) {
	if e.responseValue != nil {
		size += e.responseValue.writeTagged(bytes, classContextSpecific, TagExtendedResponseValue)
	}
	if e.responseName != nil {
		size += e.responseName.writeTagged(bytes, classContextSpecific, TagExtendedResponseName)
	}
	size += e.LDAPResult.writeComponents(bytes)
	size += bytes.WriteTagAndLength(classApplication, isCompound, TagExtendedResponse, size)
	return
}
