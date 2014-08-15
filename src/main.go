package main

import (
	"log"
	"net"
	"fmt"
)

func main() {
	ln, err := net.Listen("tcp", ":2389")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on port 2389...")

	msgchan := make(chan []byte)
	go printMessages(msgchan)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("New connection accepted")
		go handleConnection(conn, msgchan)
	}
}

func printMessages(msgchan <-chan []byte) {
	for msg := range msgchan {
		result := ""
		for _, onebyte := range msg {
			if onebyte < 0x10 {
				result =fmt.Sprintf("%s, 0x0%x", result, onebyte)
			} else {
				result =fmt.Sprintf("%s, 0x%x", result, onebyte)
			} 
		}
		log.Printf("new message: %s", result)
	}
}

func handleConnection(c net.Conn, msgchan chan<- []byte) {
	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		msgchan <- buf[0:n]
		// ...
	}
}

//import (
//	"fmt"
//	"log"
//	"net"
//	//    "runtime"
//	"bufio"
//	//    "io"
//	"time"
//	//    "strconv"
//	//    "strings"
//	//    "io/ioutil"
//	"encoding/asn1"
//	//"reflect"
//)
//
//const PORT = "389"
//
//var ldapServerConfig = map[string]string{
//	"LDAP_SERVER_LISTEN": "0.0.0.0:2389",
//	"LDAP_SERVER_NAME":   "concertino",
//}
//
//var sem chan int // currently active clients
//var timeout time.Duration
//
//func main() {
//	timeout = time.Duration(10)
//	// Start listening for LDAP connections
//	listener, err := net.Listen("tcp", ldapServerConfig["LDAP_SERVER_LISTEN"])
//	if err != nil {
//		log.Fatalf(fmt.Sprintf("Cannot listen on port, %v", err))
//	} else {
//		log.Println(fmt.Sprintf("Listening on tcp %s", ldapServerConfig["LDAP_SERVER_LISTEN"]))
//	}
//	// var clientId int64
//	//clientId = 1
//	for {
//
//		conn, err := listener.Accept()
//		if err != nil {
//			log.Println(fmt.Sprintf("Accept error: %s", err))
//			continue
//		}
//		bufin := bufio.NewReader(conn)
//		// buffered := bufin.Buffered()
//		bytes := make([]byte, 4)
//		log.Println("==================================")
//		// var bytes [4]byte
//		for i := 0; i < 10000000; i++ {
//			b, err := bufin.ReadByte()
//			if err != nil {
//				log.Println(fmt.Sprintf("Accept error: %s", err))
//				break
//			}
//			// log.Printf("%0x (%d)\n", b, b)
//			fmt.Printf("0x")
//			if b < 16 {
//				fmt.Printf("0")
//			}
//			fmt.Printf("%x,", b)
//			bytes[0] = b
//		}
//		fmt.Println("")
//		log.Println("Message id: ", bytes)
//		var messageid int32 = 0
//		rest, parseerror := asn1.Unmarshal(bytes, &messageid)
//		log.Println("Message id parsed: ", messageid)
//		log.Println("Rest: ", rest, ", error: ", parseerror)
//
//		//		log.Println(fmt.Sprintf(" There are now "+strconv.Itoa(runtime.NumGoroutine())+" serving goroutines"))
//		//		sem <- 1 // Wait for active queue to drain.
//		//		go handleClient(&Client{
//		//			conn:        conn,
//		//			address:     conn.RemoteAddr().String(),
//		//			time:        time.Now().Unix(),
//		//			bufin:       bufio.NewReader(conn),
//		//			bufout:      bufio.NewWriter(conn),
//		//			clientId:    clientId,
//		//			savedNotify: make(chan int),
//		//		})
//		//		clientId++
//	}
//}

//type LdapPacket struct {
//	MessageID	int
//	ProtocolOp	int
//}

//func newInt(n int) *int { return &n }
//
//func main3() {
//	var in = []byte{0x02, 0x01, 0x42}
//	var expected = newInt(0x42)
//	pv := reflect.New(reflect.TypeOf(expected).Elem())
//	val := pv.Interface()
//	val = interface{}(expected)
//	var rest []byte
//	var err error
//	rest, err = asn1.Unmarshal(in, val)
//	fmt.Println("IN: ", in)
//	fmt.Println("OUT: ", reflect.Indirect(reflect.ValueOf(val)).Int())
//	fmt.Printf("EXP: %#v\n", reflect.Indirect(reflect.ValueOf(expected)).Int())
//	fmt.Println("OUT RAW: ", *val.(*int))
//	fmt.Println("EXP RAW: ", *expected)
//	fmt.Println("RST: ", rest)
//	fmt.Println("ERR: ", err)
//
//	if reflect.DeepEqual(val, expected) {
//		fmt.Println("OK !")
//	} else {
//		fmt.Println("KO !")
//	}
//
//}
//
//func main4() {
//	// var in = []byte {0x02, 0x01, 0x42}
//	var in = []byte{0x30, 0x09, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02, 0x02, 0x01, 0x03}
//
//	var expected = &[]int{0, 0, 0}
//	pv := reflect.New(reflect.TypeOf(expected).Elem())
//	val := pv.Interface()
//	val = interface{}(expected)
//	var rest []byte
//	var err error
//	rest, err = asn1.Unmarshal(in, val)
//	fmt.Println("IN: ", in)
//	// fmt.Println("OUT: ", reflect.Indirect(reflect.ValueOf(val)).Int())
//	//fmt.Printf("EXP: %#v\n", reflect.Indirect(reflect.ValueOf(expected)).Int())
//	fmt.Println("OUT RAW: ", val.(*[]int))
//	fmt.Println("EXP RAW: ", expected)
//	fmt.Println("RST: ", rest)
//	fmt.Println("ERR: ", err)
//
//	if reflect.DeepEqual(val, expected) {
//		fmt.Println("OK !")
//	} else {
//		fmt.Println("KO !")
//	}
//
//}
//
//func decodeBytes(input []byte, output interface{}) {
//	var rest []byte
//	var err error
//	rest, err = asn1.Unmarshal(input, output)
//	if len(rest) > 0 {
//		fmt.Println("Remaining bytes to read ", rest)
//	}
//	if err != nil {
//		fmt.Println("ERROR: ", err)
//		// panic(err)
//	}
//}
//
//func main5() {
//	var input = []byte{0x30, 0x08, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02, 0x02, 0x01, 0x03}
//	var output = &[]int{}
//	decodeBytes(input, output)
//	fmt.Println("IN: ", input)
//	fmt.Println("OUT:", output)
//}

//type BindRequest struct {
//	Operation	int
//	Login	string
//	Password	string
//}

//type TestContextSpecificTags struct {
//	A int `asn1:"application,tag:1"`
//	B int
//}

/*
type LdapMessage struct {
	MessageID	int
	Tag			int `asn1:"application,tag:0x60"`
//	Login		string
//	Password	string
}*/

//type BindRequest struct {
//	MessageId  int
//	ProtocolOp struct {
//		Numero   int
//		Login    []byte
//		Password []byte `asn1:"tag:0"`
//	} `asn1:"application,tag:0"`
//	Controls []interface{} `asn1:"tag:0,optional"`
//}
//
//type BindRequest2 struct {
//	MessageId  int
//	ProtocolOp struct {
//	} `asn1:"application,tag:255"`
//	Controls []interface{} `asn1:"tag:0,optional"`
//}
//
//type LdapMessage struct {
//	Bytes  []byte
//	Offset int
//}
//
//func (b *LdapMessage) Init(bytes []byte) {
//
//}

// An Ldap Message starts with:
// - byte = 0x30 (SEQUENCE)
// - the message length
//
// - the applicative tag

//func main8() {
//	//	var input = []byte {0x30, 0x0c, 0x02, 0x01, 0x01, 0x60, 0x07, 0x02, 0x01, 0x03, 0x04, 0x00, 0x80, 0x00}
//	// var input = []byte {0x30, 0x04, 0x02, 0x01, 0x01, 0x60, 0x03, 0x02, 0x01, 0x01}
//	// var input = []byte {0x30, 0x1d, 0x02, 0x01, 0x01, 0x60, 0x18, 0x02, 0x01, 0x02, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64}
//	//var input = []byte {0x60, 0x18, 0x02, 0x01, 0x02, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64}
//	//var input = []byte {0x30, 0x1d, 0x02, 0x01, 0x01, 0x60, 0x03, 0x02, 0x01, 0x01}
//
//	// var input = []byte {0x30,0x1d,0x02,0x01,0x01,0x60,0x18,0x02,0x01,0x03,0x04,0x07,0x6d,0x79,0x4c,0x6f,0x67,0x69,0x6e,0x80,0x0a,0x6d,0x79,0x50,0x61,0x73,0x73,0x77,0x6f,0x72,0x64}
//	var input = []byte{0x30, 0x1d, 0x02, 0x01, 0x05, 0x60, 0x18, 0x02, 0x01, 0x03, 0x04, 0x07, 0x6d, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x80, 0x0a, 0x6d, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64}
//
//	//var input = []byte{0x30, 0x08, 0x60, 0x03, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02};
//	// var message = &TestContextSpecificTags{}
//
//	var message BindRequest2
//
//	decodeBytes(input, &message)
//	fmt.Println("IN: ", input)
//	// fmt.Println("OUT:", message)
//	fmt.Println("Message: ", message)
//	//	fmt.Println("Login: ", message.Login)
//	//	fmt.Println("Password: ", message.Password)
//}

//func handleClient(client *Client) {
//	defer closeClient(client)
//	//	defer closeClient(client)
//	greeting := "220 " + ldapServerConfig["LDAP_SERVER_NAME"] +
//		" SMTP Guerrilla-SMTPd #" + strconv.FormatInt(client.clientId, 10) + " (" + strconv.Itoa(len(sem)) + ") " + time.Now().Format(time.RFC1123Z)
//	advertiseTls := "250-STARTTLS\r\n"
//	for i := 0; i < 100; i++ {
//		switch client.state {
//		case 0:
//			responseAdd(client, greeting)
//			client.state = 1
//		case 1:
//			input, err := readSmtp(client)
//			if err != nil {
//				log.Println(fmt.Sprintf("Read error: %v", err))
//				if err == io.EOF {
//					// client closed the connection already
//					return
//				}
//				if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
//					// too slow, timeout
//					return
//				}
//				break
//			}
//			input = strings.Trim(input, " \n\r")
//			cmd := strings.ToUpper(input)
//			switch {
//			case strings.Index(cmd, "HELO") == 0:
////				if len(input) > 5 {
////					client.helo = input[5:]
////				}
//				responseAdd(client, "250 "+ldapServerConfig["LDAP_SERVER_NAME"]+" Hello ")
//
//			default:
//				responseAdd(client, fmt.Sprintf("500 unrecognized command"))
//				client.errors++
//				if client.errors > 3 {
//					responseAdd(client, fmt.Sprintf("500 Too many unrecognized commands"))
//					killClient(client)
//				}
//			}
//		}
//		// Send a response back to the client
//		err := responseWrite(client)
//		if err != nil {
//			if err == io.EOF {
//				// client closed the connection already
//				return
//			}
//			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
//				// too slow, timeout
//				return
//			}
//		}
//		if client.kill_time > 1 {
//			return
//		}
//	}
//}
//
//type Client struct {
//	state       int
////	helo        string
////	mail_from   string
////	rcpt_to     string
////	read_buffer string
//	response    string
//	address     string
////	data        string
////	subject     string
////	hash        string
//	time        int64
////	tls_on      bool
//	conn        net.Conn
//	bufin       *bufio.Reader
//	bufout      *bufio.Writer
//	kill_time   int64
//	errors      int
//	clientId    int64
//	savedNotify chan int
//}
//
//func responseAdd(client *Client, line string) {
//	client.response = line + "\r\n"
//}
//func closeClient(client *Client) {
//	client.conn.Close()
//	<-sem // Done; enable next client to run.
//}
//func killClient(client *Client) {
//	client.kill_time = time.Now().Unix()
//}
//
//func responseWrite(client *Client) (err error) {
//	var size int
//	client.conn.SetDeadline(time.Now().Add(timeout * time.Second))
//	size, err = client.bufout.WriteString(client.response)
//	client.bufout.Flush()
//	client.response = client.response[size:]
//	return err
//}
