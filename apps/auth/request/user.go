package request

type UserUserNameAndPasswordLoginRequest struct {
	UserName string `json:"userName" form:"userName" swag:"string,用户名,required"`
	Password string `json:"password" form:"password" swag:"string,密码,required"`
}

type UserCreateRequest struct {
	UserName string `json:"userName" form:"userName" swag:"string,用户名,required"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Role     int    `json:"role" form:"role"`
}
