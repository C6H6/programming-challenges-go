package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	var portPtr = flag.String("port", "123", "Port")
	var hostPtr = flag.String("host", "localhost", "Server address")
	flag.Parse()

	var addressBuffer bytes.Buffer
	addressBuffer.WriteString(*hostPtr)
	addressBuffer.WriteString(":")
	addressBuffer.WriteString(*portPtr)

	conn, _ := net.Dial("udp", addressBuffer.String())
	defer conn.Close()

	var req = getRequest()

	_, err := conn.Write(req)
	if err != nil {
		fmt.Println("Error while sending request")
		return
	}

	buf := make([]byte, 48)

	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error while reading response")
		return
	}

	err = conn.Close()
	if err != nil {
		fmt.Println("Error while closing connection")
		return
	}

	seconds := buf[40:44]
	fraction := buf[44:48]

	timestamp := binary.BigEndian.Uint32(seconds) - 2208988800
	timestampNano := binary.BigEndian.Uint32(fraction)

	date := time.Unix(int64(timestamp), int64(timestampNano))
	fmt.Println(date)
}

func getRequest() []byte {
	var request = make([]byte, 48, 48)
	request[0] = 0x1B // 00, 011, 011 = No warning, IPv4 only, Client
	return request
}
