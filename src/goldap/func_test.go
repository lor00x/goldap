package goldap

import (
	"reflect"
	"testing"
)

var ReadMessageIDTestData = []struct {
	bytes     *Bytes
	messageID MessageID
	offset    int
}{
	{messageID: MessageID(0x0f), offset: 5, bytes: &Bytes{offset: 2, bytes: []byte{0x30, 0x03, 0x02, 0x01, 0x0f}}},
	{messageID: MessageID(0x0f), offset: 5, bytes: &Bytes{offset: 2, bytes: []byte{0x30, 0x16, 0x02, 0x01, 0x0f, 0x60, 0x11, 0x02, 0x01, 0x03, 0x04, 0x00, 0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35}}},
	{messageID: MessageID(maxInt), offset: 8, bytes: &Bytes{offset: 2, bytes: []byte{0x30, 0x19, 0x02, 0x04, 0x7f, 0xff, 0xff, 0xff, 0x60, 0x11, 0x02, 0x01, 0x03, 0x04, 0x00, 0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35}}},
}

func TestReadMessageID(t *testing.T) {
	for i, test := range ReadMessageIDTestData {
		message := NewLDAPMessage()
		var err error
		message.messageID, err = ReadMessageID(test.bytes)
		if err != nil {
			t.Errorf("#%d failed at offset %d (%#x): %s", i, test.bytes.offset, test.bytes.bytes[test.bytes.offset], err)
		}
		if !reflect.DeepEqual(test.messageID, message.messageID) {
			t.Errorf("#%d: Wrong MessageID, expected %#v, got %#v", i, test.messageID, message.messageID)
		}
		if test.offset != test.bytes.offset {
			t.Errorf("#%d: Wrong Offset, expected %#v, got %#v", i, test.offset, test.bytes.offset)
		}
	}
}

var ReadLDAPMessageTestData = []struct {
	bytes *Bytes
	out LDAPMessage
}{
	{
		bytes: &Bytes{
			offset: 0,
			bytes: []byte{
				0x30, 0x1d,
					0x02, 0x01, 0x01,	// messageID
					0x60, 0x18,			// Application, tag 0 => this is a Bind request
						0x02, 0x01, 0x03,	// Version 3
						0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e,	// login = myLogin
						0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, // simple authentication: myPassword
						},
		},
		out: LDAPMessage{
			messageID: MessageID(int(0x01)),
			protocolOp: BindRequest{
				version: 0x03,
				name: "myLogin",
				authentication: OCTETSTRING([]byte("myPassword")),
			},
		},
	},
}

func TestReadLDAPMessage(t *testing.T) {
	for i, test := range ReadLDAPMessageTestData {
		message, err := ReadLDAPMessage(test.bytes)
		if err != nil {
			t.Errorf("#%d failed at offset %d (%#x): %s", i, test.bytes.offset, test.bytes.bytes[test.bytes.offset], err)
		}
		if !reflect.DeepEqual(message, test.out) {
			t.Errorf("#%d:\nhave %#+v\nwant %#+v", i, message, test.out)
		}
	}
}

var ReadAuthenticationChoiceTestData = []struct {
	bytes *Bytes
	out   AuthenticationChoice
}{
	{
		bytes: &Bytes{offset: 19, bytes: []byte{
			0x30, 0x29,
			0x02, 0x01, 0x05,
			0x60, 0x24,
			0x02, 0x01, 0x03,
			0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
			0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64}},
		out: OCTETSTRING("myPassword"),
	},
	{
		bytes: &Bytes{offset: 12, bytes: []byte{
			0x30, 0x29,
			0x02, 0x01, 0x05,
			0x60, 0x24,
			0x02, 0x01, 0x03,
			0x04, 0x00,
			0xa3, 0x16, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35,
			0x04, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64}},
		out: SaslCredentials{
			mechanism: LDAPString("CRAM-MD5"),
			credentials: NewOCTETSTRING([]byte("myPassword")),
		},
	},

}


func TestReadReadAuthenticationChoice(t *testing.T) {
	for i, test := range ReadAuthenticationChoiceTestData {
		authenticationChoice, err := ReadAuthenticationChoice(test.bytes)
		if err != nil {
			t.Errorf("#%d failed at offset %d (%#x): %s", i, test.bytes.offset, test.bytes.bytes[test.bytes.offset], err)
		} else if !reflect.DeepEqual(authenticationChoice, test.out) {
			t.Errorf("#%d:\nhave %#+v\nwant %#+v", i, authenticationChoice, test.out)
		}
	}
}

func NewOCTETSTRING(bytes []byte) *OCTETSTRING {
	octetstring := OCTETSTRING(bytes)
	return &octetstring
}

var ReadSaslCredentialsTestData = []struct {
	bytes *Bytes
	out   SaslCredentials
}{
	{
		bytes: &Bytes{offset: 19, bytes: []byte{
			0x30, 0x29,
			0x02, 0x01, 0x05,
			0x60, 0x24,
			0x02, 0x01, 0x03,
			0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
			0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35}},
		out: SaslCredentials{
			mechanism: LDAPString("CRAM-MD5"),
		},
	},
	{
		bytes: &Bytes{offset: 0, bytes: []byte{
			0xa3, 0x14, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35,
			0x04, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}},
		out: SaslCredentials{
			mechanism:   LDAPString("CRAM-MD5"),
			credentials: NewOCTETSTRING([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}),
		},
	},
}

func TestReadSaslCredentials(t *testing.T) {
	for i, test := range ReadSaslCredentialsTestData {
		saslCredentials, err := ReadSaslCredentials(test.bytes)
		if err != nil {
			t.Errorf("#%d failed: %s", i, err)
		} else if !reflect.DeepEqual(saslCredentials, test.out) {
			t.Errorf("#%d:\nhave %#+v\nwant %#+v", i, saslCredentials, test.out)
		}
	}
}
