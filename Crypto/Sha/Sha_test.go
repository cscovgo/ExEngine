package Sha_test

import (
	"ExEngine/Crypto/Sha"
	"fmt"
	"testing"
)

func TestSum1(t *testing.T) {
	fmt.Println(Sha.Sum1([]byte("a")).Hex())
}

func TestSum256(t *testing.T) {
	fmt.Println(Sha.Sum256([]byte("a")).Hex())
}

func TestSum512(t *testing.T) {
	fmt.Println(Sha.Sum512([]byte("a")).Hex())
}
