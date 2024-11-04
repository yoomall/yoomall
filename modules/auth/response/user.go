package authresponse

import "yoomall/modules/auth/model"

type LoginResult struct {
	User  *model.User      `json:"user"`
	Token *model.UserToken `json:"token"`
}
