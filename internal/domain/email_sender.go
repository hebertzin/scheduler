package domain

type EmailMessage struct {
	To      []string
	From    string
	Subject string
	Message string
}
