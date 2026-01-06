package mailer

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
)

type Mailer struct {
	host     string
	port     int
	username string
	password string
	sender   string
}

func New(host string, port int, username, password, sender string) Mailer {
	return Mailer{
		host:     host,
		port:     port,
		username: username,
		password: password,
		sender:   sender,
	}
}

func (m Mailer) Send(recipient, subject, body string, attachments map[string][]byte) error {
	boundary := "f46d043c813270fc6b0522c12013"

	// Build the email header
	header := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: multipart/related; boundary=%s\r\n"+
		"\r\n", m.sender, recipient, subject, boundary)

	// Build the email body
	message := fmt.Sprintf("--%s\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
		"\r\n"+
		"%s\r\n", boundary, body)

	// Attachments
	for filename, data := range attachments {
		if filename == "logo.jpg" {
			// Inline logo
			message += fmt.Sprintf("\r\n--%s\r\n"+
				"Content-Type: image/jpeg\r\n"+
				"Content-Transfer-Encoding: base64\r\n"+
				"Content-Disposition: inline; filename=\"%s\"\r\n"+
				"Content-ID: <logo>\r\n"+
				"\r\n"+
				"%s\r\n", boundary, filename, base64.StdEncoding.EncodeToString(data))
		} else {
			// Regular attachment
			message += fmt.Sprintf("\r\n--%s\r\n"+
				"Content-Type: application/octet-stream\r\n"+
				"Content-Transfer-Encoding: base64\r\n"+
				"Content-Disposition: attachment; filename=\"%s\"\r\n"+
				"\r\n"+
				"%s\r\n", boundary, filename, base64.StdEncoding.EncodeToString(data))
		}
	}

	// Close the boundary
	message += fmt.Sprintf("\r\n--%s--\r\n", boundary)

	// Combine header and message
	fullMessage := []byte(header + message)

	auth := smtp.PlainAuth("", m.username, m.password, m.host)
	addr := fmt.Sprintf("%s:%d", m.host, m.port)

	// Extract just the email address for the envelope 'from'
	// m.sender might be "Name <email@example.com>"
	fromEmail := m.sender
	if addr, err := mail.ParseAddress(m.sender); err == nil {
		fromEmail = addr.Address
	}

	// Extract just the email address for the envelope 'to'
	toEmail := recipient
	if addr, err := mail.ParseAddress(recipient); err == nil {
		toEmail = addr.Address
	}

	return smtp.SendMail(addr, auth, fromEmail, []string{toEmail}, fullMessage)
}
