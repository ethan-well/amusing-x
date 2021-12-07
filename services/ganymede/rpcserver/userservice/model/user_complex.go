package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	amusinguser2 "amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"github.com/ItsWewin/superfactory/xerror"
	"golang.org/x/crypto/argon2"
)

type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

type UserComplex struct {
	ID             int64  `db:"id"`
	Nickname       string `db:"nickname"`
	AreaCode       string `json:"area_code"`
	Phone          string `db:"phone"`
	PasswordDigest string `db:"password_digest"`
	Password       string
	Salt           string
	CreateTime     string `db:"create_time"`
	UpdateTime     string `db:"update_time"`
}

// 密码配置
var passwordConfig = &PasswordConfig{
	time:    2,
	memory:  64 * 1024,
	threads: 4,
	keyLen:  64,
}

func (u *UserComplex) ExistedWithPhone(ctx context.Context) (bool, *xerror.Error) {
	if len(u.Phone) == 0 {
		return false, xerror.NewError(nil, xerror.Code.CParamsError, "params is invalid. ")
	}

	udb, err := FindUserByPhone(ctx, u.AreaCode, u.Phone)
	if err != nil {
		return false, err
	}

	return udb != nil, nil
}

func (u *UserComplex) ExistedWithNicknameOrPhone(ctx context.Context) (bool, *xerror.Error) {
	if len(u.Nickname) == 0 || len(u.Phone) == 0 {
		return false, xerror.NewError(nil, xerror.Code.CParamsError, "params is invalid. ")
	}

	udb, err := amusinguser2.QueryUserByNicknameOrPhone(ctx, u.Nickname, u.Phone)
	if err != nil {
		return false, err
	}

	switch {
	case udb == nil:
		return false, nil
	case udb.Nickname == u.Nickname:
		return true, xerror.NewError(nil, xerror.Code.BDataIsNotAllow, "该昵称已经被占用")
	case udb.Phone == u.Phone:
		return true, xerror.NewError(nil, xerror.Code.BDataIsNotAllow, "此手机号已经被占用")
	default:
		return true, xerror.NewError(nil, xerror.Code.BDataIsNotAllow, "该用户已经存在")
	}
}

func (u *UserComplex) ComparePassword(password string) (bool, *xerror.Error) {
	if len(password) == 0 ||
		len(u.Salt) == 0 ||
		len(u.PasswordDigest) == 0 ||
		passwordConfig == nil {

		return false, xerror.NewError(nil, xerror.Code.BDataIsNotAllow, "密码不符合要求")
	}

	salt, err := base64.RawStdEncoding.DecodeString(u.Salt)
	if err != nil {
		return false, xerror.NewError(err, xerror.Code.SUnexpectedErr, "Salt decode failed. ")
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(u.PasswordDigest)
	if err != nil {
		return false, xerror.NewError(err, xerror.Code.SUnexpectedErr, "Password decode failed. ")
	}

	c := passwordConfig
	c.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}

func (u *UserComplex) ResetPassword(ctx context.Context, password string) *xerror.Error {
	u.Password = password
	err := u.GeneratePassword()
	if err != nil {
		return err
	}
	defer clearPassword(u)

	udb := &ganymede.UserComplex{
		Phone:          u.Phone,
		PasswordDigest: u.PasswordDigest,
		Salt:           u.Salt,
	}
	_, err = amusinguser2.UpdatePassword(ctx, udb)
	if err != nil {
		return xerror.NewError(err, err.Code, err.Message)
	}

	return nil
}

func (u *UserComplex) GeneratePassword() *xerror.Error {
	if len(u.Password) == 0 || passwordConfig == nil {
		return xerror.NewError(nil, xerror.Code.BDataIsNotAllow, "Password is blank")
	}

	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "Unexpected err")
	}

	c := passwordConfig
	hash := argon2.IDKey([]byte(u.Password), salt, c.time, c.memory, c.threads, c.keyLen)

	u.PasswordDigest = base64.RawStdEncoding.EncodeToString(hash)
	u.Salt = base64.RawStdEncoding.EncodeToString(salt)

	return nil
}

func FindUserByID(ctx context.Context, id int64) (*UserComplex, *xerror.Error) {
	user, err := amusinguser2.QueryUserByIdContext(ctx, id)
	if err != nil {
		return nil, err
	}

	u := &UserComplex{
		ID:             user.ID,
		Nickname:       user.Nickname,
		Phone:          user.Phone,
		PasswordDigest: user.PasswordDigest,
		CreateTime:     user.CreateTime,
		UpdateTime:     user.UpdateTime,
	}

	return u, nil
}

func FindUserByPhone(ctx context.Context, areaCode, phone string) (*UserComplex, *xerror.Error) {
	user, err := amusinguser2.QueryUserByPhone(ctx, areaCode+"-"+phone)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	u := &UserComplex{
		ID:             user.ID,
		Nickname:       user.Nickname,
		Phone:          user.Phone,
		PasswordDigest: user.PasswordDigest,
		CreateTime:     user.CreateTime,
		UpdateTime:     user.UpdateTime,
		Salt:           user.Salt,
	}

	return u, nil
}

func Create(ctx context.Context, user *UserComplex) (*UserComplex, *xerror.Error) {
	err := user.GeneratePassword()

	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SUnexpectedErr, "Generate password failed. ")
	}

	udb := &ganymede.UserComplex{
		Nickname:       user.Nickname,
		Phone:          user.Phone,
		PasswordDigest: user.PasswordDigest,
		Salt:           user.Salt,
	}

	udb, err = amusinguser2.InsertUserComplex(ctx, udb)
	if err != nil {
		return nil, err
	}

	user.ID = udb.ID

	return user, nil
}

func clearPassword(u *UserComplex) {
	u.Password = ""
	u.PasswordDigest = ""
}
