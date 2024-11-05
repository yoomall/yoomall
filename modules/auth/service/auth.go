package authservice

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"yoomall/modules/auth/model"
	authresponse "yoomall/modules/auth/response"
	core "yoomall/yoo"
	"yoomall/yoo/driver"
	"yoomall/yoo/helper/response"
	"yoomall/yoo/result"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB *driver.DB
}

func NewAuthService(db *driver.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (s *AuthService) createToken(userId uint, ctx *gin.Context) *model.UserToken {
	str := uuid.New().String()

	ip := ctx.ClientIP()

	userAgent := ctx.Request.UserAgent()
	device := "unknown"
	os := "unknown"
	browser := "unknown"
	if ua, err := s.parseUserAgent(userAgent); err == nil {
		device = ua.Device
		os = ua.OS
		browser = ua.Browser
	}

	token := model.UserToken{
		Token:      str,
		ExpireTime: time.Now().Add(24 * time.Hour),
		UserId:     userId,
		IP:         ip,
		Agent:      userAgent,
		Device:     device,
		OS:         os,
		Browser:    browser,
	}
	if err := s.DB.Create(&token).Error; err == nil {
		return &token
	}
	panic("生成 token 失败，should not reach here")
}

type UserAgent struct {
	Device  string
	OS      string
	Browser string
}

func (s *AuthService) parseUserAgent(userAgent string) (*UserAgent, error) {
	ua := &UserAgent{
		Device:  "unknown",
		OS:      "unknown",
		Browser: "unknown",
	}
	osArr := []string{"Windows", "Macintosh", "Linux", "Android", "IOS"}
	browserArr := []string{"Chrome", "Firefox", "Safari", "Opera", "IE", "Edge"}

	osRe := fmt.Sprintf("(%s)", strings.Join(osArr, "|"))
	browserRe := fmt.Sprintf("(%s)", strings.Join(browserArr, "|"))

	if match := regexp.MustCompile(osRe).FindStringSubmatch(userAgent); len(match) > 0 {
		ua.OS = match[0]
		switch ua.OS {
		case "Windows":
			ua.Device = "pc"
		case "Macintosh":
			ua.Device = "mac"
		case "Linux":
			ua.Device = "pc"
		case "Android":
			ua.Device = "mobile"
		case "IOS":
			ua.Device = "mobile"
		}
	}

	if match := regexp.MustCompile(browserRe).FindStringSubmatch(userAgent); len(match) > 0 {
		ua.Browser = match[0]
	}

	return ua, nil
}
func (s *AuthService) HashedPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func (s *AuthService) LoginWithUsernameAndPassword(username string, password string, ctx *gin.Context) *result.Result[*authresponse.LoginResult] {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("AuthService.LoginWithUsernameAndPassword", err)
		}
	}()

	if username == "" || password == "" {
		return result.Err[*authresponse.LoginResult](fmt.Errorf("username or password is empty"))
	}

	var user model.User
	err := s.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return result.Err[*authresponse.LoginResult](fmt.Errorf("用户不存在"))
	}

	user.LastLoginAt = (core.LocalTime)(time.Now())

	if err := s.DB.Save(&user).Error; err != nil {
		return result.Err[*authresponse.LoginResult](fmt.Errorf("更新用户信息失败"))
	}

	// check passwrod
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return result.Err[*authresponse.LoginResult](fmt.Errorf("密码错误"))
	}

	// 复用近期的 token todo:需要一个设备信息的粗 id，来处理多设备登录的情况/或者踢掉其他的所有 token/尽量避免多设备共用一个 token 的情况
	var findToken *model.UserToken = new(model.UserToken)
	if err := s.DB.Where("user_id = ?", user.ID).First(&findToken).Error; err == nil {
		if findToken.ExpireTime.After(time.Now()) {
			findToken.ExpireTime = time.Now().Add(24 * time.Hour)
			s.DB.Save(findToken)
			return result.Ok(&authresponse.LoginResult{
				User:  &user,
				Token: findToken,
			})
		}
	}

	// 生成新 token
	var userToken *model.UserToken = s.createToken(user.ID, ctx)
	return result.Ok(&authresponse.LoginResult{
		User:  &user,
		Token: userToken,
	})
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

	find := s.DB.Where("username = ?", user.UserName).Where("email = ?", user.Email).Where("phone = ?", user.Phone).First(&model.User{}).Error
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

// logout

func (s *AuthService) Logout(ctx *gin.Context) {
	userToken := ctx.MustGet("token").(model.UserToken)
	if err := s.DB.Delete(&userToken).Error; err != nil {
		response.Error(response.ErrInternalError, err.Error()).Done(ctx)
		return
	}
	response.Success("退出成功").Done(ctx)
}
