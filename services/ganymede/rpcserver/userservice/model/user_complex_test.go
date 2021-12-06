package model

import (
	"testing"
)

func TestSignPassword(t *testing.T) {
	u := UserComplex{
		ID:             0,
		Nickname:       "",
		Phone:          "",
		PasswordDigest: "",
		Password:       "12345678",
		CreateTime:     "",
		UpdateTime:     "",
	}

	err := u.GeneratePassword()
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("salt: %s, passwordDigest: %s", u.Salt, u.PasswordDigest)
}

func TestComparePassword(t *testing.T) {
	u := &UserComplex{
		ID:             0,
		Nickname:       "",
		Phone:          "",
		PasswordDigest: "Ugq4fOGHApEUMj7jqTGfZGsJZWJkeOGxpi3a29KWGjd733TNgPqu8AgbsoyLzq32obXbt8eXv5MEMYpk3GoCEw",
		Salt:           "+MRRIKt7KHbIo2SuV/446g",
		CreateTime:     "",
		UpdateTime:     "",
	}

	result, err := u.ComparePassword("12345678")
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("result: %t", result)
}
