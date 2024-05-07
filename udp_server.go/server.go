package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8001")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()
	slog.Info("Listening on port 8001")

	buf := make([]byte, 1024)
	for {
		n, addr, err := ln.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		slog.Info(fmt.Sprintf("Received %s from %s \n", string(buf[:n]), addr))
		_, err = ln.WriteToUDP([]byte(time.Now().String()), addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
