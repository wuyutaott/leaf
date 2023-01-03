package main

import (
	"bufio"
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
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string err:", err)
		}
		conn.Write([]byte(str))
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
