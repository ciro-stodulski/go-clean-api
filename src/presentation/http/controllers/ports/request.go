package ports_http

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) Get(name string) string {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value
		}
	}
	return ""
}

type HttpRequest struct {
	Body    interface{}
	Params  Params
	Query   map[string][]string
	Headers map[string][]string
	Next    func()
}
