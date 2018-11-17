// To satisfy authboss's mailer interface
package main

import (
	"context"
	"github.com/volatiletech/authboss"
)

type AuthMailer struct {
	*server
}

func (a AuthMailer) Send(ctx context.Context, email authboss.Email) error {
	sender := email.FromName + "<" + email.From + ">"

	recipient := ""
	for k, v := range email.To {
		if k == 0 {
			recipient = recipient + ", "
		}
		recipient = recipient + v
	}

	message := a.mailer.NewMessage(sender, email.Subject, email.TextBody, recipient)
	message.SetHtml(email.HTMLBody)

	_, _, err := a.mailer.Send(message)
	if err != nil {
		checkError(err)
	}

	return err
}
