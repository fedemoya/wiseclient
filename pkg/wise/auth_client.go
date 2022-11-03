package wise

import (
	"fmt"
	"net/http"
)

type AuthClient struct {
	token      string
	httpClient IHTTPClient
}

func NewAuthClient(token string, httpClient IHTTPClient) *AuthClient {
	return &AuthClient{token: token, httpClient: httpClient}
}

func (a AuthClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))
	return a.httpClient.Do(req)
}
