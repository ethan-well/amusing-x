package xconst

type GlobalConst struct {
	Login Login
}

type LoginType map[string]int
type Login struct {
	LoginType LoginType
}

var globalConst GlobalConst

func InitGlobalConst() {
	globalConst = GlobalConst{Login: Login{LoginType: LoginType{"password": 0, "verificationCode": 1}}}
}

func (t LoginType) IsPasswordType() bool {
	for k, v := range t {

	}
}
