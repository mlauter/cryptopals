package hex2base64

import (
	"testing"
)

type hexBase64Test struct {
	hexString    string
	base64String string
}

// Obviously ideally there would be more tests in here,
// but for now just using the example from cryptopals
var hexBase64Tests = []hexBase64Test{
	{
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
}

func TestHex2Base64Handmade(t *testing.T) {
	for i, test := range hexBase64Tests {
		base64String := Hex2Base64Handmade(test.hexString)
		if base64String != test.base64String {
			t.Errorf("#%d: got: %#v expected: %#v", i, base64String, test.base64String)
		}
	}
}

func TestHex2Base64(t *testing.T) {
	for i, test := range hexBase64Tests {
		base64String := Hex2Base64(test.hexString)
		if base64String != test.base64String {
			t.Errorf("#%d: got: %#v expected: %#v", i, base64String, test.base64String)
		}
	}

}
