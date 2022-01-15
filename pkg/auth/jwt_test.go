package auth_test

import (
	"blog/pkg/auth"
	"os"
	"testing"
)

func TestAes(t *testing.T) {
	os.Setenv("aes_key", "0123456789123456")
	c, err := auth.AesEncrypt([]byte("123456"))
	if err != nil {
		t.Error("encode error", err.Error())
	}
	o, err := auth.AesDecrypt([]byte(c))
	if err != nil {
		t.Error("DECODE ERROR", err.Error())
	}
	if string(o) != "123456" {
		t.Error("compare error")
	}
}
