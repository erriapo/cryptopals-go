package set1

import (
	"testing"
)

func TestChallenge1(t *testing.T) {
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got := Base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if got != expected {
		t.Errorf("Base64(%q) == %q, want %q", "4927....6f6d", got, expected)
	}
}

func TestChallenge2(t *testing.T) {
	expected := "746865206b696420646f6e277420706c6179"
	arg1 := "1c0111001f010100061a024b53535009181c"
	arg2 := "686974207468652062756c6c277320657965"
	got, e := XorString(arg1, arg2) // the kid don't play
	if e != nil {
		t.Errorf("XorString(%q, %q) threw error %q", arg1, arg2, e)
	}
	if got != expected {
		t.Errorf("XorString(%q, %q) == %q, want %q", arg1, arg2, got, expected)
	}
}
