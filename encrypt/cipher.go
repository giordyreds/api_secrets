package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func Encrypt(key string, plaintext string) (string, error) {
	block, err := newCipherBlock(key)

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(key string, cipherHex string) (string, error) {
	block, err := newCipherBlock(key)

	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cipherHex)

	if err != nil {
		return "", err
	}

	// no initialization vector here
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// in this way you add some salt to ciphered text,
	// if someone finds enc key he cannot retrieve the text
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	block, err := aes.NewCipher(cipherKey)

	if err != nil {
		return nil, err
	}

	return block, nil
}
