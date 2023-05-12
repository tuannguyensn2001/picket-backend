package dto

type IncreaseBalanceWalletInput struct {
	UserId int
	Amount int `json:"amount" form:"amount" binding:"required"`
}
