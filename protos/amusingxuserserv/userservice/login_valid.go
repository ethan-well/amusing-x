package userservice

func (x *LoginRequest) LoginByPassword() bool {
	return x.Type == 0
}

func (x *LoginRequest) LoginByVerificationCode() bool {
	return x.Type == 1
}
