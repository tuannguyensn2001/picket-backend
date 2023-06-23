package entities

type User struct {
	Id              int      `json:"id" gorm:"column:id"`
	Email           string   `json:"email" gorm:"column:email"`
	Password        string   `json:"password" gorm:"column:password"`
	Username        string   `json:"username" gorm:"column:username"`
	EmailVerifiedAt int64    `json:"email_verified_at" gorm:"column:email_verified_at"`
	CreatedAt       int64    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       int64    `json:"updated_at" gorm:"column:updated_at"`
	Wallet          *Wallet  `json:"wallet,omitempty" `
	Version         int      `json:"version" gorm:"column:version"`
	Profile         *Profile `json:"profile,omitempty" `
}

type Profile struct {
	Id        int    `json:"id" gorm:"column:id"`
	UserId    int    `json:"user_id" gorm:"column:user_id"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url"`
	Nickname  string `json:"nickname" gorm:"column:nickname"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at"`
}
