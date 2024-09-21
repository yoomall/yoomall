package authresponse

import "lazyfury.github.com/yoomall-server/apps/auth/model"

type LoginResult struct {
	User  *model.User      `json:"user"`
	Token *model.UserToken `json:"token"`
}
