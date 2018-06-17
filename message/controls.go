package message

import "fmt"

//
//        Controls ::= SEQUENCE OF control Control
func readTaggedControls(bytes *Bytes, class int, tag int) (controls Controls, err error) {
	err = bytes.ReadSubBytes(class, tag, controls.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedControls:\n%s", err.Error())}
		return
	}
	return
}
func (controls *Controls) readComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var control Control
		control, err = readControl(bytes)
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		*controls = append(*controls, control)
	}
	return
}
func (controls Controls) Pointer() *Controls { return &controls }
