package Aes

import (
	"ExEngine/Crypto"
	"crypto/aes"
	"crypto/cipher"
)

type Aes struct {
	key     []byte
	iv      []byte
	padding int
	mode    int
}

//填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func New(key, iv string, padding, mode int) *Aes {
	a := &Aes{
		key:     []byte(key),
		iv:      []byte(iv),
		padding: padding,
		mode:    mode,
	}
	l := len(a.key)
	if l != 16 && l != 24 && l != 32 {
		panic("AES key length must 16/24/32 Bits")
	}
	return a
}

func (a *Aes) cipher(block cipher.Block, de bool) (cipher.BlockMode, cipher.Stream) {
	if de {
		switch a.mode {
		case Crypto.ECB:
			return Crypto.NewECBDecrypter(block), nil
		case Crypto.CBC:
			return cipher.NewCBCDecrypter(block, a.iv), nil
		case Crypto.CTR:
			return nil, cipher.NewCTR(block, a.iv)
		case Crypto.OFB:
			return nil, cipher.NewOFB(block, a.iv)
		case Crypto.CFB:
			return nil, cipher.NewCFBDecrypter(block, a.iv)
		default:
			return nil, nil
		}
	} else {
		switch a.mode {
		case Crypto.ECB:
			return Crypto.NewECBEncrypter(block), nil
		case Crypto.CBC:
			return cipher.NewCBCEncrypter(block, a.iv), nil
		case Crypto.CTR:
			return nil, cipher.NewCTR(block, a.iv)
		case Crypto.OFB:
			return nil, cipher.NewOFB(block, a.iv)
		case Crypto.CFB:
			return nil, cipher.NewCFBEncrypter(block, a.iv)
		default:
			return nil, nil
		}
	}
}

func (a *Aes) Encrypt(origin string) *Ret {
	b := []byte(origin)
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return &Ret{err: err}
	}
	var padded []byte
	bs := block.BlockSize()
	padded = Crypto.Padding(b, a.padding, bs)
	en := make([]byte, len(padded))
	bm, s := a.cipher(block, false)
	if s != nil {
		s.XORKeyStream(en, padded)
		return &Ret{v: en}
	}
	bm.CryptBlocks(en, padded)
	return &Ret{v: en}
}

func (a *Aes) Decrypt(origin []byte) *Ret {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return &Ret{err: err}
	}
	de := make([]byte, len(origin))
	bm, s := a.cipher(block, true)
	if s != nil {
		s.XORKeyStream(de, origin)
	} else {
		bm.CryptBlocks(de, origin)
	}
	unPadded := Crypto.UnPadding(de, a.padding)
	return &Ret{v: unPadded}
}
