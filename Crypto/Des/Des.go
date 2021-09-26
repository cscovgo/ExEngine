package Des

import (
	"ExEngine/Crypto"
	"crypto/cipher"
	"crypto/des"
)

type Des struct {
	key     []byte
	iv      []byte
	padding int
	mode    int
}

func New(key, iv string, padding, mode int) *Des {
	d := &Des{
		key:     []byte(key),
		iv:      []byte(iv),
		padding: padding,
		mode:    mode,
	}
	kl := len(d.key)
	ivl := len(d.iv)
	if kl != 8 && kl != 16 && kl != 24 && kl != 32 {
		panic("DES key length must 8/16/24/32.")
	}
	if ivl != 8 && ivl != 16 && ivl != 24 && ivl != 32 {
		panic("DES iv length must 8/16/24/32.")
	}
	return d
}

func (d *Des) cipher(block cipher.Block, de bool) (cipher.BlockMode, cipher.Stream) {
	if de {
		switch d.mode {
		case Crypto.ECB:
			return Crypto.NewECBDecrypter(block), nil
		case Crypto.CBC:
			return cipher.NewCBCDecrypter(block, d.iv), nil
		case Crypto.CTR:
			return nil, cipher.NewCTR(block, d.iv)
		case Crypto.OFB:
			return nil, cipher.NewOFB(block, d.iv)
		case Crypto.CFB:
			return nil, cipher.NewCFBDecrypter(block, d.iv)
		default:
			return nil, nil
		}
	} else {
		switch d.mode {
		case Crypto.ECB:
			return Crypto.NewECBEncrypter(block), nil
		case Crypto.CBC:
			return cipher.NewCBCEncrypter(block, d.iv), nil
		case Crypto.CTR:
			return nil, cipher.NewCTR(block, d.iv)
		case Crypto.OFB:
			return nil, cipher.NewOFB(block, d.iv)
		case Crypto.CFB:
			return nil, cipher.NewCFBEncrypter(block, d.iv)
		default:
			return nil, nil
		}
	}
}

func (d *Des) Encrypt(origin string) *Ret {
	b := []byte(origin)
	block, err := des.NewCipher(d.key)
	if err != nil {
		return &Ret{err: err}
	}
	padded := Crypto.Padding(b, d.padding, block.BlockSize())
	en := make([]byte, len(padded))
	bm, s := d.cipher(block, false)
	if s != nil {
		s.XORKeyStream(en, padded)
		return &Ret{v: en}
	}
	bm.CryptBlocks(en, padded)
	return &Ret{v: en}
}

func (d *Des) Decrypt(origin []byte) *Ret {
	block, err := des.NewCipher(d.key)
	if err != nil {
		return &Ret{err: err}
	}
	de := make([]byte, len(origin))
	bm, s := d.cipher(block, true)
	if s != nil {
		s.XORKeyStream(de, origin)
	} else {
		bm.CryptBlocks(de, origin)
	}
	unPadded := Crypto.UnPadding(de, d.padding)
	return &Ret{v: unPadded}
}
