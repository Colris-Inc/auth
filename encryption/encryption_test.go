package encryption

import (
	"testing"
)

func TestEncryptString(t *testing.T) {
	plainTxt := "Hello World1234"
	encryptionKey = "012345679qwertya"

	encTxt := EncryptString(plainTxt)
	decTxt := DecryptString(encTxt)

	if plainTxt != decTxt {
		t.Errorf("Input %v, got '%v'", plainTxt, decTxt)
	}
}
