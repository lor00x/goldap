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
