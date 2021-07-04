package mail

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
)

// Mailer defines function for sending email
type Mailer interface {
	Send(*EmailContent) error
}

// EmailUser defines email address info
type EmailUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (eu *EmailUser) String() string {
	b, _ := json.Marshal(eu)
	return string(b)
}

// EmailContent defines email content info
type EmailContent struct {
	ID               int64      `json:"id"`
	Subject          string     `json:"subject"`
	FromUser         *EmailUser `json:"from"`
	ToUser           *EmailUser `json:"to"`
	PlainTextContent string     `json:"plaintext_content"`
	HtmlContent      string     `json:"html_content"`
}

func (em *EmailContent) String() string {
	b, _ := json.Marshal(em)
	return string(b)
}

// Validate will check whether the email content is valid
func (em *EmailContent) Validate() error {
	if em == nil || em.FromUser == nil || em.ToUser == nil || em.PlainTextContent == "" {
		return errors.New("wrong content")
	}
	return nil
}

// NewSendgrid creates new Sendgrid client
func NewSendgrid(apiKey string) *Sendgrid {
	client := sendgrid.NewSendClient(apiKey)
	return &Sendgrid{
		ApiKey: apiKey,
		Client: client,
	}
}

// Sendgrid implements logic to send email to destination email address via sendgrid
type Sendgrid struct {
	ApiKey string `json:"api_key"`
	Client *sendgrid.Client
}

// Send will send email based on email content
func (m *Sendgrid) Send(em *EmailContent) error {
	if err := em.Validate(); err != nil {
		return err
	}

	//TODO Un-comment that
	// from := mail.NewEmail(em.FromUser.Name, em.FromUser.Email)
	// subject := em.Subject
	// to := mail.NewEmail(em.ToUser.Name, em.ToUser.Email)
	// plainTextContent := em.PlainTextContent
	// htmlContent := em.HtmlContent
	// message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	// response, err := m.Client.Send(message)
	// if err != nil {
	// 	fmt.Println("Cannot send email due to error: ", err)
	// 	return err
	// }
	// fmt.Printf("Email sent with Response code: %v, Response body: %v, Response headers: %v\n", response.StatusCode, response.Body, response.Headers)
	fmt.Println("Sending email: ", em)
	return nil
}
