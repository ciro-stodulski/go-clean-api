package domainexceptions

func UserAlreadyExists() error {
	return New("USER_ALREADY_EXISTS", "user already exists")
}
