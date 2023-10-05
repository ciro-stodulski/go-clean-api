package cacheclient

type (
	CacheClient interface {
		Set(key string, value any, timeEx int) error
		Get(key string) (any, error)
	}
)
