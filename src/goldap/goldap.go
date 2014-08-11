package goldap

import (
//	"fmt"
//	"goldap/asn1"
)




//type GoldapMessage struct {
//	MessageId	int
//	// Choice
//	BindRequest	GoldapBindRequest	`asn1:"application,tag:0,optional"`
//	Controls	[]interface{}	`asn1:"tag:0,optional"`
//}
//
//type GoldapBindRequest struct {
//	Version  int
//	Login    []byte
//	// Choice
//	Password []byte `asn1:"tag:0,optional"`
//	Authentication GoldapSaslCredentials `asn1:"tag:3,optional"`
//}
//
//type GoldapSaslCredentials struct {
//	Mechanism	[]byte 
//	Credentials []byte	`asn1:"optional"`
//}
//
//func ReadGoldapMessage(bytes []byte) (message GoldapMessage, rest []byte, err error ) {
//	rest, err = asn1.Unmarshal(bytes, &message)
//	if err != nil {
//		fmt.Println("ERR: %s", err)
//	}
//	
//	if ! message.check() {
//	}
//	return
//}
//
//
//func (m *GoldapMessage) check () bool {
//	return true
//}



//func (r GoldapBindRequest) check() bool {
//	if r.Password != nil && r.Authentication == nil {
//		return true
//	}
//	if r.Password == nil && r.Authentication != nil {
//		return true
//	}
//	return false
//}

//// SIMPLE LOGIN/PASSWORD BIND REQUEST
//type BindMessage struct {
//	MessageId  int
//	ProtocolOp BindRequest        `asn1:"application,tag:0"`
//	Controls   []interface{} `asn1:"tag:0,optional"`
//}
//
//type BindRequest struct {
//	Version  int
//	Login    []byte
//	Password []byte `asn1:"tag:0"`
//}
//
//func (b *BindMessage) GetVersion() int {
//	return b.ProtocolOp.Version
//}
//
//func (b *BindMessage) GetLogin() string {
//	return string(b.ProtocolOp.Login)
//}
//
//func (b *BindMessage) GetPassword() string {
//	return string(b.ProtocolOp.Password)
//}
//
//func ReadBindMessage(bytes []byte) *BindMessage {
//	var bindMessage BindMessage
//	_, err := asn1.Unmarshal(bytes, &bindMessage)
//	if err != nil {
//		fmt.Println("ERR: %s", err)
//	}
//	return &bindMessage
//}
//
//
//// SASL BIND REQUEST
//type SaslBindMessage struct {
//	MessageId  int
//	ProtocolOp SaslBindRequest        `asn1:"application,tag:0"`
//	Controls   []interface{} `asn1:"tag:0,optional"`
//}
//
//type SaslBindRequest struct {
//	Version  int
//	Name    []byte
//	Authentication SaslCredentials `asn1:"tag:3"`
//}
//
//type SaslCredentials struct {
//	Mechanism	[]byte 
//	Credentials []byte	`asn1:"optional"`
//}
//
//func (b *SaslBindMessage) GetVersion() int {
//	return b.ProtocolOp.Version
//}
//
//func (b *SaslBindMessage) GetName() string {
//	return string(b.ProtocolOp.Name)
//}
//
//func (b *SaslBindMessage) GetMechanism() string {
//	return string(b.ProtocolOp.Authentication.Mechanism)
//}
//
//func ReadSaslBindMessage(bytes []byte) *SaslBindMessage {
//	var saslBindRequest SaslBindMessage
//	_, err := asn1.Unmarshal(bytes, &saslBindRequest)
//	if err != nil {
//		fmt.Println("ERR: %s", err)
//	}
//	return &saslBindRequest
//}
