package aes

import (
	"fmt"
	"testing"
)

func Test_AesBase64Encrypt(t *testing.T) {
	aesEcp, err := InitAesEncrypt("qwertyuiopasdfghjkl", "zxcvbnmasdfghjklqw")
	if err != nil {
		t.Fatal(err)
	}
	str := "{\"imei\":\"10\",\"content\":\"防守打法施工方大哥嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎嘎\",\"phone\":\"65535\",\"amount\":\"1.0100000000\",\"bank_name\":\"招商银行\",\"number\":\"2222\",\"name\":\"张三\",\"date\":\"2021-05-05 13:00\"}\n"
	base64Encrypt, err := AesBase64Encrypt(str, aesEcp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("====== base64Encrypt: %v", base64Encrypt)
}

func Test_AesBase64Decrypt(t *testing.T) {
	aesEcp, err := InitAesEncrypt("aaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbb")
	if err != nil {
		t.Fatal(err)
	}
	str := "sHvuTnofxFb8qZdST99ANpQNo7SscCT0M3JKkHzuwTEpoXVHs8PGDPLO3LTCxgduRLvOYnYG3UY3kQW0J5jS9dJsJsSOzPGrLBjAN7Qoays="
	base64Decrypt, err := AesBase64Decrypt(str, aesEcp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("====== base64Decrypt: %v", base64Decrypt)
}
