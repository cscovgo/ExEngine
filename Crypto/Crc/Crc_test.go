package Crc_test

import (
	"ExEngine/Crypto/Crc"
	"fmt"
	"testing"
)

func TestCrc32Q(t *testing.T) {
	str := "测试字符串"
	fmt.Println(Crc.C32Q([]byte(str)).Hex())
}

func TestCrcIEEE(t *testing.T) {
	str := "测试字符串"
	fmt.Println(Crc.IEEE([]byte(str)).Hex())
}
