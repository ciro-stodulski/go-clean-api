package domainnotificationpbgrpc

type (
	NotificationPbGrpc interface {
		Verify(id string) error
	}
)
