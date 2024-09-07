package request

type UserUserNameAndPasswordLoginRequest struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}

type UserCreateRequest struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Role     int    `json:"role" form:"role"`
}
