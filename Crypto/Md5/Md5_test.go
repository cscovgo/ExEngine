package Md5_test

import (
	"ExEngine/Crypto/Md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(Md5.Sum([]byte("a")).Hex())
	fmt.Println(Md5.Sum([]byte("a")).Hex16())
}
