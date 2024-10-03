package email

import (
	"log"

	"github.com/resend/resend-go/v2"
)

type EmailClient struct {
	email *resend.Client
}

func NewEmailClient(e *resend.Client) *EmailClient {
	return &EmailClient{
		email: e,
	}
}

func (e *EmailClient) SendEmail(params resend.SendEmailRequest) {
	_, err := e.email.Emails.Send(&params)
	if err != nil {
		log.Printf("Unable to send email: %v", err)
	}
}
