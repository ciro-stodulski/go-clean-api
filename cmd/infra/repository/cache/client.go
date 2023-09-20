package cacheclient

type (
	CacheClient interface {
		Set(key string, value interface{}, timeEx int) error
		Get(key string) (interface{}, error)
	}
)
