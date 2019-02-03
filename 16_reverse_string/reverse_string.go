package main

import (
	"bytes"
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "STRING")
		return
	}

	str := os.Args[1]

	fmt.Println(reverse(str))
}

func reverse(str string) string {
	var buf bytes.Buffer

	for i := len(str)-1; i >= 0; i-- {
		buf.WriteByte(byte(str[i]))
	}

	return buf.String()
}
