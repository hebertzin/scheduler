package emailprovider

import (
	"net/smtp"
	"strings"

	"github.com/hebertzin/scheduler/internal/domain"
)

type SMPTConfig struct {
	Port     string
	Password string
	Host     string
}

func NewSMPT(port string, password string, host string) *SMPTConfig {
	return &SMPTConfig{
		Port:     port,
		Password: password,
		Host:     host,
	}
}

func (config SMPTConfig) Send(s domain.EmailMessage) {
	auth := smtp.PlainAuth("", s.From, config.Password, config.Host)
	addr := config.Host + ":" + config.Port

	msg := []byte("To: " + strings.Join(s.To, ",") + "\r\n" +
		"Subject: " + s.Subject + "\r\n" +
		"\r\n" +
		s.Message + "\r\n")

	_ = smtp.SendMail(addr, auth, s.From, s.To, msg)

	return
}
