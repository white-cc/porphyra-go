package models_test

import (
	"blog/models"
	"os"
	"testing"
)

func setENV() {
	os.Setenv("MONGODB_URI", "mongodb://admin:123456@127.0.0.1:27017/?retryWrites=true&w=majority")
	os.Setenv("aes_key", "0123456789123456")
}

func TestUser(t *testing.T) {
	setENV()
	models.Init()
	models.LoadUserdbModel()
	id_1, err := models.AddUser("admin2", "1234567", "123@qq.com")
	if err != nil {
		t.Error(err)
	}
	id_2, err := models.FindUser("admin2", "1234567")
	if err != nil {
		t.Error(err)
	}
	if id_2.Username != "admin2" {
		t.Error(id_1)
	}
	t.Log(id_1)
}
