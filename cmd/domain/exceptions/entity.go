package domainexceptions

func InvalidEntity() error {
	return New("INVALID_ENTITY", "invalid entity")
}
