package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	data := []byte("[zhe才是一个完整的数据包]")
	l := len(data)
	dataLen := make([]byte, 4)
	magicNum := make([]byte, 4)
	binary.BigEndian.PutUint32(magicNum, 0x123456)
	binary.BigEndian.PutUint32(dataLen, uint32(l))
	packetBuff := bytes.NewBuffer(magicNum)
	packetBuff.Write(dataLen)
	packetBuff.Write(data)
	conn, err := net.DialTimeout("tcp", "localhost:6666", time.Second*5)
	if err != nil {
		fmt.Println("connect failed,err:", err)
		return
	}
	for i := 0; i < 200; i++ {
		_, err := conn.Write(packetBuff.Bytes())
		if err != nil {
			fmt.Println("write failed,err:", err.Error())
		}
	}
}
