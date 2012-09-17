package email

import (
	"net/mail"
)

var Config = Configuration{
	Port: 587,
}

type Configuration struct {
	Host     string
	Port     uint16
	Username string
	Password string
	From     mail.Address
}

func (self *Configuration) Name() string {
	return "email"
}

func (self *Configuration) Init() error {
	if self.From.Address == "" {
		self.From.Address = self.Username
	}
	_, err := ValidateAddress(self.From.Address)
	return err
}

func (self *Configuration) Close() error {
	return nil
}

func InitGmail(email, password string) error {
	return InitGmailFrom(email, email, password)
}

// InitGmailFrom uses a different fromAddress for sending emails
// than the loginAddress used together with the password for authentication
// The fromAddress has to be verified as a valid sender address in Gmail.
func InitGmailFrom(fromAddress, loginAddress, password string) error {
	if _, err := ValidateAddress(fromAddress); err != nil {
		return err
	}
	if _, err := ValidateAddress(loginAddress); err != nil {
		return err
	}
	Config = Configuration{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: loginAddress,
		Password: password,
		From:     mail.Address{Name: "", Address: fromAddress},
	}
	return nil
}
