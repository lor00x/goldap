package goldap

import (
	"goldap/asn1"
	"reflect"
	"testing"
)

var bindMessageData = []struct {
	in  []byte
	out *BindMessage
}{
	{
		in: []byte{0x30, 0x1d, 0x02, 0x01, 0x05, 0x60, 0x18, 0x02, 0x01, 0x03, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64},
		out: &BindMessage{
			MessageId: int(0x05),
			ProtocolOp: BindRequest{
				Version:  int(0x03),
				Login:    []byte("myLogin"),
				Password: []byte("myPassword"),
			},
		},
	},
}

func TestReadBindMessage(t *testing.T) {
	for i, test := range bindMessageData {
		var bindRequest = ReadBindMessage(test.in)
		if test.out.GetVersion() != bindRequest.GetVersion() {
			t.Errorf("#%d: Bad version: %d (expected %d)", i, bindRequest.GetVersion(), test.out.GetVersion())
		}
		if test.out.GetLogin() != bindRequest.GetLogin() {
			t.Errorf("#%d: Bad login: %s (expected %s)", i, bindRequest.GetLogin(), test.out.GetLogin())
		}
		if test.out.GetPassword() != bindRequest.GetPassword() {
			t.Errorf("#%d: Bad password: %s (expected %s)", i, bindRequest.GetPassword(), test.out.GetPassword())
		}
	}
}

var saslBindMessageData = []struct {
	in  []byte
	out *SaslBindMessage
}{
	{
		in: []byte{0x30, 0x16, 0x02, 0x01, 0x01, 0x60, 0x11, 0x02, 0x01, 0x03, 0x04, 0x00, 0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35},
		out: &SaslBindMessage{
			MessageId:		int(0x01),
			ProtocolOp:		SaslBindRequest{
				Version:		int(0x03),
				Name:			[]byte(""),
				Authentication:	SaslCredentials{
					Mechanism: []byte("CRAM-MD5"),
				},
			},
		},
	},
}

func TestReadSaslBindMessage(t *testing.T) {
	for i, test := range saslBindMessageData {
		var saslBindRequest = ReadSaslBindMessage(test.in)
		if test.out.GetVersion() != saslBindRequest.GetVersion() {
			t.Errorf("#%d: Bad version: %d (expected %d)", i, saslBindRequest.GetVersion(), test.out.GetVersion())
		}
		if test.out.GetName() != saslBindRequest.GetName() {
			t.Errorf("#%d: Bad login: %s (expected %s)", i, saslBindRequest.GetName(), test.out.GetName())
		}
		if test.out.GetMechanism() != saslBindRequest.GetMechanism() {
			t.Errorf("#%d: Bad mechanism: %d (expected %d)", i, saslBindRequest.GetMechanism(), test.out.GetMechanism())
		}
	}
}

func newInt(n int) *int { return &n }

type TestContextSpecificTags struct {
	A int `asn1:"tag:3"`
}
//
//type TestBindOp struct {
//	Version int
//}
//
type TestBindRequest struct {
	ProtocolOp BindRequest `asn1:"application,tag:0"`
}

var unmarshalTestData = []struct {
	in  []byte
	out interface{}
}{
	{[]byte{0x02, 0x01, 0x42}, newInt(0x42)},
	//	{[]byte{0x30, 0x08, 0x06, 0x06, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d}, &TestObjectIdentifierStruct{[]int{1, 2, 840, 113549}}},
	//	{[]byte{0x03, 0x04, 0x06, 0x6e, 0x5d, 0xc0}, &BitString{[]byte{110, 93, 192}, 18}},
	//	{[]byte{0x30, 0x09, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02, 0x02, 0x01, 0x03}, &[]int{1, 2, 3}},
	//	{[]byte{0x02, 0x01, 0x10}, newInt(16)},
	//	{[]byte{0x13, 0x04, 't', 'e', 's', 't'}, newString("test")},
	//	{[]byte{0x16, 0x04, 't', 'e', 's', 't'}, newString("test")},
	//	{[]byte{0x16, 0x04, 't', 'e', 's', 't'}, &RawValue{0, 22, false, []byte("test"), []byte("\x16\x04test")}},
	//	{[]byte{0x04, 0x04, 1, 2, 3, 4}, &RawValue{0, 4, false, []byte{1, 2, 3, 4}, []byte{4, 4, 1, 2, 3, 4}}},
	{[]byte{0x30, 0x03, 0x83, 0x01, 0x02}, &TestContextSpecificTags{2}},
	//	{[]byte{0x30, 0x03, 0x83, 0x01, 0x02}, &TestBindRequest{ TestContextSpecificTags{2}}},
	//	{[]byte{0x30, 0x08, 0xa1, 0x03, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02}, &TestContextSpecificTags2{1, 2}},
	//	{[]byte{0x01, 0x01, 0x00}, newBool(false)},
	//	{[]byte{0x01, 0x01, 0xff}, newBool(true)},
	//	{[]byte{0x30, 0x0b, 0x13, 0x03, 0x66, 0x6f, 0x6f, 0x02, 0x01, 0x22, 0x02, 0x01, 0x33}, &TestElementsAfterString{"foo", 0x22, 0x33}},
	//	{[]byte{0x30, 0x05, 0x02, 0x03, 0x12, 0x34, 0x56}, &TestBigInt{big.NewInt(0x123456)}},
	//	{[]byte{0x30, 0x0b, 0x31, 0x09, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02, 0x02, 0x01, 0x03}, &TestSet{Ints: []int{1, 2, 3}}},
	{
		in: []byte{0x30, 0x1d, 0x02, 0x01, 0x05, 0x60, 0x18, 0x02, 0x01, 0x03, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64},
		out: &BindMessage{
			MessageId: int(0x05),
			ProtocolOp: BindRequest{
				Version:  int(0x03),
				Login:    []byte("myLogin"),
				Password: []byte("myPassword"),
			},
		},
	},

	{
		in: []byte{0x30, 0x1a, 0x60, 0x18, 0x02, 0x01, 0x03, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64},
		out: &TestBindRequest{
			ProtocolOp: BindRequest{
				Version:  int(0x03),
				Login:    []byte("myLogin"),
				Password: []byte("myPassword"),
			},
		},
	},

	{
		in: []byte{0x30, 0x18, 0x02, 0x01, 0x03, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64},
		out: &BindRequest{
			Version:  int(0x03),
			Login:    []byte("myLogin"),
			Password: []byte("myPassword"),
		},
	},

	{
		in: []byte{0x30, 0x08, 0x04, 0x02, 0x66, 0x61, 0x04, 0x02, 0x65, 0x65},
		out: &SaslCredentials{
			Mechanism: []byte("fa"),
			Credentials: []byte("ee"),
		},
	},

	{
		in: []byte{0x30, 0x04, 0x04, 0x02, 0x66, 0x61},
		out: &SaslCredentials{
			Mechanism: []byte("fa"),
		},
	},

	{
		in: []byte{0x30, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35},
		out: &SaslCredentials{
			Mechanism: []byte("CRAM-MD5"),
		},
	},

	{
		in: []byte{0x30, 0x11, 0x02, 0x01, 0x03, 0x04, 0x00, 0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35},
		out: &SaslBindRequest{
			Version:		int(0x03),
			Name:			[]byte(""),
			Authentication:	SaslCredentials{
				Mechanism: []byte("CRAM-MD5"),
			},
		},
	},
	
	{
		in: []byte{0x30, 0x16, 0x02, 0x01, 0x01, 0x60, 0x11, 0x02, 0x01, 0x03, 0x04, 0x00, 0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35},
		out: &SaslBindMessage{
			MessageId:		int(0x01),
			ProtocolOp:		SaslBindRequest{
				Version:		int(0x03),
				Name:			[]byte(""),
				Authentication:	SaslCredentials{
					Mechanism: []byte("CRAM-MD5"),
				},
			},
		},
	},
	
}

func TestUnmarshal(t *testing.T) {
	for i, test := range unmarshalTestData {
		pv := reflect.New(reflect.TypeOf(test.out).Elem())
		val := pv.Interface()
		rest, err := asn1.Unmarshal(test.in, val)
		if len(rest) > 0 {
			t.Errorf("Rest bytes %v", rest)
		}
		if err != nil {
			t.Errorf("Unmarshal failed at index %d %v", i, err)
		}
		if !reflect.DeepEqual(val, test.out) {
			t.Errorf("#%d:\nhave %#v\nwant %#v", i, val, test.out)
		}
	}
}
