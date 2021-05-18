package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"

	config "github.com/Colris-Inc/auth/configurations"
)

var (
	encryptionKey = config.AppConfig.EncryptionKey
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
	cipherData := make([]byte, aes.BlockSize+len(unencData))

	// initialization vector
	iv := cipherData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(cipherData[aes.BlockSize:], unencData)

	// authenticate the cipher text
	hash := hmac.New(sha256.New, encKey)

	io.WriteString(hash, string(cipherData))

	// return authenticated cipher
	return string(hash.Sum(nil))
}

func DecryptString(encryptedString string) string {
	var response string

	return response
}
