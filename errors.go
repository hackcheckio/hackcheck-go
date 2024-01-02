package hackcheck

import "errors"

type RateLimitError struct {
	Limit             int
	RemainingRequests int
	message           string
}

func (e *RateLimitError) Error() string {
	return e.message
}

func newRateLimitError(limit, remainingRequests int, message string) *RateLimitError {
	return &RateLimitError{
		Limit:             limit,
		RemainingRequests: remainingRequests,
		message:           message,
	}
}

var (
	ErrUnauthorizedIPAddress = errors.New("unauthorized ip address")
	ErrInvalidAPIKey         = errors.New("invalid api key")

	ErrServerError = errors.New("server returned an error")
)
