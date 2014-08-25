package goldap

import (
	//	"bufio"
	"errors"
	"fmt"
//	"github.com/kr/pretty"
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
	messageid := 1
	for {
		var err error
		var bytes *[]byte
		bytes, err = p.readLdapMessageBytes(p.clientConn)
		if err != nil {
			p.clientConn.Close()
			log.Printf("%s: %s", p.name, "CLIENT DISCONNECTED")
			break
		}

		messageid++
		message := Message{id: messageid, source: "CLIENT", bytes: *bytes}
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
	messageid := 0
	for {
		var err error
		var bytes *[]byte
		bytes, err = p.readLdapMessageBytes(p.serverConn)
		if err != nil {
			p.serverConn.Close()
			log.Printf("%s: %s", p.name, "SERVER DISCONNECTED")
			break
		}
		messageid++
		message := Message{id: messageid, source: "SERVER", bytes: *bytes}
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
		log.Print(msg.bytes)
//		result := ""
//		for _, onebyte := range msg.bytes {
//			if onebyte < 0x10 {
//				result = fmt.Sprintf("%s, 0x0%x", result, onebyte)
//			} else {
//				result = fmt.Sprintf("%s, 0x%x", result, onebyte)
//			}
//		}
//		// Now decode the message
//		message, err := p.DecodeMessage(msg.bytes)
//		if err != nil {
//			result = fmt.Sprintf("%s\n%s", result, err.Error())
//		} else {
//			result = fmt.Sprintf("%s\n%# v", result, pretty.Formatter(message))
//		}
//		log.Printf("%s - %s - msg %d %s", p.name, msg.source, msg.id, result)
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


func (p *Proxy) readLdapMessageBytes(conn net.Conn) (ret *[]byte, err error){
	var bytes []byte
	var tagAndLength tagAndLength
	tagAndLength, err = p.readTagAndLength(conn, &bytes)
	if err != nil {
		return
	}
	p.readBytes(conn, &bytes, tagAndLength.length)
	return &bytes, err
}


// Read "length" bytes from the connection
// Append the read bytes to "bytes"
// Return the last read byte
func (p *Proxy) readBytes(conn net.Conn, bytes *[]byte, length int) (b byte, err error){
	newbytes := make([]byte, length)
	n, err := conn.Read(newbytes)
	if n != length {
		fmt.Errorf("%d bytes read instead of %d", n, length)
	} else if err != nil {
		return
	}
	*bytes = append(*bytes, newbytes...)
	b = (*bytes)[len(*bytes)-1]
	return
}

// parseTagAndLength parses an ASN.1 tag and length pair from a live connection
// into a byte slice. It returns the parsed data and the new offset. SET and
// SET OF (tag 17) are mapped to SEQUENCE and SEQUENCE OF (tag 16) since we
// don't distinguish between ordered and unordered objects in this code.
func (p *Proxy) readTagAndLength(conn net.Conn, bytes *[]byte) (ret tagAndLength, err error) {
	// offset = initOffset
	//b := bytes[offset]
	//offset++
	var b byte
	b, err = p.readBytes(conn, bytes, 1)
	if err != nil {
		return
	}
	ret.class = int(b >> 6)
	ret.isCompound = b&0x20 == 0x20
	ret.tag = int(b & 0x1f)

//	// If the bottom five bits are set, then the tag number is actually base 128
//	// encoded afterwards
//	if ret.tag == 0x1f {
//		ret.tag, err = parseBase128Int(conn, bytes)
//		if err != nil {
//			return
//		}
//	}
	// We are expecting the LDAP sequence tag 0x30 as first byte
	if b != 0x30 {
		panic(fmt.Sprintf("Expecting 0x30 as first byte, but got %#x instead", b))
	}


	b, err = p.readBytes(conn, bytes, 1)
	if err != nil {
		return
	}
	if b&0x80 == 0 {
		// The length is encoded in the bottom 7 bits.
		ret.length = int(b & 0x7f)
	} else {
		// Bottom 7 bits give the number of length bytes to follow.
		numBytes := int(b & 0x7f)
		if numBytes == 0 {
			err = SyntaxError{"indefinite length found (not DER)"}
			return
		}
		ret.length = 0
		for i := 0; i < numBytes; i++ {

			b, err = p.readBytes(conn, bytes, 1)
			if err != nil {
				return
			}
			if ret.length >= 1<<23 {
				// We can't shift ret.length up without
				// overflowing.
				err = StructuralError{"length too large"}
				return
			}
			ret.length <<= 8
			ret.length |= int(b)
			if ret.length == 0 {
				// DER requires that lengths be minimal.
				err = StructuralError{"superfluous leading zeros in length"}
				return
			}
		}
	}

	return
}
