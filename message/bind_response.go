package message

import "fmt"

//        BindResponse ::= [APPLICATION 1] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             serverSaslCreds    [7] OCTET STRING OPTIONAL }
func readBindResponse(bytes *Bytes) (bindresponse BindResponse, err error) {
	err = bytes.ReadSubBytes(classApplication, TagBindResponse, bindresponse.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readBindResponse:\n%s", err.Error())}
		return
	}
	return
}
func (bindresponse *BindResponse) readComponents(bytes *Bytes) (err error) {
	bindresponse.LDAPResult.readComponents(bytes)
	if bytes.HasMoreData() {
		var tag TagAndLength
		tag, err = bytes.PreviewTagAndLength()
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		if tag.Tag == TagBindResponseServerSaslCreds {
			var serverSaslCreds OCTETSTRING
			serverSaslCreds, err = readTaggedOCTETSTRING(bytes, classContextSpecific, TagBindResponseServerSaslCreds)
			if err != nil {
				err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
				return
			}
			bindresponse.serverSaslCreds = serverSaslCreds.Pointer()
		}
	}
	return
}

//        BindResponse ::= [APPLICATION 1] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             serverSaslCreds    [7] OCTET STRING OPTIONAL }
func (b BindResponse) write(bytes *Bytes) (size int) {
	if b.serverSaslCreds != nil {
		size += b.serverSaslCreds.writeTagged(bytes, classContextSpecific, TagBindResponseServerSaslCreds)
	}
	size += b.LDAPResult.writeComponents(bytes)
	size += bytes.WriteTagAndLength(classApplication, isCompound, TagBindResponse, size)
	return
}
