package main

import "crypto/cipher"

type ctr struct {
	b  cipher.Block
	iv []byte
}

func (x *ctr) XOR(src []byte, ctr uint64) []byte {
	iv := make([]byte, x.b.BlockSize())
	out := make([]byte, x.b.BlockSize())
	copy(out, src)
	copy(iv, x.iv)
	for i := 0; i < len(iv); i++ {
		iv[i] = iv[i] ^ byte(ctr)
		ctr = ctr >> 8
	}
	key := make([]byte, x.b.BlockSize())
	x.b.Encrypt(key, iv)
	for i := 0; i < len(out); i++ {
		out[i] = out[i] ^ key[i]
	}
	return out
}

func newCTR(block cipher.Block, iv []byte) *ctr {
	if len(iv) != block.BlockSize() {
		panic("IV length must equal block size")
	}
	return &ctr{
		b:  block,
		iv: iv[:],
	}
}
