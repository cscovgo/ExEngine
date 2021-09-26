package Chacha20_test

import (
	"ExEngine/Crypto"
	"ExEngine/Crypto/Chacha20"
	"fmt"
	"testing"
)

func TestChacha20_XorKeyStream(t *testing.T) {
	key := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	iv := []byte("12345678")
	text := []byte("欢迎使用ExEngine")
	c := Chacha20.New(key, iv, 20)
	fmt.Println(c.XorKeyStream(text).Base64())
}

func TestChacha20_XorKeyStream2(t *testing.T) {
	key, _ := Crypto.FromHex("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	iv, _ := Crypto.FromHex("1234567898765412")
	b, _ := Crypto.FromBase64("JTyQHN3UNt1B5PzKVvPPxk82dKA=")
	c := Chacha20.New(key, iv, 8)
	fmt.Println(c.XorKeyStream(b).String())
}
