package domainexceptions

func PasswordIncorrect() *ApplicationException {
	return new("PASSWORD_INCORRECT", "password incorrect")
}
