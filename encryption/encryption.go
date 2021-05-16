package encryption

import (
	config "github.com/Colris-Inc/auth/configurations"
)

var (
	encryptionKey = config.AppConfig.EncryptionKey
)

func EncryptString(unencryptedString string) string {
	var response string

	return response
}

func DecryptString(encryptedString string) string {
	var response string

	return response
}
