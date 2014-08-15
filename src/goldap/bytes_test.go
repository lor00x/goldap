package goldap

import (
	"reflect"
	"testing"
)

var ParseInt32TestData = []struct {
	bytes  *Bytes // Input
	length int    // Input
	value  int32  // Expected output
	offset int    // Expected output
}{
	{value: 0x09, offset: 1, length: 1, bytes: &Bytes{offset: 0, bytes: []byte{0x09}}},
	{value: 0x0987, offset: 2, bytes: &Bytes{offset: 0, bytes: []byte{0x09, 0x87}}, length: 2},
	{value: 0x098765, offset: 3, length: 3, bytes: &Bytes{offset: 0, bytes: []byte{0x09, 0x87, 0x65}}},
	{value: 0x09876543, offset: 4, length: 4, bytes: &Bytes{offset: 0, bytes: []byte{0x09, 0x87, 0x65, 0x43}}},
	{value: 0x0f, offset: 5, length: 1, bytes: &Bytes{offset: 4, bytes: []byte{0x30, 0x03, 0x02, 0x01, 0x0f}}},
	{value: 0x0f, offset: 5, length: 1, bytes: &Bytes{offset: 4, bytes: []byte{0x30, 0x16, 0x02, 0x01, 0x0f, 0x60, 0x11, 0x02, 0x01, 0x03, 0x04, 0x00, 0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35}}},
	{value: 0x7fffffff, offset: 8, length: 4, bytes: &Bytes{offset: 4, bytes: []byte{0x30, 0x19, 0x02, 0x04, 0x7f, 0xff, 0xff, 0xff, 0x60, 0x11, 0x02, 0x01, 0x03, 0x04, 0x00, 0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35}}},
}

func TestParseInt32(t *testing.T) {
	for i, test := range ParseInt32TestData {
		value, err := test.bytes.ParseInt32(test.length)
		if err != nil {
			t.Errorf("#%d failed: %s", i, err)
		}
		if !reflect.DeepEqual(test.value, value) {
			t.Errorf("#%d: Wrong int32, expected %#v, got %#v", i, &test.value, &value)
		}
		if test.offset != test.bytes.offset {
			t.Errorf("#%d: Wrong Offset, expected %#v, got %#v", i, test.offset, test.bytes.offset)
		}
	}
}

var ParseUTF8STRINGTestData = []struct {
	bytes  *Bytes     // Input
	length int        // Input
	value  UTF8STRING // Expected output
	offset int        // Expected output
}{
	{value: UTF8STRING("CRAM-MD5"), offset: 8, length: 8, bytes: &Bytes{offset: 0, bytes: []byte{0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35}}},
	{value: UTF8STRING("Hello, 世界"), offset: 13, length: 13, bytes: &Bytes{offset: 0, bytes: []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}}},
	{value: UTF8STRING("myLogin"), offset: 19, length: 7, bytes: &Bytes{offset: 12, bytes: []byte{0x30, 0x1d, 0x02, 0x01, 0x05, 0x60, 0x18, 0x02, 0x01, 0x03, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64}}},
}

func TestParseUTF8STRING(t *testing.T) {
	for i, test := range ParseUTF8STRINGTestData {
		value, err := test.bytes.ParseUTF8STRING(test.length)
		if err != nil {
			t.Errorf("#%d failed: %s", i, err)
		}
		if !reflect.DeepEqual(test.value, value) {
			t.Errorf("#%d: Wrong UTF8 String, expected %#+v, got %#+v", i, test.value, value)
		}
		if test.offset != test.bytes.offset {
			t.Errorf("#%d: Wrong Offset, expected %#v, got %#v", i, test.offset, test.bytes.offset)
		}
	}
}

var ParseOCTETSTRINGTestData = []struct {
	bytes  *Bytes     // Input
	length int        // Input
	value  OCTETSTRING // Expected output
	offset int        // Expected output
}{
	{value: OCTETSTRING([]byte{0x41, 0x4d, 0x2d}), offset: 5, length: 3, bytes: &Bytes{offset: 2, bytes: []byte{0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35}}},
	{value: OCTETSTRING([]byte{0xe4, 0xb8, 0x96}), offset: 10, length: 3, bytes: &Bytes{offset: 7, bytes: []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}}},
	{value: OCTETSTRING([]byte("myLogin")), offset: 19, length: 7, bytes: &Bytes{offset: 12, bytes: []byte{0x30, 0x1d, 0x02, 0x01, 0x05, 0x60, 0x18, 0x02, 0x01, 0x03, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64}}},
}

func TestParseOCTETSTRING(t *testing.T) {
	for i, test := range ParseOCTETSTRINGTestData {
		value, err := test.bytes.ParseOCTETSTRING(test.length)
		if err != nil {
			t.Errorf("#%d failed: %s", i, err)
		}
		if !reflect.DeepEqual(test.value, value) {
			t.Errorf("#%d: Wrong UTF8 String, expected %#+v, got %#+v", i, test.value, value)
		}
		if test.offset != test.bytes.offset {
			t.Errorf("#%d: Wrong Offset, expected %#v, got %#v", i, test.offset, test.bytes.offset)
		}
	}
}
