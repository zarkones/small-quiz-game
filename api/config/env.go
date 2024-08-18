package config

import (
	"errors"
	"net"
	"strconv"
)

// errWrap modifies the passed reference to an error
// object "currentErrors" in order to associate newly
// occurring error with an environment variable.
func errWrap(currentErrors *error, varName string, err error) error {
	*currentErrors = errors.Join(*currentErrors, errors.New(varName+": "+err.Error()))
	return *currentErrors
}

// ValidateEnv would assign default values to
// environment variables where applicable.
// It also assures that all of the variables have the right type & value.
// We're wrapping errors to make environment validation more informative.
func ValidateEnv() (envErrs error) {
	if len(HOST) == 0 {
		HOST = "127.0.0.1"
	}
	trial := net.ParseIP(HOST)
	if trial.To4() == nil {
		errWrap(&envErrs, "HOST", ErrInvalidHost)
	}

	if len(PORT) == 0 {
		PORT = "8080"
	}
	port, err := strconv.Atoi(PORT)
	if err != nil {
		errWrap(&envErrs, "PORT", err)
	}
	if port <= 0 || port > 65535 {
		errWrap(&envErrs, "PORT", ErrInvalidPort)
	}

	return envErrs
}
