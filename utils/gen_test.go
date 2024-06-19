package utils

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	wif, address, _ := GenerateKey()
	fmt.Println(wif)
	fmt.Println(address)
}
