package main

import (
	"github.com/lor00x/goldap/proxy/server"
)

func main() {
	server.Forward(":2389", "127.0.0.1:10389")
}
