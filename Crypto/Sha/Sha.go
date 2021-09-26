package Sha

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func Sum1(r []byte) *Ret {
	h := sha1.New()
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}

func Sum256(r []byte) *Ret {
	h := sha256.New()
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}

func Sum512(r []byte) *Ret {
	h := sha512.New()
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}
