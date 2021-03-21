package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encryption-service/pkg/config"
	"log"
)

func Decrypt(encryptedText string) (decryptedText string) {
	keyString := hex.EncodeToString([]byte(config.GetSecret()))
	keyBytes, _ := hex.DecodeString(keyString)
	encryptedTextBytes, _ := hex.DecodeString(encryptedText)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
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