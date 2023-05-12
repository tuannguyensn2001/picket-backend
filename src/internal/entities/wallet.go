package entities

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	Id        int   `json:"id" gorm:"column:id"`
	UserId    int   `json:"user_id" gorm:"column:user_id"`
	Balance   int   `json:"balance" gorm:"column:balance"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

func (p *Wallet) BeforeSave(tx *gorm.DB) error {
	// Check if the record has been modified since it was loaded
	if tx.Statement.Changed("UpdatedAt") {
		// Record has been modified, abort the save operation
		return fmt.Errorf("record has been modified by another user")
	}
	// Record has not been modified, update the UpdatedAt field
	p.UpdatedAt = time.Now().Unix()
	return nil
}
