package twitchirc

import (
	"errors"
)

var (
	ErrDial   = errors.New("Dial was unable to connect")
	ErrReader = errors.New("IRC reader error")
	ErrLogin  = errors.New("IRC Login Authentication error")
)
