package fixedXOR

import (
	"testing"
)

type fixedXORTest struct {
	a   string
	b   string
	out string
}

var fixedXORTests = []fixedXORTest{
	{
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965",
		"746865206b696420646f6e277420706c6179",
	},
}

func TestFixedXOR(t *testing.T) {
	for i, test := range fixedXORTests {
		out := FixedXORString(test.a, test.b)
		if out != test.out {
			t.Errorf("#%d: got: %#v expected %#v", i, out, test.out)
		}
	}
}
