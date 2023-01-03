package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

func reader(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read string from server err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func writer(conn net.Conn) {
	for {
		var input string
		fmt.Scanln(&input)
		buf := []byte(input)
		buf = append(buf, '\n')
		conn.Write(buf)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3564")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	go reader(conn)
	go writer(conn)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)
	sig := <-c
	fmt.Printf("finish with signal: %v \n", sig)
}
