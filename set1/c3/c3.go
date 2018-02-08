package main

import (
	"encoding/hex"
	"fmt"
	"github.com/mlauter/cryptopals/set1/c2/fixedXOR"
	"log"
	"sort"
)

var freqMap = map[string]float64{
	"e": 12.49,
	"t": 9.28,
	"a": 8.04,
	"o": 7.64,
	"i": 7.57,
	"n": 7.23,
	"s": 6.51,
	"r": 6.28,
	"h": 5.05,
	"l": 4.07,
	"d": 3.82,
	"c": 3.34,
	"u": 2.73,
	"m": 2.51,
	"f": 2.40,
	"p": 2.14,
	"g": 1.87,
	"w": 1.68,
	"y": 1.66,
	"b": 1.48,
	"v": 1.05,
	"k": 0.54,
	"x": 0.23,
	"j": 0.16,
	"q": 0.12,
	"z": 0.09,
}

type cypher struct {
	K       string
	Message string
}

func score(m []byte) float64 {
	var score float64
	for _, v := range m {
		if freq, ok := freqMap[string(v)]; ok {
			score += freq
		} else {
			score -= 5
		}
	}
	return score
}

func main() {
	src := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	encrypted := dst[:n]

	// We should try all the single byte characters
	// technically this could be 0-255, but ascii is only 0-127
	candidates := make(map[float64]cypher)
	for i := 0; i < 128; i++ {
		candidateKeyStr := make([]byte, len(encrypted))
		for j, _ := range candidateKeyStr {
			candidateKeyStr[j] = byte(i)
		}

		message, _ := fixedXOR.FixedXOR(encrypted, candidateKeyStr)
		score := score(message)

		candidates[score] = cypher{
			string(i),
			string(message),
		}
	}

	var keys []float64
	for k := range candidates {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	winningKey := keys[len(keys)-1]
	winner := candidates[winningKey]
	fmt.Printf("score: %f key: %s message: %s\n", winningKey, winner.K, winner.Message)
}
