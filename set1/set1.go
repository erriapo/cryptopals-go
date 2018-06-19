package set1

import (
	"encoding/hex"
	"log"
	//"fmt"
	//"reflect"
	"encoding/base64"
	"errors"
)

// ErrBufferNotSameSize is an error for mismatched buffer sizes.
var ErrBufferNotSameSize = errors.New("arguments must be same size")

// XorString implements Fixed XOR described in
// http://cryptopals.com/sets/1/challenges/2
func XorString(buffer1, buffer2 string) (string, error) {
	decoded1, err := hex.DecodeString(buffer1)
	if err != nil {
		return "", err
	}
	decoded2, err := hex.DecodeString(buffer2)
	if err != nil {
		return "", err
	}

	if len(decoded1) != len(decoded2) {
		return "", ErrBufferNotSameSize
	}

	return xor([]byte(decoded1), []byte(decoded2)), nil
}

// Xor performs the XOR operation on two equal-lngth slices
func xor(buffer1, buffer2 []byte) string {
	raw := make([]byte, len(buffer1))
	for i, num := range buffer1 {
		raw[i] = num ^ buffer2[i]
	}

	return hex.EncodeToString(raw)
}

// Base64 converts bytes represented in a hex string into a base64 representation.
func Base64(src string) string {
	decoded, err := hex.DecodeString(src)
	if err != nil {
		log.Fatal(err)
	}
	//result := fmt.Sprintf("%X", decoded)
	//fmt.Printf("%v\n", reflect.TypeOf(decoded))
	return base64.StdEncoding.EncodeToString([]byte(decoded))
}
