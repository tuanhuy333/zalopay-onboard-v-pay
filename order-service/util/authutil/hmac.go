package authutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// BuildMAC convert params to slice of string and join with sep '|',
// the build MAC with SHA256 algorithm, the output is a hexadecimal encoding
func BuildMAC(key string, params ...interface{}) string {
	var str []string
	for _, p := range params {
		str = append(str, fmt.Sprint(p))
	}

	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(strings.Join(str, "|")))
	s := strings.Join(str, "|")
	fmt.Println(s)
	return hex.EncodeToString(h.Sum(nil))
}

// ValidMAC check the provided mac is equal with the result of BuildMAC
func ValidMAC(key string, mac string, params ...interface{}) bool {
	return BuildMAC(key, params...) == mac
}
