package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Payload map[string]string

func (p Payload) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Payload) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, p)
}

type Notification struct {
	ID        int     `json:"id"`
	From      int     `json:"from"`
	To        int     `json:"to"`
	Type      int     `json:"type"`
	Payload   Payload `json:"payload"`
	Template  string  `json:"template"`
	ReadAt    int     `json:"read_at"`
	CreatedAt int     `json:"created_at"`
	UpdatedAt int     `json:"updated_at"`
	FromUser  *User   `json:"from_user" gorm:"-"`
	ToUser    *User   `json:"to_user" gorm:"-"`
}
