package hipchat

const (
	DefaultBaseUrl = "https://api.hipchat.com/v2"
)

type Client struct {
	Token   string
	BaseUrl string
}

func NewClient(token string) *Client {
	return &Client{
		Token:   token,
		BaseUrl: DefaultBaseUrl,
	}
}
