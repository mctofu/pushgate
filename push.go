package pushgate

import "fmt"

type RetryableError struct {
	Cause error
}

func (e *RetryableError) Error() string {
	return fmt.Sprintf("Retryable error: %v", e.Cause)
}

// Message is a basic push message with title and body
type Message struct {
	Title string
	Body  string
}
