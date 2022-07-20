package client

type CacheClient interface {
	Set(key string, value string, timeEx int) error
	Get(key string) (string, error)
}
