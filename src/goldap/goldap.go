package goldap

import (
	"fmt"
	"goldap/asn1"
)


// SIMPLE LOGIN/PASSWORD BIND REQUEST
type BindMessage struct {
	MessageId  int
	ProtocolOp BindRequest        `asn1:"application,tag:0"`
	Controls   []interface{} `asn1:"tag:0,optional"`
}

type BindRequest struct {
	Version  int
	Login    []byte
	Password []byte `asn1:"tag:0"`
}

func (b *BindMessage) GetVersion() int {
	return b.ProtocolOp.Version
}

func (b *BindMessage) GetLogin() string {
	return string(b.ProtocolOp.Login)
}

func (b *BindMessage) GetPassword() string {
	return string(b.ProtocolOp.Password)
}

func ReadBindMessage(bytes []byte) *BindMessage {
	var bindMessage BindMessage
	_, err := asn1.Unmarshal(bytes, &bindMessage)
	if err != nil {
		fmt.Println("ERR: %s", err)
	}
	return &bindMessage
}


// SASL BIND REQUEST
type SaslBindMessage struct {
	MessageId  int
	ProtocolOp SaslBindRequest        `asn1:"application,tag:0"`
	Controls   []interface{} `asn1:"tag:0,optional"`
}

type SaslBindRequest struct {
	Version  int
	Name    []byte
	Authentication SaslCredentials `asn1:"tag:3"`
}

type SaslCredentials struct {
	Mechanism	[]byte 
	Credentials []byte	`asn1:"optional"`
}

func (b *SaslBindMessage) GetVersion() int {
	return b.ProtocolOp.Version
}

func (b *SaslBindMessage) GetName() string {
	return string(b.ProtocolOp.Name)
}

func (b *SaslBindMessage) GetMechanism() string {
	return string(b.ProtocolOp.Authentication.Mechanism)
}

func ReadSaslBindMessage(bytes []byte) *SaslBindMessage {
	var saslBindRequest SaslBindMessage
	_, err := asn1.Unmarshal(bytes, &saslBindRequest)
	if err != nil {
		fmt.Println("ERR: %s", err)
	}
	return &saslBindRequest
}
