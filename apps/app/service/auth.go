package service

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"lazyfury.github.com/yoomall-server/apps/app/model"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type AuthService struct {
	DB *driver.DB
}

func NewAuthService(db *driver.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (s *AuthService) genToken() string {
	return uuid.NewString()
}

func (s *AuthService) HashedPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func (s *AuthService) LoginWithUsernameAndPassword(username string, password string) (*model.User, *model.UserToken, error) {
	if username == "" || password == "" {
		return nil, nil, fmt.Errorf("username or password is empty")
	}

	var user model.User
	err := s.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, nil, fmt.Errorf("用户不存在")
	}

	// check passwrod
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, fmt.Errorf("密码错误")
	}

	var userToken = model.UserToken{
		UserId:     user.ID,
		Token:      s.genToken(),
		ExpireTime: time.Now().Add(24 * time.Hour),
	}

	err = s.DB.Create(&userToken).Error

	if err != nil {
		return nil, nil, fmt.Errorf("登录失败")
	}

	return &user, &userToken, nil
}

func (s *AuthService) CheckPasswordStrength(password string) error {
	lenth := len(password)
	if lenth < 8 {
		return fmt.Errorf("密码太短")
	}
	if lenth > 32 {
		return fmt.Errorf("密码太长")
	}

	// has letter and number
	if !regexp.MustCompile(`[a-zA-Z0-9]+`).MatchString(password) {
		return fmt.Errorf("密码必须包含字母和数字")
	}

	return nil
}

func (s *AuthService) CreateUser(user *model.User) error {
	if err := s.CheckPasswordStrength(user.Password); err != nil {
		return err
	}

	find := s.DB.Where("username = ?", user.UserName).First(&model.User{}).Error
	if find == nil {
		return fmt.Errorf("用户已存在")
	}

	user.Password, _ = s.HashedPassword(user.Password)
	return s.DB.Create(user).Error
}

func (s *AuthService) UpdateUser(user *model.User) error {
	return s.DB.Save(user).Error
}

func (s *AuthService) GetUser(id int) (*model.User, error) {
	var user model.User
	err := s.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (s *AuthService) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := s.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
