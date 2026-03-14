package smtp

import (
	"fmt"
	"net/smtp"
	"strings"
)

type SMPTConfig struct {
	Port     string
	Password string
	Host     string
}

type SMPTSendEmail struct {
	From    string
	To      []string
	Message string
	Subject string
}

func NewSMPT(port string, password string, host string) *SMPTConfig {
	return &SMPTConfig{
		Port:     port,
		Password: password,
		Host:     host,
	}
}

func (config SMPTConfig) Send(s SMPTSendEmail) error {
	auth := smtp.PlainAuth("", s.From, config.Password, config.Host)
	addr := config.Host + ":" + config.Port

	msg := []byte("To: " + strings.Join(s.To, ",") + "\r\n" +
		"Subject: " + s.Subject + "\r\n" +
		"\r\n" +
		s.Message + "\r\n")

	err := smtp.SendMail(addr, auth, s.From, s.To, msg)
	if err != nil {
		return fmt.Errorf("some error has been ocurred send email: %w", err)
	}

	return nil
}
