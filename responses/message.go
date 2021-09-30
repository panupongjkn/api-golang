package responses

type Message struct {
	StatusCode uint              `json:"status_code"`
	Message    string            `json:"message"`
	Fields     map[string]string `json:"fields"`
}
