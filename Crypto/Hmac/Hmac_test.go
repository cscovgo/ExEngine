package Hmac_test

import (
	"ExEngine/Crypto/Hmac"
	"fmt"
	"testing"
)

func TestSum1(t *testing.T) {
	fmt.Println(Hmac.Sum1([]byte("a"), []byte("a")).Hex())
}
func TestSum256(t *testing.T) {
	fmt.Println(Hmac.Sum256([]byte("a"), []byte("a")).Hex())
}
func TestSum512(t *testing.T) {
	fmt.Println(Hmac.Sum512([]byte("a"), []byte("a")).Hex())
}
