package crypto

import (
	"testing"
)



func Test_Aes(t *testing.T) {
	aes_key := "0102030405060708" //0910203040506070"
	ivEn := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	
	Aes := NewGoAes(aes_key,ivEn)
	Aes.Init()
	
	v1 := Aes.Encrypt([]byte("1234"))
	Aes.Decrypt(v1)
	v2 := Aes.Encrypt([]byte("1234哈哈哈哈哈哈哈哈"))
	Aes.Decrypt(v2)
	v3 := Aes.Encrypt([]byte("23456123456123456123456123456123456123456123456123456哈哈"))
	Aes.Decrypt(v3)
	

}


