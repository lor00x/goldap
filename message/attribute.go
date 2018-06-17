package message

import "fmt"

//
//        Attribute ::= PartialAttribute(WITH COMPONENTS {
//             ...,
//             vals (SIZE(1..MAX))})
func readAttribute(bytes *Bytes) (ret Attribute, err error) {
	var par PartialAttribute
	par, err = readPartialAttribute(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAttribute:\n%s", err.Error())}
		return
	}
	if len(par.vals) == 0 {
		err = LdapError{"readAttribute: expecting at least one value"}
		return
	}
	ret = Attribute(par)
	return

}

//
//        Attribute ::= PartialAttribute(WITH COMPONENTS {
//             ...,
//             vals (SIZE(1..MAX))})
func (a Attribute) write(bytes *Bytes) (size int) {
	return PartialAttribute(a).write(bytes)
}
