package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"io"

	config "github.com/Colris-Inc/auth/configurations"
)

var (
	encryptionKey = config.AppConfig.EncryptionKey
)

const (
	sha512Bytes = 1 << 6
	sha256Bytes = 1 << 5
)

// EncryptString - will encrypt a plain string and return an encrypted string
func EncryptString(unencrypted string) string {

	// convert to byte slice
	encKey := []byte(encryptionKey)
	unencData := []byte(unencrypted)

	// create a cipher block
	block, err := aes.NewCipher(encKey)

	if err != nil {
		panic(err)
	}

	// make cipher data
	cipherData := make([]byte, aes.BlockSize+len(unencData)+sha512Bytes)

	// initialization vector
	iv := cipherData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(cipherData[aes.BlockSize:], unencData)

	// authenticate the cipher text
	hash := hmac.New(sha512.New, encKey)

	copy(cipherData[len(cipherData)-sha512Bytes:], hash.Sum(nil))
	// io.WriteString(hash, string(cipherData))

	// return authenticated cipher
	return string(cipherData)
}

func DecryptString(encryptedString string) string {

	cipherData := []byte(encryptedString)
	encKey := []byte(encryptionKey)

	hash := hmac.New(sha512.New, encKey)

	encDataHash := cipherData[len(cipherData)-sha512Bytes:]

	if ok := hmac.Equal(hash.Sum(nil), encDataHash); !ok {
		panic("MAC does not match")
	}

	block, err := aes.NewCipher(encKey)
	if err != nil {
		panic(err)
	}

	if len(cipherData) < aes.BlockSize {
		panic("data too short")
	}
	iv := cipherData[:aes.BlockSize]
	cipherData = cipherData[aes.BlockSize : len(cipherData)-sha512Bytes]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherData, cipherData)
	return string(cipherData)
}
