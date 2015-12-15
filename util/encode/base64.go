package encode 


import (
	"encoding/base64"
)


func Base64Encrypt(data []byte) string{
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

func Base64Decrypt(data string) ([]byte,error) {
	b,err := base64.StdEncoding.DecodeString(data)
	return b,err
}