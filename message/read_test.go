package message

import (
	"reflect"
	"testing"
)

type LDAPMessageTestData struct {
	bytes Bytes
	out   LDAPMessage
}

func getLDAPMessageTestData() (ret []LDAPMessageTestData) {
	return []LDAPMessageTestData{
		// A bind request with empty login / password
		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					0x30, 0x0c,
					0x02, 0x01, 0x01, // messageID
					0x60, 0x07, // tag application (class = 1) value 0 => this is a bind request
					0x02, 0x01, 0x03, // version 3
					0x04, 0x00, // empty login
					0x80, 0x00, // empty credentials
				},
			},
			out: LDAPMessage{
				messageID: MessageID(int(0x01)),
				protocolOp: BindRequest{
					version:        0x03,
					name:           LDAPDN(""),
					authentication: OCTETSTRING([]byte("")),
				},
			},
		},
		// A bind response for a bind request with empty login / password
		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					0x30, 0x0c,
					0x02, 0x01, 0x01,
					0x61, 0x07, // tag application 1 => this a bind response
					0x0a, 0x01, 0x00, // result code 0
					0x04, 0x00, // matchedDN empty
					0x04, 0x00, // diagnosticMessage empty
				},
			},
			out: LDAPMessage{
				messageID: MessageID(int(0x01)),
				protocolOp: BindResponse{
					LDAPResult: LDAPResult{
						resultCode:        0,
						matchedDN:         LDAPDN(""),
						diagnosticMessage: LDAPString(""),
					},
				},
			},
		},

		// A bind request with a simple login / password authentication
		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					0x30, 0x1d,
					0x02, 0x01, 0x01, // messageID
					0x60, 0x18, // Application, tag 0 => this is a Bind request
					0x02, 0x01, 0x03, // Version 3
					0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, // login = myLogin
					0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, // simple authentication: myPassword
				},
			},
			out: LDAPMessage{
				messageID: MessageID(int(0x01)),
				protocolOp: BindRequest{
					version:        0x03,
					name:           LDAPDN("myLogin"),
					authentication: OCTETSTRING([]byte("myPassword")),
				},
			},
		},
		// A bind request with SASL (CRAM-MD5)
		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					0x30, 0x16,
					0x02, 0x01, 0x01, // messageID
					0x60, 0x11,
					0x02, 0x01, 0x03, // version 3
					0x04, 0x00, // no login
					0xa3, 0x0a, 0x04, 0x08, 0x43, 0x52, 0x41, 0x4d, 0x2d, 0x4d, 0x44, 0x35, // SASL mechanism "CRAM-MD5", no credentials
				},
			},
			out: LDAPMessage{
				messageID: MessageID(int(0x01)),
				protocolOp: BindRequest{
					version: 0x03,
					name:    LDAPDN(""),
					authentication: SaslCredentials{
						mechanism: LDAPString("CRAM-MD5"),
					},
				},
			},
		},
		// An unbind request
		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					0x30, 0x05,
					0x02, 0x01, 0x0b, // messageID
					0x42, 0x00, // tag application 2 => unbind request with NULL
				},
			},
			out: LDAPMessage{
				messageID:  MessageID(int(0x0b)),
				protocolOp: UnbindRequest{},
			},
		},
		// A search request

		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					0x30, 0x38,
					0x02, 0x01, 0x02,
					//        SearchRequest ::= [APPLICATION 3] SEQUENCE {
					0x63, 0x33,
					//             baseObject      LDAPDN,
					0x04, 0x00,
					//             scope           ENUMERATED {
					//                  baseObject              (0),
					//                  singleLevel             (1),
					//                  wholeSubtree            (2),
					//                  ...  },
					0x0a, 0x01, 0x00,
					//             derefAliases    ENUMERATED {
					//                  neverDerefAliases       (0),
					//                  derefInSearching        (1),
					//                  derefFindingBaseObj     (2),
					//                  derefAlways             (3) },
					0x0a, 0x01, 0x03,
					//             sizeLimit       INTEGER (0 ..  maxInt),
					0x02, 0x01, 0x00,
					//             timeLimit       INTEGER (0 ..  maxInt),
					0x02, 0x01, 0x00,
					//             typesOnly       BOOLEAN,
					0x01, 0x01, 0x00,
					//             filter          Filter,
					0x87, 0x0b, // present
					0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, // objectClass"
					//             attributes      AttributeSelection }
					0x30, 0x13,
					0x04, 0x11, 0x73, 0x75, 0x62, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x75, 0x62, 0x65, 0x6e, 0x74, 0x72, 0x79, // "subschemaSubentry"
				},
			},
			out: LDAPMessage{
				messageID: MessageID(int(0x02)),
				protocolOp: SearchRequest{
					baseObject:   LDAPDN(""),
					scope:        ENUMERATED(0),
					derefAliases: ENUMERATED(3),
					sizeLimit:    INTEGER(0),
					timeLimit:    INTEGER(0),
					typesOnly:    BOOLEAN(false),
					filter:       FilterPresent("objectClass"),
					attributes: AttributeSelection(
						[]LDAPString{
							LDAPString("subschemaSubentry"),
						},
					),
				},
			},
		},
		// Search result entry
		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					0x30, 0x2b,
					0x02, 0x01, 0x02,
					//        SearchResultEntry ::= [APPLICATION 4] SEQUENCE {
					0x64, 0x26,
					//             objectName      LDAPDN,
					0x04, 0x00,
					//             attributes      PartialAttributeList }
					//        PartialAttributeList ::= SEQUENCE OF
					0x30, 0x22,
					//                             partialAttribute PartialAttribute
					0x30, 0x20,
					//        PartialAttribute ::= SEQUENCE {
					//             type       AttributeDescription,
					0x04, 0x11,
					0x73, 0x75, 0x62, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x75, 0x62, 0x65, 0x6e, 0x74, 0x72, 0x79,
					//             vals       SET OF value AttributeValue }
					0x31, 0x0b,
					0x04, 0x09,
					0x63, 0x6e, 0x3d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
				},
			},
			out: LDAPMessage{
				messageID: MessageID(int(0x02)),
				protocolOp: SearchResultEntry{
					objectName: LDAPDN(""),
					attributes: PartialAttributeList(
						[]PartialAttribute{
							PartialAttribute{
								type_: AttributeDescription(string("subschemaSubentry")),
								vals: []AttributeValue{
									AttributeValue(string("cn=schema")),
								},
							},
						},
					),
				},
			},
		},
		// SearchResultDone
		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					0x30, 0x0c,
					0x02, 0x01, 0x02,
					0x65, 0x07,
					0x0a,
					0x01, 0x00,
					0x04, 0x00,
					0x04, 0x00,
				},
			},
			out: LDAPMessage{
				messageID: MessageID(int(0x02)),
				protocolOp: SearchResultDone{
					resultCode:        ENUMERATED(0), // 0: success
					matchedDN:         LDAPDN(""),
					diagnosticMessage: LDAPString(""),
				},
			},
		},

		// ModifyDNrequest
		{
			bytes: Bytes{
				offset: NewInt(0),
				bytes: []byte{
					// 302102010d6c1c04096f753d636f6e666967040a6f753d636f6e666967670101ff8000
					0x30, 0x21,
					0x02, 0x01, 0x0d,
					0x6c, 0x1c,
					0x04, 0x09, 0x6f, 0x75, 0x3d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
					0x04, 0x0a, 0x6f, 0x75, 0x3d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x67,
					0x01, 0x01, 0xff,
					0x80, 0x00,
				},
			},
			out: LDAPMessage{
				messageID: MessageID(13),
				protocolOp: ModifyDNRequest{
					entry:        LDAPDN("ou=config"),
					newrdn:       RelativeLDAPDN("ou=configg"),
					deleteoldrdn: BOOLEAN(true),
					newSuperior:  NewLDAPDN(LDAPDN("")),
				},
				controls: (*Controls)(nil),
			},
		},
	}
}

func NewLDAPDN(ldapdn LDAPDN) *LDAPDN {
	return &ldapdn
}

func TestReadLDAPMessage(t *testing.T) {
	for i, test := range getLDAPMessageTestData() {
		message, err := ReadLDAPMessage(test.bytes)
		if err != nil {
			t.Errorf("#%d failed reading bytes at offset %d (%s): %s", i, *test.bytes.offset, test.bytes.DumpCurrentBytes(), err)
		} else if !reflect.DeepEqual(message, test.out) {
			t.Errorf("#%d:\nhave %#+v\nwant %#+v", i, message, test.out)
		}
	}
}

//func TestWriteLDAPMessage(t *testing.T) {
//	for i, test := range getLDAPMessageTestData() {
//		message, err := ReadLDAPMessage(test.bytes)
//		if err != nil {
//			t.Errorf("#%d failed reading bytes at offset %d (%s): %s", i, test.bytes.offset, test.bytes.DumpCurrentBytes(), err)
//		}
//		var bytes []byte
//		bytes, err = WriteBytes()
//		if err != nil {
//			t.Errorf("#%d failed writing bytes: %s", err)
//		} else if !reflect.DeepEqual(bytes, test.bytes) {
//			t.Errorf("#%d:\nhave %#+v\nwant %#+v", i, bytes, test.bytes)
//		}
//	}
//}
