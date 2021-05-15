package validation

import (
	"encryption-service/pkg/decryption"
	"testing"
)

func TestCompare(t *testing.T) {
	cases := []struct {
		caseName, encryptedText, decryptedText string
	}{
		{"case1", "6ac0c4f09977fd57e6d8a4d6af3731e38e9a1363470e8e5ca338c0ca1bc7183980",
			"admin"},
		{"case2", "42b85fda305e307673115f70276835d5052196ff181e83a865d2e702f8e0c5b3029e",
			"admin1"},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			decryptedText := decryption.Decrypt(tc.encryptedText)
			if tc.decryptedText == decryptedText {
				t.Log("successful operation")
				return
			} else {
				t.Errorf("test failed. required=%v, got=%v\n", tc.decryptedText, decryptedText)
				return
			}
		})
	}
}
