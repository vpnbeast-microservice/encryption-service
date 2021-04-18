package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encryption-service/pkg/config"
	"testing"
)

func TestDecrypt(t *testing.T) {
	cases := []struct{
		caseName, encryptedText, decryptedText string
	}{
		{"case1", "6ac0c4f09977fd57e6d8a4d6af3731e38e9a1363470e8e5ca338c0ca1bc7183980",
			"admin"},
		{"case2", "42b85fda305e307673115f70276835d5052196ff181e83a865d2e702f8e0c5b3029e",
			"admin1"},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			keyString := hex.EncodeToString([]byte(config.GetSecret()))
			keyBytes, _ := hex.DecodeString(keyString)
			encryptedTextBytes, _ := hex.DecodeString(tc.encryptedText)

			block, err := aes.NewCipher(keyBytes)
			if err != nil {
				t.Fatal(err)
			}

			aesGCM, err := cipher.NewGCM(block)
			if err != nil {
				t.Fatal(err)
			}

			nonceSize := aesGCM.NonceSize()
			nonce, ciphertext := encryptedTextBytes[:nonceSize], encryptedTextBytes[nonceSize:]
			plainTextBytes, err := aesGCM.Open(nil, nonce, ciphertext, nil)
			if err != nil {
				t.Fatal(err)
			}

			if string(plainTextBytes) != tc.decryptedText {
				t.Errorf("Decryption was incorrect, got: %s, want: %s.", string(plainTextBytes), tc.decryptedText)
			}
		})
	}
}