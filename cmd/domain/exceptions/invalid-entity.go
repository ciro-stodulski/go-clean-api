package domainexceptions

func InvalidEntity() *ApplicationException {
	return new("INVALID_ENTITY", "invalid entity")
}
