package events

import (
	"github.com/gookit/event"
	"github.com/rs/zerolog/log"
)

func HandlerUserRegisterEvent() {
	event.On(USER_REGISTER, event.ListenerFunc(func(e event.Event) error {
		log.Info().Interface("event", e.Data()).Send()
		return nil
	}))
}
