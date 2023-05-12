package app

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewRawError(message string, code int) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}

func NewBadRequestError(message string) Error {
	return NewRawError(message, 400)
}

func NewForbiddenError(message string) Error {
	return NewRawError(message, 403)
}
