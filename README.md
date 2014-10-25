[![GoDoc](https://godoc.org/github.com/lor00x/goldap?status.svg)](https://godoc.org/github.com/lor00x/goldap)
[![Build Status](https://travis-ci.org/lor00x/goldap.svg)](https://travis-ci.org/lor00x/goldap)

goldap
======

Implementation of the LDAP protocol in Go

This is experimental, I'm learning Golang ;-)

I'm co-learning Golang with [Valère JEANTET](https://github.com/vjeantet). Have a look at its own implementation of LDAP in Golang: https://github.com/vjeantet/ldapserver

Feel free to contribute, comment :)

# Installation

```bash
go get github.com/kr/pretty
go get github.com/kr/pty
go get github.com/lor00x/goldap
```

# How to run the LDAP proxy

```bash
cd proxy
go run main.go
```

By default the proxy will listen on the port 2389 on all interfaces and forward to 127.0.0.1:10389.
This is the configuration I use for testing with Apache Directory Studio.

You can change the configuration by editing the file "proxy/main.go"

Once the proxy is running, you still need an LDAP proxy and client for testing.

# Quickly setup an LDAP server and client for testing

For testing I use Apache Directory Studio, you can download it here : http://directory.apache.org/studio/downloads.html

## Setup the LDAP server

* Launch Apache Directory Studio.
* In the lower left you should see the "LDAP Servers" panel (if not click on Window => Show view => LDAP Servers).
* Click on the "New server" button
* Select "ApacheDS 2.0.0" and click on finish
* Double-click on the new server in the list, this should open the configuration in the central panel
* Set 10389 as the LDAP server port
* In the "Partitions" section, you should see an example dataset, this is perfect for our tests :-)
* Save the configuration, select the server and click on the "Run" button

## Setup the LDAP client

* Launch Apache Directory Studio
* In the lower left you should see the "Connections" panel (if not click on Window => Show view => Connections)
* Click on the "New connection" button
* Enter a name, then setup 127.0.0.1 for the hostname and 10389 for the port
* Click on next, select "No Authentication" as authentication method
* Click on finish, the connection should open automatically (if not, double click on it)
* The content of the LDAP server should be displayed in the "LDAP Browser" panel
* Now return to the console where you previously launched the goldap proxy, you should see the LDAP messages dumped as you browse the server :-)


# Output example:

```bash
$ cd goldap/proxy
$ go run main.go
2014/10/25 17:53:52 Listening on port :2389...
2014/10/25 17:54:07 New connection accepted
2014/10/25 17:54:07 Message: PROXY1 - CLIENT - msg 2 
message.LDAPMessage{
    messageID:  1,
    protocolOp: message.BindRequest{
        version:        3,
        name:           "",
        authentication: "",
    },
    controls: (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - SERVER - msg 1 
message.LDAPMessage{
    messageID:  1,
    protocolOp: message.BindResponse{},
    controls:   (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - CLIENT - msg 3 
message.LDAPMessage{
    messageID:  2,
    protocolOp: message.SearchRequest{
        baseObject:   "",
        scope:        0,
        derefAliases: 3,
        sizeLimit:    0,
        timeLimit:    0,
        typesOnly:    false,
        filter:       "objectClass",
        attributes:   {"subschemaSubentry"},
    },
    controls: (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - SERVER - msg 2 
message.LDAPMessage{
    messageID:  2,
    protocolOp: message.SearchResultEntry{
        objectName: "",
        attributes: {
            {
                type_: "subschemaSubentry",
                vals:  {"cn=schema"},
            },
        },
    },
    controls: (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - SERVER - msg 3 
message.LDAPMessage{
    messageID:  2,
    protocolOp: message.SearchResultDone{},
    controls:   (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - CLIENT - msg 4 
message.LDAPMessage{
    messageID:  3,
    protocolOp: message.SearchRequest{
        baseObject:   "cn=schema",
        scope:        0,
        derefAliases: 3,
        sizeLimit:    0,
        timeLimit:    0,
        typesOnly:    false,
        filter:       message.FilterEqualityMatch{attributeDesc:"objectClass", assertionValue:"subschema"},
        attributes:   {"createTimestamp", "modifyTimestamp"},
    },
    controls: (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - SERVER - msg 4 
message.LDAPMessage{
    messageID:  3,
    protocolOp: message.SearchResultEntry{
        objectName: "cn=schema",
        attributes: {
            {
                type_: "modifyTimestamp",
                vals:  {"20090818022733Z"},
            },
            {
                type_: "createTimestamp",
                vals:  {"20090818022733Z"},
            },
        },
    },
    controls: (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - SERVER - msg 5 
message.LDAPMessage{
    messageID:  3,
    protocolOp: message.SearchResultDone{},
    controls:   (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - CLIENT - msg 5 
message.LDAPMessage{
    messageID:  4,
    protocolOp: message.SearchRequest{
        baseObject:   "",
        scope:        0,
        derefAliases: 0,
        sizeLimit:    0,
        timeLimit:    0,
        typesOnly:    false,
        filter:       "objectClass",
        attributes:   {"namingContexts", "subschemaSubentry", "supportedLDAPVersion", "supportedSASLMechanisms", "supportedExtension", "supportedControl", "supportedFeatures", "vendorName", "vendorVersion", "+", "objectClass"},
    },
    controls: (*message.Controls)(nil),
}

2014/10/25 17:54:07 Message: PROXY1 - SERVER - msg 6 
message.LDAPMessage{
    messageID:  4,
    protocolOp: message.SearchResultEntry{
        objectName: "",
        attributes: {
            {
                type_: "vendorName",
                vals:  {"Apache Software Foundation"},
            },
            {
                type_: "vendorVersion",
                vals:  {"2.0.0-M14"},
            },
            {
                type_: "objectClass",
                vals:  {"top", "extensibleObject"},
            },
            {
                type_: "subschemaSubentry",
                vals:  {"cn=schema"},
            },
            {
                type_: "supportedLDAPVersion",
                vals:  {"3"},
            },
            {
                type_: "supportedControl",
                vals:  {"2.16.840.1.113730.3.4.3", "1.3.6.1.4.1.4203.1.10.1", "2.16.840.1.113730.3.4.2", "1.3.6.1.4.1.4203.1.9.1.4", "1.3.6.1.4.1.42.2.27.8.5.1", "1.3.6.1.4.1.4203.1.9.1.1", "1.3.6.1.4.1.4203.1.9.1.3", "1.3.6.1.4.1.4203.1.9.1.2", "1.3.6.1.4.1.18060.0.0.1", "2.16.840.1.113730.3.4.7", "1.2.840.113556.1.4.319"},
            },
            {
                type_: "supportedExtension",
                vals:  {"1.3.6.1.4.1.1466.20036", "1.3.6.1.4.1.1466.20037", "1.3.6.1.4.1.18060.0.1.5", "1.3.6.1.4.1.18060.0.1.3", "1.3.6.1.4.1.4203.1.11.1"},
            },
            {
                type_: "supportedSASLMechanisms",
                vals:  {"NTLM", "GSSAPI", "GSS-SPNEGO", "CRAM-MD5", "SIMPLE", "DIGEST-MD5"},
            },
            {
                type_: "entryUUID",
                vals:  {"f290425c-8272-4e62-8a67-92b06f38dbf5"},
            },
            {
                type_: "namingContexts",
                vals:  {"ou=system", "ou=schema", "dc=example,dc=com", "ou=config"},
            },
            {
                type_: "supportedFeatures",
                vals:  {"1.3.6.1.4.1.4203.1.5.1"},
            },
        },
    },
    controls: (*message.Controls)(nil),
}

```
