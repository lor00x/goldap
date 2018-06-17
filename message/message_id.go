package message

import "fmt"

func readTaggedMessageID(bytes *Bytes, class int, tag int) (ret MessageID, err error) {
	var integer INTEGER
	integer, err = readTaggedPositiveINTEGER(bytes, class, tag)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedMessageID:\n%s", err.Error())}
		return
	}
	return MessageID(integer), err
}

//        MessageID ::= INTEGER (0 ..  maxInt)
//
//        maxInt INTEGER ::= 2147483647 -- (2^^31 - 1) --
//
func readMessageID(bytes *Bytes) (ret MessageID, err error) {
	return readTaggedMessageID(bytes, classUniversal, tagInteger)
}
