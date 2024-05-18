package main

import (
	"fmt"
	"net"
)

func main() {
	l, _ := net.Listen("tcp", ":6633")

	for {
		conn, _ := l.Accept()

		go func() {
			defer conn.Close()

			var buf = make([]byte, 512)
			n, _ := conn.Read(buf)
			fmt.Printf("Read from %d byte data.\r\n", n)

			buf = make([]byte, 0)
			n, _ = conn.Write(buf)
			fmt.Printf("Write to %d byte data.\r\n", n)
		}()
	}
}
