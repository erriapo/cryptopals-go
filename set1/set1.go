package set1

import (
	"encoding/hex"
	"log"
	//"fmt"
	//"reflect"
	"encoding/base64"
)

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
