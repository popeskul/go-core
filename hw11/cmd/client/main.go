package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		return
	}
	defer conn.Close()

	go handler(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			return
		}

		if text == "exit\n" {
			fmt.Println("Bye!")
			return
		}
	}
}

// Reads responses from the server and prints them to the console
func handler(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		fmt.Print(string(msg), "\n", "> ")
	}
}
