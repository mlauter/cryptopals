package hex2base64

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
)

// Takes in a hex encoded string and returns that string base64 encoded
// Or empty string and error
// Uses my manual implementation of encoding/decoding
// rather than the standard library (well, halfway anyway
// because I didn't implement base64 encoding)
func Hex2Base64Handmade(s string) string {
	decoded, err := hexDecodeStr(s)
	if err != nil {
		log.Fatal(err)
	}
	return base64Encode(decoded)
}

// Takes in a hex encoded string and returns that string base64 encoded
// Or empty string and error
func Hex2Base64(s string) string {
	// The std library way, for future reference
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(decoded)
}

// This implementation is just to prove to myself I can do it
// I'm probably not thinking of all errors I'd need to handle,
// And I'll use go's encoding library for the rest of the
// exercises
// the go source code has this nicer, separating Decode
// and DecodeString, where Decode takes a dst and src
// and modifies the injected dst. Also does a cool thing where it
// actually uses the src array as dst, overwriting as it goes,
// and then returning the relevant part i think
func hexDecodeStr(s string) ([]byte, error) {
	src := []byte(s)
	dst := make([]byte, len(src)/2)
	if len(s)%2 != 0 {
		return dst, errors.New("Not a valid hex string")
	}

	for i := range dst {
		first, ok := unhex(src[i*2])
		if !ok {
			return dst, errors.New("Invalid hex char") // Would def be nicer to have the byte in the message but w/e
		}
		second, ok := unhex(src[i*2+1])
		if !ok {
			return dst, errors.New("Invalid hex char")
		}

		dst[i] = first<<4 | second // equivalent to first * 16 + second
	}

	return dst, nil
}

// Returns the value of the character from
// 0-15 and true if the character was valid
// otherwise 0 and an error
// this comes from the go docs lol
func unhex(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		// why c - '0'? Because this is the value of the ascii char c,
		// not c's value, the ascii numbers start with 0 at 048,
		// and go up sequentially, so subtracting the character '0'
		// from c will give us the numeric value of c
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}

	return 0, false
}

// Hrm okay turns out base64 is slightly more complicated
// let's assume standard padding character `=`
func base64Encode(src []byte) string {
	buf := make([]byte, base64EncodedLen(len(src)))

	// mehhhhh actually base64 encoding is a bit more involved
	// and it feels like not a good use of time to implement it myself right now,
	// I should come back to it. Basically you take 3 bytes of input,
	// and then break it into 4 groups of 6 bits, and then encode that
	// to produce 4 bytes of base64
	base64.StdEncoding.Encode(buf, src)
	return string(buf)
}

// given an input of n bytes, the output will be 4[n/3] bytes long,
// including padding characters.
// (you'll have to pad up to 2 bytes)
func base64EncodedLen(n int) int {
	return ((n + 2) / 3) * 4 // integer division
}
