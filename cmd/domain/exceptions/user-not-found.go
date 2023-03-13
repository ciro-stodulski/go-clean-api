package domainexceptions

func UserNotFound() error {
	return New("USER_NOT_FOUND", "user not found")
}
