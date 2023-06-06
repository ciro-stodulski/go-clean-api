package exception

func PasswordIncorrect() *ApplicationException {
	return new("PASSWORD_INCORRECT", "password incorrect")
}
