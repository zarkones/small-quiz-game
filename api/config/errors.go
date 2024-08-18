package config

import "errors"

var (
	ErrInvalidPort = errors.New("invalid port range")
	ErrInvalidHost = errors.New("invalid host address")
)
