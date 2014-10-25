[![GoDoc](https://godoc.org/github.com/lor00x/goldap?status.svg)](https://godoc.org/github.com/lor00x/goldap)
[![Build Status](https://travis-ci.org/lor00x/goldap.svg)](https://travis-ci.org/lor00x/goldap)

goldap
======

Implementation of the LDAP protocol in Go

This is experimental, I'm learning Golang ;-)

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
