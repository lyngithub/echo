package utils

import (
	"echo/conf"
	"echo/utils/aes"
	"github.com/btcsuite/btcd/btcec/v2"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

// 离线生成波场 密钥、地址
func GenerateKey() (wif string, address string, err error) {
	pri, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", err
	}
	if len(pri.Key.Bytes()) != 32 {
		for {
			pri, err = btcec.NewPrivateKey()
			if err != nil {
				continue
			}
			if len(pri.Key.Bytes()) == 32 {
				break
			}
		}
	}
	address = addr.PubkeyToAddress(pri.ToECDSA().PublicKey).String()
	s := pri.Key.String()
	solve := Solve(s)

	aesEcp, err := aes.InitAesEncrypt(conf.Config.Wallet.Key, conf.Config.Wallet.Pk)
	if err != nil {
		return "", "", err
	}
	wif, err = aes.AesBase64Encrypt(solve, aesEcp)
	if err != nil {
		return "", "", err
	}
	return
}

// Solve
func Solve(str string) string {
	if len(str) == 0 {
		return ""
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func X(s string) (string, error) {
	solve := Solve(s)
	aesEcp, err := aes.InitAesEncrypt(conf.Config.Wallet.Key, conf.Config.Wallet.Pk)
	if err != nil {
		return "", err
	}
	wif, err := aes.AesBase64Encrypt(solve, aesEcp)
	if err != nil {
		return "", err
	}
	return wif, nil
}
