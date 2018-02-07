package fixedXOR

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

// Takes in two equal length byte arrays and
// returns the result of XOR-ing a against b
// If a and b are not the same length returns error
func FixedXOR(a []byte, b []byte) ([]byte, error) {
	c := make([]byte, len(a))

	if len(a) != len(b) {
		return c, errors.New("a and b must be the same length")
	}

	for i, v := range a {
		c[i] = v ^ b[i]
	}

	return c, nil
}

// Takes in two hex encoded strings
// And returns the hex encoded result of
// XOR-ing a against b
func FixedXORString(aStr string, bStr string) string {
	a, err := hex.DecodeString(aStr)
	if err != nil {
		log.Fatal(err)
	}

	b, err := hex.DecodeString(bStr)
	if err != nil {
		log.Fatal(err)
	}

	c, err := FixedXOR(a, b)
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(c)
}

func main() {
	out := FixedXORString("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	fmt.Printf("%s\n", out)
}
