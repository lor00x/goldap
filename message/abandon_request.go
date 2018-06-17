package message

import "fmt"

//
//        AbandonRequest ::= [APPLICATION 16] MessageID
func (a AbandonRequest) size() int {
	return MessageID(a).sizeTagged(TagAbandonRequest)
}

//
//        AbandonRequest ::= [APPLICATION 16] MessageID
func readAbandonRequest(bytes *Bytes) (ret AbandonRequest, err error) {
	var mes MessageID
	mes, err = readTaggedMessageID(bytes, classApplication, TagAbandonRequest)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAbandonRequest:\n%s", err.Error())}
		return
	}
	ret = AbandonRequest(mes)
	return
}

//
//        AbandonRequest ::= [APPLICATION 16] MessageID
func (a AbandonRequest) write(bytes *Bytes) int {
	return MessageID(a).writeTagged(bytes, classApplication, TagAbandonRequest)
}
