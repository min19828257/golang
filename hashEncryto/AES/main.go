// package main

// import (
// 	"crypto/aes"
// 	"fmt"
// )

// func main() {
// 	key := "Hello, Key 12345"    // 16 바이트
// 	source := "Hello, world! 12" // 16 바이트

// 	block, err := aes.NewCipher([]byte(key)) //AES 대칭키 암호화 블록 생성
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	ciphertext := make([]byte, len(source))
// 	block.Encrypt(ciphertext, []byte(source))
// 	fmt.Printf("%x\n", ciphertext)

// 	plaintext := make([]byte, len(source))
// 	block.Decrypt(plaintext, ciphertext) // AES 알고리즘으로 암호화된 데이터를 평문으로 복호화

// 	fmt.Println(string(plaintext))
// }

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var (
	initialVector = "1234567890123456"
	passphrase    = "Impassphrasegood"
)

func main() {
	var plainText = "hello wosaaaaaaassssssssssssssrld"

	encryptedData := AESEncrypt(plainText, []byte(passphrase))
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
	fmt.Println(encryptedString)

	encryptedData, _ = base64.StdEncoding.DecodeString(encryptedString)
	decryptedText := AESDecrypt(encryptedData, []byte(passphrase))
	fmt.Println(string(decryptedText))
}

func AESEncrypt(src string, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func AESDecrypt(crypt []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
