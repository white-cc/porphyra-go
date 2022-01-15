package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"os"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData []byte) ([]byte, error) {
	aes_key := os.Getenv("aes_key")
	if aes_key == "" {
		err := errors.New("no aes_key")
		return nil, err
	}

	key := []byte(aes_key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()                              //块大小
	origData = PKCS7Padding(origData, blockSize)                //填充
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //加密器
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData) //加密
	return crypted, nil
}

func AesDecrypt(crypted []byte) ([]byte, error) {
	aes_key := os.Getenv("aes_key")
	if aes_key == "" {
		err := errors.New("no aes_key")
		return nil, err
	}
	key := []byte(aes_key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil

}
