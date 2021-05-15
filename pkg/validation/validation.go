package validation

import (
	"encryption-service/pkg/decryption"
)

// Compare compares plainText and encryptedText
func Compare(plainText string, encryptedText string) bool {
	var isEqual bool
	decryptedText := decryption.Decrypt(encryptedText)
	if plainText == decryptedText {
		isEqual = true
	} else {
		isEqual = false
	}

	return isEqual
}
