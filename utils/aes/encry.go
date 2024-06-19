package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AesEncrypt struct {
	key   []byte
	iv    []byte
	block cipher.Block
}

func InitAesEncrypt(key, pk string) (*AesEncrypt, error) {
	aesEcp := &AesEncrypt{}
	aesEcp.key = []byte(key[:16])
	aesEcp.iv = []byte(pk[:16])
	var err error
	aesEcp.block, err = aes.NewCipher(aesEcp.key)
	if err != nil {
		return nil, err
	}
	return aesEcp, nil
}

// AesBase64Encrypt
func AesBase64Encrypt(in string, a *AesEncrypt) (string, error) {
	origData := []byte(in)
	origData = pKCS5Adding(origData, a.block.BlockSize())
	crypted := make([]byte, len(origData))
	bm := cipher.NewCBCEncrypter(a.block, a.iv)
	bm.CryptBlocks(crypted, origData)
	var b = base64.StdEncoding.EncodeToString(crypted)
	return b, nil
}

func AesBase64Decrypt(b string, a *AesEncrypt) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return "", err
	}
	origData := make([]byte, len(crypted))
	bm := cipher.NewCBCDecrypter(a.block, a.iv)
	bm.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	var out = string(origData)
	return out, nil
}

// PKCS5Adding
func pKCS5Adding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS5UnPadding
func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
