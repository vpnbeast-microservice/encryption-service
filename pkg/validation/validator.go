package validation

import (
	"encryption-service/pkg/decryption"
)

func Compare(plainText string, encryptedText string) bool {
	decryptedText := decryption.Decrypt(encryptedText)
	if plainText == decryptedText {
		return true
	}
	return false
}