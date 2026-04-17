package email

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type EmailSender struct {
	SMTPHost    string
	SMTPPort    int
	SMTPUser    string
	SMTPPass    string
	FromEmail   string
	FromName    string
	FrontendURL string
}

func (s *EmailSender) SendMagicLink(email, token string) error {
	link := fmt.Sprintf("%s/auth/verify?token=%s", s.FrontendURL, token)

	subject := "Вход в ваш аккаунт"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Вход в %s</h2>
			<p>Нажмите на ссылку для входа (действует 15 минут):</p>
			<a href="%s" style="background:#007bff;color:white;padding:10px 20px;text-decoration:none;border-radius:5px;">
				Войти
			</a>
			<p>Если вы не запрашивали вход — просто проигнорируйте это письмо.</p>
			<p>Ссылка: %s</p>
		</body>
		</html>
	`, s.FromName, link, link)

	msg := fmt.Sprintf("From: %s <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n%s",
		s.FromName, s.FromEmail, email, subject, body)

	addr := net.JoinHostPort(s.SMTPHost, fmt.Sprintf("%d", s.SMTPPort))

	conn, err := (&net.Dialer{Timeout: 10 * time.Second}).Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("dial smtp: %w", err)
	}

	_ = conn.SetDeadline(time.Now().Add(20 * time.Second))

	client, err := smtp.NewClient(conn, s.SMTPHost)
	if err != nil {
		_ = conn.Close()
		return fmt.Errorf("create smtp client: %w", err)
	}

	defer func() {
		_ = client.Close()
	}()

	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(&tls.Config{ServerName: s.SMTPHost}); err != nil {
			return fmt.Errorf("starttls: %w", err)
		}
	}

	if ok, _ := client.Extension("AUTH"); ok {
		auth := smtp.PlainAuth("", s.SMTPUser, normalizeSMTPPassword(s.SMTPPass), s.SMTPHost)
		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("smtp auth: %w", err)
		}
	}

	if err := client.Mail(s.FromEmail); err != nil {
		return fmt.Errorf("smtp mail from: %w", err)
	}

	if err := client.Rcpt(email); err != nil {
		return fmt.Errorf("smtp rcpt to: %w", err)
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("smtp data: %w", err)
	}

	if _, err := writer.Write([]byte(msg)); err != nil {
		_ = writer.Close()
		return fmt.Errorf("smtp write message: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("smtp close writer: %w", err)
	}

	if err := client.Quit(); err != nil {
		return fmt.Errorf("smtp quit: %w", err)
	}

	return nil
}

func normalizeSMTPPassword(pass string) string {
	return strings.Join(strings.Fields(pass), "")
}
