package authresponse

import "yoomall/apps/auth/model"

type LoginResult struct {
	User  *model.User      `json:"user"`
	Token *model.UserToken `json:"token"`
}
