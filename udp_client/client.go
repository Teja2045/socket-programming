package main

import (
	"log"
	"log/slog"
	"net"
	"time"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8001")
	handleError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	handleError(err)

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		_, err = conn.Write([]byte("hello udp server"))
		handleError(err)

		n, err := conn.Read(buf)
		handleError(err)

		slog.Info(string(buf[:n]))
		time.Sleep(time.Second)
	}
}
