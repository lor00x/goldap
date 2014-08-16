package goldap

import (
	//	"bufio"
	"errors"
	"fmt"
	"github.com/kr/pretty"
	"log"
	"net"
)

// The Proxy is a program-in-the-middle which will dump every LDAP structures
// exchanged between the client and the server
type Proxy struct {
	name       string
	dumpChan   chan Message
	clientConn net.Conn
	serverConn net.Conn
	clientChan chan Message
	serverChan chan Message
}

// To dump each request we have to read the ASN.1 first bytes to get the lengh of the message
// then build a slice of bytes with the correct amount of data
type Message struct {
	id     int
	source string
	bytes  []byte
}

func Forward(local string, remote string) {
	listener, err := net.Listen("tcp", local)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on port %s...", local)

	i := 0
	for {
		i++
		var err error
		proxy := Proxy{name: fmt.Sprintf("PROXY%d", i)}
		proxy.clientConn, err = listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		proxy.serverConn, err = net.Dial("tcp", remote)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("New connection accepted")
		go proxy.Start()
	}
}

func (p *Proxy) Start() {
	p.dumpChan = make(chan Message)
	p.clientChan = make(chan Message)
	p.serverChan = make(chan Message)
	go p.Dump()
	go p.ReadClient()
	go p.WriteServer()
	go p.ReadServer()
	go p.WriteClient()
}

func (p *Proxy) ReadClient() {
	buf := make([]byte, 1024)
	messageid := 1
	for {
		// @TODO: read the tag and length, then get the right amount of data
		n, err := p.clientConn.Read(buf)
		if err != nil || n == 0 {
			p.clientConn.Close()
			log.Printf("%s: %s", p.name, "CLIENT DISCONNECTED")
			break
		}
		messageid++
		message := Message{id: messageid, source: "CLIENT", bytes: buf[0:n]}
		p.dumpChan <- message
		p.serverChan <- message
	}
}

func (p *Proxy) WriteServer() {
	for msg := range p.serverChan {
		p.serverConn.Write(msg.bytes)
	}
}

func (p *Proxy) ReadServer() {
	buf := make([]byte, 1024*1024)
	messageid := 0
	for {
		// @TODO: read the tag and length, then get the right amount of data
		n, err := p.serverConn.Read(buf)
		if err != nil || n == 0 {
			p.serverConn.Close()
			log.Printf("%s: %s", p.name, "SERVER DISCONNECTED")
			break
		}

		messageid++
		message := Message{id: messageid, source: "SERVER", bytes: buf[0:n]}
		p.dumpChan <- message
		p.clientChan <- message
	}
}

func (p *Proxy) WriteClient() {
	for msg := range p.clientChan {
		p.clientConn.Write(msg.bytes)
	}
}

func (p *Proxy) Dump() {
	for msg := range p.dumpChan {
		result := ""
		for _, onebyte := range msg.bytes {
			if onebyte < 0x10 {
				result = fmt.Sprintf("%s, 0x0%x", result, onebyte)
			} else {
				result = fmt.Sprintf("%s, 0x%x", result, onebyte)
			}
		}
		// Now decode the message
		message, err := p.DecodeMessage(msg.bytes)
		if err != nil {
			result = fmt.Sprintf("%s\n%s", result, err.Error())
		} else {
			result = fmt.Sprintf("%s\n%# v", result, pretty.Formatter(message))
		}
		log.Printf("%s - %s - msg %d %s", p.name, msg.source, msg.id, result)
	}
}

func (p *Proxy) DecodeMessage(bytes []byte) (ret LDAPMessage, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%s", e))
		}
	}()
	ret, err = ReadLDAPMessage(&Bytes{offset: 0, bytes: bytes})
	return
}
