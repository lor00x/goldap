package message

import "fmt"

//
//        Referral ::= SEQUENCE SIZE (1..MAX) OF uri URI
func readTaggedReferral(bytes *Bytes, class int, tag int) (referral Referral, err error) {
	err = bytes.ReadSubBytes(class, tag, referral.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedReferral:\n%s", err.Error())}
		return
	}
	return
}
func (referral *Referral) readComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var uri URI
		uri, err = readURI(bytes)
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		*referral = append(*referral, uri)
	}
	if len(*referral) == 0 {
		return LdapError{"readComponents: expecting at least one URI"}
	}
	return
}
func (referral Referral) Pointer() *Referral { return &referral }
