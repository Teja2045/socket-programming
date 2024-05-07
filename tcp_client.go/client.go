package main

import (
	"bufio"
	"log"
	"log/slog"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		slog.Info("message to send is " + text)

		conn.Write([]byte(text + "\n"))

		message, _ := bufio.NewReader(conn).ReadString('\n')

		slog.Info("recieved message is: " + message)
	}
}
