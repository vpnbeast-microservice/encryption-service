package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encryption-service/pkg/options"
	commons "github.com/vpnbeast/golang-commons"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	opts   *options.EncryptionServiceOptions
)

func init() {
	logger = commons.GetLogger()
	opts = options.GetEncryptionServiceOptions()
}

// Decrypt decrypts specified string with secret
func Decrypt(encryptedText string) (decryptedText string) {
	keyString := hex.EncodeToString([]byte(opts.Secret))
	keyBytes, _ := hex.DecodeString(keyString)
	encryptedTextBytes, _ := hex.DecodeString(encryptedText)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		logger.Error("an error occurred while creating a new cipher block from key, returning",
			zap.ByteString("keyBytes", keyBytes), zap.String("error", err.Error()))
		return
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Error("an error occurred while creating a new GCM, returning", zap.String("error", err.Error()))
		return
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := encryptedTextBytes[:nonceSize], encryptedTextBytes[nonceSize:]

	//Decrypt the data
	plainTextBytes, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}

	return string(plainTextBytes)
}
