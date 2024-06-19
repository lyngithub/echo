package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func splitBytes(b []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(b); i += chunkSize {
		end := i + chunkSize
		if end > len(b) {
			end = len(b)
		}
		chunks = append(chunks, b[i:end])
	}
	return chunks
}

func openssl_RSADecrypt(src, priKey []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, errors.New("key is invalid format")
	}

	// x509 parse
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	dst, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), src)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func openssl_RSAEncrypt(src, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("key is invalid format")
	}

	// x509 parse
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	dst, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), src)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func DecryptFile(private_pem_path, data string) ([]byte, error) {
	if len(data) == 0 {
		return []byte{}, errors.New("empty data")
	}

	privateKey, err := ioutil.ReadFile(private_pem_path)
	if err != nil {
		return []byte{}, err
	}

	// c.setupPriKey()
	encrypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return []byte{}, err
	}
	var decrypted []byte
	if len(encrypted) > 256 {
		for _, chunk := range splitBytes(encrypted, 256) {

			decryptData, err := openssl_RSADecrypt(chunk, privateKey)
			if err != nil {
				return []byte{}, err
			}
			decrypted = append(decrypted, decryptData...)
		}
	} else {
		decryptData, err := openssl_RSADecrypt(encrypted, privateKey)
		if err != nil {
			return []byte{}, err
		}
		decrypted = decryptData
	}

	//dat, err := base64.StdEncoding.DecodeString(string(decrypted))
	//if err != nil {
	//	return []byte{}, err
	//}
	return decrypted, nil
}

func EncryptFile(public_pem_path string, data []byte) (string, error) {
	if len(data) == 0 {
		return "", errors.New("empty data")
	}

	publicKey, err := ioutil.ReadFile(public_pem_path)
	if err != nil {
		return "", err
	}

	dat := []byte(base64.StdEncoding.EncodeToString(data))

	var decrypted []byte
	if len(dat) > 245 {
		for _, chunk := range splitBytes(dat, 245) {

			decryptData, err := openssl_RSAEncrypt(chunk, publicKey)
			if err != nil {
				return "", err
			}
			decrypted = append(decrypted, decryptData...)
		}
	} else {
		decryptData, err := openssl_RSAEncrypt(dat, publicKey)
		if err != nil {
			return "", err
		}
		decrypted = decryptData
	}

	if err != nil {
		return "", err
	}
	encrypted := base64.StdEncoding.EncodeToString(decrypted)
	return encrypted, nil
}
