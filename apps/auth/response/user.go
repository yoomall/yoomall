package authresponse

import "github.com/yoomall/yoomall/apps/auth/model"

type LoginResult struct {
	User  *model.User      `json:"user"`
	Token *model.UserToken `json:"token"`
}
