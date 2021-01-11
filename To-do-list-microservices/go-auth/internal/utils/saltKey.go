package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"hash"
)

const (
	SaltSize   = 64
	iterations = 1e4
)

func ComparePassword(hash, pw []byte) bool {
	return bytes.Equal(hash, HashPassword(pw, hash[:SaltSize]))
}

// hash the password with the provided salt using the pbkdf2 algorithm
// return byte slice containing salt (first 64 bytes) and hash (last 32 bytes) => total of 96 bytes
func HashPassword(pw, salt []byte) []byte {
	ret := make([]byte, len(salt))
	copy(ret, salt)
	return append(ret, Key(pw, salt, iterations, sha256.Size, sha256.New)...)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func Key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	prf := hmac.New(h, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen

	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}
