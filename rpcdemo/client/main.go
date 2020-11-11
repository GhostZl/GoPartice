package main

import "net"

func main() {
	conn, err := net.Dial("tcp", ":1234")
}
