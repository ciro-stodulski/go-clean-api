package domainexceptions

func UserAlreadyExists() *ApplicationException {
	return new("USER_ALREADY_EXISTS", "user already exists")
}
