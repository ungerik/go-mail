package email

import "net/mail"

type Configuration struct {
	Host        string
	Port        uint16
	Username    string
	Password    string
	DefaultFrom mail.Address
}

var Config Configuration

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
		Host:        "smtp.gmail.com",
		Port:        587,
		Username:    loginAddress,
		Password:    password,
		DefaultFrom: mail.Address{"", fromAddress},
	}
	return nil
}
