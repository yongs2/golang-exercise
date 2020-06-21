package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func AesGcmEncrypt(key []byte, plaintext string) (ciphertext, nonce []byte) {
	plaintextByte := []byte(plaintext)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce = make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintextByte, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	fmt.Printf("Nonce: %x\n", nonce)

	return
}

func AesGcmDecrypt(key, ciphertext, nonce []byte) (plaintext string) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintextByte, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	plaintext = string(plaintextByte)
	fmt.Printf("%s\n", plaintext)

	return
}

func main() {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err.Error())
	}

	plaintext := "Lorem Ipsum"
	ciphertext, nonce := AesGcmEncrypt(key, plaintext)

	plaintext = AesGcmDecrypt(key, ciphertext, nonce)
}
