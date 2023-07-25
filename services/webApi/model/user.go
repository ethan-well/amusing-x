package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/services/webApi/mysql/amusingxwebapi"
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"github.com/ItsWewin/superfactory/aerror"
	"golang.org/x/crypto/argon2"
)

type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

type User struct {
	ID             int64  `db:"id"`
	Nickname       string `db:"nickname"`
	AreaCode       string `db:"area_code"`
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

func (u *User) ExistedWithPhone(ctx context.Context) (bool, aerror.Error) {
	if len(u.Phone) == 0 {
		return false, aerror.NewError(nil, aerror.Code.CParamsErr, "params is invalid. ")
	}

	udb, err := FindUserByPhone(ctx, u.AreaCode, u.Phone)
	if err != nil {
		return false, err
	}

	return udb != nil, nil
}

func (u *User) ExistedWithNicknameOrPhone(ctx context.Context) (bool, aerror.Error) {
	if len(u.Nickname) == 0 || len(u.Phone) == 0 {
		return false, aerror.NewError(nil, aerror.Code.CParamsErr, "params is invalid. ")
	}

	udb, err := amusingxwebapi.QueryUserByNicknameOrPhone(ctx, u.Nickname, u.Phone)
	if err != nil {
		return false, err
	}

	switch {
	case udb == nil:
		return false, nil
	case udb.Nickname == u.Nickname:
		return true, aerror.NewError(nil, aerror.Code.BDataIsNotAllow, "该昵称已经被占用")
	case udb.Phone == u.Phone:
		return true, aerror.NewError(nil, aerror.Code.BDataIsNotAllow, "此手机号已经被占用")
	default:
		return true, aerror.NewError(nil, aerror.Code.BDataIsNotAllow, "该用户已经存在")
	}
}

func (u *User) ComparePassword(password string) (bool, aerror.Error) {
	if len(password) == 0 ||
		len(u.Salt) == 0 ||
		len(u.PasswordDigest) == 0 ||
		passwordConfig == nil {

		return false, aerror.NewError(nil, aerror.Code.BDataIsNotAllow, "密码不符合要求")
	}

	salt, err := base64.RawStdEncoding.DecodeString(u.Salt)
	if err != nil {
		return false, aerror.NewError(err, aerror.Code.SUnexpectedErr, "Salt decode failed. ")
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(u.PasswordDigest)
	if err != nil {
		return false, aerror.NewError(err, aerror.Code.SUnexpectedErr, "Password decode failed. ")
	}

	c := passwordConfig
	c.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}

func (u *User) ResetPassword(ctx context.Context, password string) aerror.Error {
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
	_, err = amusingxwebapi.UpdatePassword(ctx, udb)
	if err != nil {
		return aerror.NewError(err, err.Code(), err.Message())
	}

	return nil
}

func (u *User) GeneratePassword() aerror.Error {
	if len(u.Password) == 0 || passwordConfig == nil {
		return aerror.NewError(nil, aerror.Code.BDataIsNotAllow, "Password is blank")
	}

	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "Unexpected err")
	}

	c := passwordConfig
	hash := argon2.IDKey([]byte(u.Password), salt, c.time, c.memory, c.threads, c.keyLen)

	u.PasswordDigest = base64.RawStdEncoding.EncodeToString(hash)
	u.Salt = base64.RawStdEncoding.EncodeToString(salt)

	return nil
}

func FindUserByID(ctx context.Context, id int64) (*User, aerror.Error) {
	user, err := amusingxwebapi.QueryUserByIdContext(ctx, id)
	if err != nil {
		return nil, err
	}

	u := &User{
		ID:             user.ID,
		Nickname:       user.Nickname,
		Phone:          user.Phone,
		PasswordDigest: user.PasswordDigest,
		CreateTime:     user.CreateTime,
		UpdateTime:     user.UpdateTime,
	}

	return u, nil
}

func FindUserByPhone(ctx context.Context, areaCode, phone string) (*User, aerror.Error) {
	user, err := amusingxwebapi.QueryUserByPhone(ctx, areaCode+"-"+phone)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	u := &User{
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

func Create(ctx context.Context, user *User) (*User, aerror.Error) {
	err := user.GeneratePassword()

	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SUnexpectedErr, "Generate password failed. ")
	}

	udb := &ganymede.UserComplex{
		Nickname:       user.Nickname,
		Phone:          user.Phone,
		PasswordDigest: user.PasswordDigest,
		Salt:           user.Salt,
	}

	udb, err = amusingxwebapi.Insert(ctx, udb)
	if err != nil {
		return nil, err
	}

	user.ID = udb.ID

	return user, nil
}

func clearPassword(u *User) {
	u.Password = ""
	u.PasswordDigest = ""
}
