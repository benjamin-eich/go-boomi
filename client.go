package boomi

type Client struct {
	baseUrl   string
	accountId string
	username  string
	password  string
}

func NewClient(baseUrl, accountId, username, password string) *Client {
	return &Client{baseUrl, accountId, username, password}
}
