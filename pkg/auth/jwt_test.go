package auth_test

import (
	"blog/pkg/auth"
	"encoding/hex"
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

func TestJwt(t *testing.T) {
	os.Setenv("jwt_key", "123456")
	token, err := auth.CreateToken("123454", []byte("12345"))
	if err != nil {
		t.Error(err.Error())
	}
	result, err := auth.ParseToken(token)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result.Username + " " + result.Password)

}

func TestTEST(t *testing.T) {
	temp := []byte("12345")
	str := hex.EncodeToString(temp)
	t.Log(str)
}
