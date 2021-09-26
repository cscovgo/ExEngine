package TripleDes_test

import (
	"ExEngine/Crypto"
	"ExEngine/Crypto/TripleDes"
	"fmt"
	"testing"
)

func TestTripleDes_Encrypt(t *testing.T) {
	td := TripleDes.New("aaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbb", Crypto.PaddingPKCS7, Crypto.ECB)
	fmt.Println(td.Encrypt("欢迎使用ExEngine ").Hex())
}

func TestTripleDes_Decrypt(t *testing.T) {
	td := TripleDes.New("aaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbb", Crypto.PaddingPKCS7, Crypto.ECB)
	b, _ := Crypto.FromHex("9c21ec39e3a5a3febd3686c31180fac03388c6f76a52d2f1")
	fmt.Println(td.Decrypt(b).String())
}
