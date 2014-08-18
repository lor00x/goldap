package main

import (
	"goldap"
)

func main() {
	goldap.Forward(":2389", "127.0.0.1:10389")
}
