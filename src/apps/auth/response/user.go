package authresponse

import "yoomall/src/apps/auth/model"

type LoginResult struct {
	User  *model.User      `json:"user"`
	Token *model.UserToken `json:"token"`
}
