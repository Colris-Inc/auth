package encryption

import "testing"

func TestEncryptString(t *testing.T) {
	plainTxt := "Hello World1234"
	encryptionKey = "012345679qwertya"

	encTxt := EncryptString(plainTxt)

	if encTxt == "" && len(plainTxt) != 0 {
		t.Errorf("Input %v, got '', wanted encrypted string", plainTxt)
	}
}
