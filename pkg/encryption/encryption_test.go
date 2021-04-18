package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encryption-service/pkg/config"
	"io"
	"testing"
)

func TestEncrypt(t *testing.T) {
	cases := []struct{
		caseName, clearText string
	}{
		{"case1", "admin"},
		{"case1", "admin1"},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			keyString := hex.EncodeToString([]byte(config.GetSecret()))
			key, _ := hex.DecodeString(keyString)
			plaintext := []byte(tc.clearText)

			block, err := aes.NewCipher(key)
			if err != nil {
				t.Fatal(err)
			}

			aesGCM, err := cipher.NewGCM(block)
			if err != nil {
				t.Fatal(err)
			}

			nonce := make([]byte, aesGCM.NonceSize())
			if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
				t.Fatal(err)
			}

			_ = aesGCM.Seal(nonce, nonce, plaintext, nil)
		})
	}
}