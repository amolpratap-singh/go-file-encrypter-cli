package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	plainText, err := os.ReadFile(source)

	if err != nil {
		panic(err.Error())
	}

	encrytKey := password
	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(encrytKey, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherText := aesgcm.Seal(nil, nonce, plainText, nil)
	cipherText = append(cipherText, nonce...)

	dstFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}

	defer dstFile.Close()

	_, err = dstFile.Write(cipherText)
	if err != nil {
		panic(err.Error())
	}
}

func Decrypt(source string, password []byte) {
	fmt.Println("Decrypt function called")
}