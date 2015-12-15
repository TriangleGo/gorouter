package crypto



import (
	"fmt"
	"pkg/crypto/aes"
	"crypto/cipher"
	
)

type GoAes struct {
	AesKey string
	ivEn []byte
	cib cipher.Block
}


//this Aes model uses AES-128-CBC block mode
func NewGoAes(key string,iv []byte) *GoAes{
	return &GoAes{AesKey:key,ivEn:iv}
}

func (this *GoAes)Init() error{
	var err error
	this.cib,err = aes.NewCipher( []byte(this.AesKey) )
	if err != nil {
		return err
	}	
	return nil
}

func (this *GoAes)Encrypt(data []byte)  []byte {
	//encrypt the encode
	in := this.pad( data )
	enc := make([]byte,len(in))
	blockMode := cipher.NewCBCEncrypter(this.cib,this.ivEn)
	blockMode.CryptBlocks(enc,in)
	//fmt.Printf("Encrypte result = %v \n",enc)
	return enc
}

func (this *GoAes)Decrypt(data []byte) []byte {
	dec := make([]byte,len(data))
	DecryptoBlockMode := cipher.NewCBCDecrypter(this.cib,this.ivEn)
	DecryptoBlockMode.CryptBlocks(dec,data)
	//fmt.Printf("Decrypte:%v\n %v\n",dec,string(this.unpad(dec)))
	return this.unpad(dec)
}

func (this *GoAes) pad(data []byte) []byte{
	//count padding 
	paddingLen := aes.BlockSize - (  len(data) % aes.BlockSize )
	
	// in must  x*aes.BlockSize
	in := make([]byte, aes.BlockSize * ( (len(data) /  aes.BlockSize) + 1))
	
	//fmt.Printf(" in size = %v \n\n",len(in))
	// fill data with paddingLen
	for  i :=0;i<len(in);i++ {
		in[i] = byte(paddingLen)
	}
	copy(in ,data)
	
	//fmt.Printf("in is %v :length %v\n",in,len(in))
	
	return in
}

func (this *GoAes) unpad(data []byte) []byte{
	pad := int(data[ len(data) - 1])
	
	if pad == 0 {
		panic("aes unpad error\n")
	}
	
	if pad > aes.BlockSize {
		panic("aes unpad error\n")
	}
	
	return data[0:len(data) - pad]
}

