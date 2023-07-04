package dto

type RegisterInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginGoogleInput struct {
	Code string `json:"code" form:"code" binding:"required"`
}

type NewUserRegisterPayload struct {
	UserId int `json:"user_id"`
}

type InitPasswordInput struct {
	Password string `json:"password" form:"password" binding:"required"`
	UserId   int
}

type ResetPasswordInput struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
	UserId      int
}
