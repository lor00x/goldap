package message

import (
  "testing"
  "fmt"
)

func toHex(b []byte) (r string) {
  r = "[ "
  for _, e := range b {
    r += fmt.Sprintf("0x%x ", e)
  }
  return r + "]"
}

func TestMessageID(t *testing.T) {
	m := NewLDAPMessageWithProtocolOp(UnbindRequest{})
	m.SetMessageID(128)
	buf, err := m.Write()
	if err != nil {
		t.Errorf("marshalling failed with %v", err)
	}
	t.Logf("%v", toHex(buf.Bytes()))

	ret, err := ReadLDAPMessage(NewBytes(0, buf.Bytes()))
	if err != nil {
		t.Errorf("unmarshalling failed with %v", err)
	}
	if _, ok := ret.ProtocolOp().(UnbindRequest); !ok {
		t.Errorf("should be an unbind request")
	}
	if ret.MessageID() != 128 {
		t.Errorf("Expect message id 128, got %d", ret.MessageID())
	}
	t.Log("Done, marshal/unmarshall worked")
}
