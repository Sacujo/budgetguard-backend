package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("qwerty123")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hash)

	if err := CheckPassword("qwerty123", hash); err != nil {
		t.Error("верный пароль не прошёл проверку")
	}
	if err := CheckPassword("wrong", hash); err == nil {
		t.Error("неверный пароль прошёл проверку")
	}
}
