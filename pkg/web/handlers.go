package web

import (
	"encoding/json"
	"encryption-service/pkg/decryption"
	"encryption-service/pkg/encryption"
	"encryption-service/pkg/logging"
	"encryption-service/pkg/validation"
	"errors"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	var request encryptRequest
	err := decodeJSONBody(w, r, &request)
	if err != nil {
		logger.Error("an error occurred while decoding json body", zap.String("error", err.Error()))
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	encryptedText := encryption.Encrypt(request.PlainText)
	response := encryptResponse{
		Tag:    "encrypt",
		Output: encryptedText,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("an error occurred while marshaling response", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseBytes)
	if err != nil {
		logger.Error("an error occurred while writing response", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	var request decryptRequest
	err := decodeJSONBody(w, r, &request)
	if err != nil {
		logger.Error("an error occurred while decoding json body", zap.String("error", err.Error()))
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	decryptedText := decryption.Decrypt(request.EncryptedText)
	response := decryptResponse{
		Tag:    "decrypt",
		Output: decryptedText,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("an error occurred while marshaling response", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseBytes)
	if err != nil {
		logger.Error("an error occurred while writing response", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	var request checkRequest
	err := decodeJSONBody(w, r, &request)
	if err != nil {
		logger.Error("an error occurred while decoding json body", zap.String("error", err.Error()))
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	isMatched := validation.Compare(request.PlainText, request.EncryptedText)
	response := checkResponse{
		Tag:    "check",
		Status: isMatched,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("an error occurred while marshaling response", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseBytes)
	if err != nil {
		logger.Error("an error occurred while writing response", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		logger.Error("an error occurred while writing response", zap.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
