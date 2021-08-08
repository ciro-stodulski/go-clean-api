package http_service

type HttpClient interface {
	Get(url string) ([]byte, error)
}
