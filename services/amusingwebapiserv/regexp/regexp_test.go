package regexp

import "testing"

func TestPhoneNumberValid(t *testing.T) {
	if err := PhoneNumberValid("18710565589"); err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("succeed")
}

func TestAreaCodeValid(t *testing.T) {
	if err := AreaCodeValid("1277"); err != nil {
		t.Fatalf("some error: %s", err)
	}
}

func TestPasswordValid(t *testing.T) {
	if err := PasswordValid("Aa121212"); err != nil {
		t.Fatalf("some error: %s", err)
	}
}

func TestNicknameValid(t *testing.T) {
	if err := NicknameValid("xxxxx想—wewe12132"); err != nil {
		t.Fatalf("some error: %s", err)
	}
}

func TestVerificationCodeValid(t *testing.T) {
	if err := VerificationCodeValid("3333"); err != nil {
		t.Fatalf("some error: %s", err)
	}
}
