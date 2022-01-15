package jwt_test

import (
	"blog/pkg/jwt"
	"os"
	"testing"
)

func TestAes(t *testing.T) {
	os.Setenv("aes_key", "helloworld123456")
	c, err := jwt.AesEncrypt([]byte("123456"))
	if err != nil {
		t.Error("encode error", err.Error())
	}
	o, err := jwt.AesDecrypt([]byte(c))
	if err != nil {
		t.Error("DECODE ERROR", err.Error())
	}
	if o != "123456" {
		t.Error("compare error")
	}
}
