package events

import "github.com/gookit/event"

const (
	USER_REGISTER = "user.register"
)

func NewUserRegisterEvent(userId int) event.Event {
	return event.NewBasic(USER_REGISTER, map[string]any{
		"user_id": userId,
	})
}
