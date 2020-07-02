package pushgate

import "github.com/gregdel/pushover"

// PushoverSender provides a simple wrapper around the pushover api/lib
type PushoverSender struct {
	app *pushover.Pushover
}

// NewPushoverSender creates a new *PushoverSender using the provided
// api key.
func NewPushoverSender(key string) *PushoverSender {
	return &PushoverSender{
		app: pushover.New(key),
	}
}

func (p *PushoverSender) Send(rcpt string, msg *Message) error {
	pMsg := pushover.NewMessageWithTitle(
		trim(msg.Body, 1024),
		trim(msg.Title, 250),
	)
	pRcpt := pushover.NewRecipient(rcpt)

	_, err := p.app.SendMessage(pMsg, pRcpt)
	if err != nil {
		if err == pushover.ErrHTTPPushover {
			return &RetryableError{Cause: err}
		}
		return err
	}

	return nil
}

// trim s to a max length of limit runes. append ... if trimmed.
func trim(s string, limit int) string {
	if len(s) <= limit {
		return s
	}
	runes := []rune(s)
	if len(runes) <= limit {
		return s
	}
	return string(runes[:limit-3]) + "..."
}
