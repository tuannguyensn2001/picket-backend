package entities

type User struct {
	Id              int     `json:"id" gorm:"column:id"`
	Email           string  `json:"email" gorm:"column:email"`
	Password        string  `json:"password" gorm:"column:password"`
	Username        string  `json:"username" gorm:"column:username"`
	EmailVerifiedAt int64   `json:"email_verified_at" gorm:"column:email_verified_at"`
	CreatedAt       int64   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       int64   `json:"updated_at" gorm:"column:updated_at"`
	Wallet          *Wallet `json:"wallet,omitempty" `
	Version         int     `json:"version" gorm:"column:version"`
}
