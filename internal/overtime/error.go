package overtime

import "errors"

var (
	ErrOvertimeTooMuch  = errors.New("overtime cannot be more than 3 hours")
	ErrAlreadySubmitted = errors.New("overtime already submitted for the day")
)
