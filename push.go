package pushgate

import "fmt"

// RetryableError indicates a temporary problem that may be resolved
// by retrying.
type RetryableError struct {
	Cause error
}

func (e *RetryableError) Error() string {
	return fmt.Sprintf("retryable error: %v", e.Cause)
}

// Message is a basic push message with title and body
type Message struct {
	Title string
	Body  string
}
