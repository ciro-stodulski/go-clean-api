package domainexceptions

func PasswordIncorrect() error {
	return New("PASSWORD_INCORRECT", "password incorrect")
}
