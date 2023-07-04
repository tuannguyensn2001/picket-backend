package dto

type UpdateAvatarInput struct {
	Url    string `json:"url" form:"url" binding:"required" validate:"required,url"`
	UserId int
}
