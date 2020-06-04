package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

func encrypt(key, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		plaintext = append(plaintext, make([]byte, aes.BlockSize-mod)...)
	}

	block, err := aes.NewCipher(hash(key))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}

func decrypt(key, ciphertext []byte) []byte {
	block, err := aes.NewCipher(hash(key))
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext
}

func hash(key []byte) []byte {
	salt := []byte{0x1e, 0x13, 0x90, 0x9f, 0xfc, 0xbc, 0x37, 0x2a}

	dk, err := scrypt.Key(key, salt, 1<<16, 8, 1, 32)
	if err != nil {
		panic(err)
	}

	return dk
}
