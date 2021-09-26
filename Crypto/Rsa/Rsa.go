package Rsa

import (
	"ExEngine/Crypto"
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"hash"
)

type Rsa struct {
	pubKey  []byte
	priKey  []byte
	priTyp  int
	signTyp crypto.Hash
}

func New(pubKey, priKey string, signTyp ...crypto.Hash) *Rsa {
	r := &Rsa{
		pubKey: []byte(pubKey),
		priKey: []byte(priKey),
	}
	if len(signTyp) > 0 {
		r.signTyp = signTyp[0]
	}
	if err := r.checkPri(); err != nil {
		panic(err.Error())
	}
	return r
}

func (r *Rsa) checkPri() error {
	block, _ := pem.Decode(r.priKey)
	if _, err := x509.ParseECPrivateKey(block.Bytes); err == nil {
		r.priTyp = Crypto.PirTypEC
		return nil
	}
	if _, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		r.priTyp = Crypto.PirTypPKCS1
		return nil
	}
	if _, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		r.priTyp = Crypto.PirTypPKCS8
		return nil
	}
	return errors.New("unknown private key Type")
}

func (r *Rsa) parsePriKey(block *pem.Block) (pri *rsa.PrivateKey, err error) {
	switch r.priTyp {
	case Crypto.PirTypEC:
		if pri, err = x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
			return pri, nil
		}
	case Crypto.PirTypPKCS1:
		if pri, err = x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
			return pri, nil
		}
	case Crypto.PirTypPKCS8:
		priIface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err == nil {
			return priIface.(*rsa.PrivateKey), nil
		}
	}
	return nil, err
}

func (r *Rsa) hash() hash.Hash {
	switch r.signTyp {
	case Crypto.SignTypSHA1:
		return sha1.New()
	case Crypto.SignTypSHA256:
		return sha256.New()
	case Crypto.SignTypSHA512:
		return sha512.New()
	case Crypto.SignTypMD5:
		return md5.New()
	}
	return nil
}

func (r *Rsa) Encrypt(origin []byte) *Ret {
	block, _ := pem.Decode(r.pubKey)
	if block == nil {
		return &Ret{err: errors.New("Rsa:Public Key Error")}
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return &Ret{err: err}
	}
	result, err := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), origin)
	return &Ret{result, err}
}

func (r *Rsa) Decrypt(origin []byte) *Ret {
	block, _ := pem.Decode(r.priKey)
	if block == nil {
		return &Ret{err: errors.New("Rsa:Private Key Error")}
	}
	pri, err := r.parsePriKey(block)
	if err != nil {
		return &Ret{err: err}
	}
	bs := splitWithSize(origin, (pri.N.BitLen()+7)/8)
	var buf bytes.Buffer
	for _, b := range bs {
		de, err := rsa.DecryptPKCS1v15(rand.Reader, pri, b)
		if err != nil {
			return &Ret{err: err}
		}
		buf.Write(de)
	}
	return &Ret{v: buf.Bytes(), err: nil}
}

func (r *Rsa) Sign(origin []byte) *Ret {
	h := r.hash()
	h.Write(origin)
	sum := h.Sum(nil)
	block, _ := pem.Decode(r.priKey)
	if block == nil {
		return &Ret{err: errors.New("Rsa:Private Key Error")}
	}
	pri, err := r.parsePriKey(block)
	if err != nil {
		return &Ret{err: err}
	}
	sign, err := rsa.SignPKCS1v15(rand.Reader, pri, r.signTyp, sum)
 	return &Ret{v: sign, err: err}
}

func (r *Rsa) Verify(origin, sign []byte) *Ret {
	h := r.hash()
	h.Write(origin)
	sum := h.Sum(nil)
	block, _ := pem.Decode(r.pubKey)
	if block == nil {
		return &Ret{err: errors.New("Rsa:Public Key Error")}
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return &Ret{err: err}
	}
	return &Ret{err: rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), r.signTyp, sum, sign)}
}

//根据长度切割字节
func splitWithSize(plain []byte, size int) [][]byte {
	var result [][]byte
	plainLen := len(plain)
	for i := 0; i < plainLen/size; i++ {
		result = append(result, plain[size*i:size*(i+1)])
	}
	plainMod := plainLen % size
	if plainMod > 0 {
		result = append(result, plain[plainLen-plainMod:])
	}
	return result
}
