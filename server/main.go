package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 6666")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			go handlisten(conn)
		}
	}
}

func packetSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == 0x123456 {
		var l int32
		binary.Read(bytes.NewReader(data[4:8]), binary.BigEndian, &l)
		pl := int(l) + 8
		fmt.Println(data[0 : pl+1])
		if pl <= len(data) {
			return pl, data[:pl], nil
		}
	}
	return
}

func handlisten(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("conn 关闭")

	result := bytes.NewBuffer(nil)
	var buf [65542]byte
	for {
		n, err := conn.Read(buf[:])
		result.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		} else {
			s := bufio.NewScanner(result)
			s.Split(packetSplit)
			for s.Scan() {
				fmt.Println(string(s.Bytes()[6:]))
			}

		}
		result.Reset()
	}
}
