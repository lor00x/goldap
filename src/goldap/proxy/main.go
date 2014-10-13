package main

import (
	"goldap/proxy/server"
)

func main() {
	server.Forward(":2389", "127.0.0.1:10389")
}
