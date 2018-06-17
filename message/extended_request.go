package message

import "fmt"

//
//        ExtendedRequest ::= [APPLICATION 23] SEQUENCE {
//             requestName      [0] LDAPOID,
//             requestValue     [1] OCTET STRING OPTIONAL }
func readExtendedRequest(bytes *Bytes) (ret ExtendedRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagExtendedRequest, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readExtendedRequest:\n%s", err.Error())}
		return
	}
	return
}
func (req *ExtendedRequest) readComponents(bytes *Bytes) (err error) {
	req.requestName, err = readTaggedLDAPOID(bytes, classContextSpecific, TagExtendedRequestName)
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
		if tag.Tag == TagExtendedRequestValue {
			var requestValue OCTETSTRING
			requestValue, err = readTaggedOCTETSTRING(bytes, classContextSpecific, TagExtendedRequestValue)
			if err != nil {
				err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
				return
			}
			req.requestValue = requestValue.Pointer()
		}
	}
	return
}

//
//        ExtendedRequest ::= [APPLICATION 23] SEQUENCE {
//             requestName      [0] LDAPOID,
//             requestValue     [1] OCTET STRING OPTIONAL }
func (e ExtendedRequest) write(bytes *Bytes) (size int) {
	if e.requestValue != nil {
		size += e.requestValue.writeTagged(bytes, classContextSpecific, TagExtendedRequestValue)
	}
	size += e.requestName.writeTagged(bytes, classContextSpecific, TagExtendedRequestName)
	size += bytes.WriteTagAndLength(classApplication, isCompound, TagExtendedRequest, size)
	return
}

//
//        ExtendedRequest ::= [APPLICATION 23] SEQUENCE {
//             requestName      [0] LDAPOID,
//             requestValue     [1] OCTET STRING OPTIONAL }
func (e ExtendedRequest) size() (size int) {
	size += e.requestName.sizeTagged(TagExtendedRequestName)
	if e.requestValue != nil {
		size += e.requestValue.sizeTagged(TagExtendedRequestValue)
	}
	size += sizeTagAndLength(TagExtendedRequest, size)
	return
}
