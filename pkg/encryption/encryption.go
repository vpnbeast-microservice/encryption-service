package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encryption-service/pkg/logging"
	"encryption-service/pkg/options"
	"fmt"
	"go.uber.org/zap"
	"io"
)

var (
	logger *zap.Logger
	opts   *options.EncryptionServiceOptions
)

func init() {
	logger = logging.GetLogger()
	opts = options.GetEncryptionServiceOptions()
}

// Encrypt encrypts the string with specified secret
func Encrypt(stringToEncrypt string) (encryptedString string) {
	// Since the key is in string, we need to convert decode it to bytes
	keyString := hex.EncodeToString([]byte(opts.Secret))
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	// Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("an error occurred while creating a new cipher block from key, returning",
			zap.ByteString("keyBytes", key), zap.String("error", err.Error()))
		return
	}

	// Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	// https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Error("an error occurred while creating a new GCM, returning", zap.String("error", err.Error()))
		return
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logger.Error("an error occurred while creating a nonce, returning", zap.String("error", err.Error()))
		return
	}

	// Encrypt the data using aesGCM.Seal
	// Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data.
	// The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}
