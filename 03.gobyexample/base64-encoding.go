package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&{}'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(1, sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(2, string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(3, uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(4, string(uDec))
}
