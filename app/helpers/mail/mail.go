package mail

import (
	"echo_shop/config"
	mailUtil "echo_shop/pkg/mail"
	"path"

	"github.com/flosch/pongo2"
)

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, templateName string, tplData map[string]interface{}) error {
	filename := path.Join(config.String("APP.TEMPLATE_DIR"), templateName)
	template := pongo2.Must(pongo2.FromCache(filename))

	body, err := template.Execute(tplData)
	if err != nil {
		return err
	}

	m := &mailUtil.Mail{
		Driver:   config.String("MAIL.DRIVER"),
		Host:     config.String("MAIL.HOST"),
		Port:     config.Int("MAIL.PORT"),
		User:     config.String("MAIL.USERNAME"),
		Password: config.String("MAIL.PASSWORD"),
		FromName: config.String("MAIL.FROM_NAME"),
		MailTo:   mailTo,
		Subject:  subject,
		Body:     body,
	}

	return m.Send()
}
