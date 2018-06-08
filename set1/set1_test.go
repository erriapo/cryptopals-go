package set1

import (
	"testing"
)

func TestChallenge(t *testing.T) {
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got := Base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if got != expected {
		t.Errorf("Base64(%q) == %q, want %q", "4927....6f6d", got, expected)
	}
}
