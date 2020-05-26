package pushgate

import "github.com/gregdel/pushover"

type PushoverSender struct {
	app *pushover.Pushover
}

func NewPushoverSender(key string) *PushoverSender {
	return &PushoverSender{
		app: pushover.New(key),
	}
}

func (p *PushoverSender) Send(rcpt string, msg *Message) error {
	pMsg := pushover.NewMessageWithTitle(msg.Body, msg.Title)
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
