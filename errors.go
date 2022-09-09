package hackcheck

import "errors"

var (
	ErrUnauthorized = errors.New("failed to lookup, this may be due to your IP address not being linked, or your api key is invalid")
	ErrServerError  = errors.New("server returned an error, try again in a few minutes")
)
