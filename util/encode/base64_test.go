package encode

import (
	"testing"
	"fmt"
)



func Test_Base64(t *testing.T) {
	str := Base64Encrypt([]byte("12345"))
	fmt.Printf("Base64Encrypt result %v \n",str)
	
	b,_:= Base64Decrypt(str)
	fmt.Printf("Base64Decrypt result %v \n",b)
}


