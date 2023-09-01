package main

import "errors"

var (
	ErrorInvalidAudio    error = errors.New("toast: invalid audio")
	ErrorInvalidDuration error = errors.New("toast: invalid duration")
)
