package main

import (
	"fmt"
	"github.com/mlauter/cryptopals/set1/c1/hex2base64"
)

const (
	s1C1HexStr = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
)

func main() {
	// s1_c1
	base64Str := hex2base64.Hex2Base64Handmade(s1C1HexStr)
	sBase64Str := hex2base64.Hex2Base64(s1C1HexStr)
	fmt.Printf("Set1 - Challenge1:\n\tmine: %s\n\tstd:  %s\n", base64Str, sBase64Str)
}
