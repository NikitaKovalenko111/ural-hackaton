package email

import (
	"fmt"
	"net/smtp"
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

	addr := fmt.Sprintf("%s:%d", s.SMTPHost, s.SMTPPort)
	auth := smtp.PlainAuth("", s.SMTPUser, s.SMTPPass, s.SMTPHost)

	return smtp.SendMail(addr, auth, s.FromEmail, []string{email}, []byte(msg))
}
