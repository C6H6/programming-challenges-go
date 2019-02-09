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
	var hostPtr = flag.String("host", "0.pl.pool.ntp.org", "Server address")
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

	seconds, fraction := parseResponse(buf)
	timestamp, timestampNano := getParsedTimestamps(seconds, fraction)

	date := time.Unix(int64(timestamp), int64(timestampNano))
	fmt.Println(date)
}

func getRequest() []byte {
	var request = make([]byte, 48, 48)
	request[0] = 0x1B // 00, 011, 011 = No warning, IPv4 only, Client
	return request
}

func parseResponse(res []byte) ([]byte, []byte) {
	seconds := res[40:44]
	fraction := res[44:48]

	return seconds, fraction
}

func getParsedTimestamps(seconds []byte, fraction []byte) (uint32, uint32) {
	return binary.BigEndian.Uint32(seconds) - 2208988800, binary.BigEndian.Uint32(fraction)
}
