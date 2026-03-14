package domain

type EmailMessage struct {
	To      []string
	From    string
	Subject string
	Message string
}

type EmailSender interface {
	Send(message EmailMessage)
}
