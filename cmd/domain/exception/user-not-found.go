package exception

func UserNotFound() *ApplicationException {
	return new("USER_NOT_FOUND", "user not found")
}
