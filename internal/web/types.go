package web

type checkRequest struct {
	PlainText     string `json:"plainText"`
	EncryptedText string `json:"encryptedText"`
}

type checkResponse struct {
	Status bool `json:"status"`
}

type encryptRequest struct {
	PlainText string `json:"plainText"`
}

type encryptResponse struct {
	Output string `json:"output"`
}

type decryptRequest struct {
	EncryptedText string `json:"encryptedText"`
}

type decryptResponse struct {
	Output string `json:"output"`
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}
