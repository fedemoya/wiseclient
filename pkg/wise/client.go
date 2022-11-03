package wise

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type httpResponse struct {
	res        *http.Response
	statusCode int
	body       []byte
}

type Client struct {
	baseUrl    *url.URL
	httpClient IHTTPClient
}

func NewClient(baseUrl *url.URL, httpClient IHTTPClient) *Client {
	return &Client{
		baseUrl:    baseUrl,
		httpClient: httpClient,
	}
}

func (c *Client) get(ctx context.Context, url *url.URL) (httpResponse, error) {
	req, err := c.prepareRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return httpResponse{}, err
	}
	return c.do(req)
}

func (c *Client) post(ctx context.Context, url *url.URL, body []byte) (httpResponse, error) {
	req, err := c.prepareRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return httpResponse{}, err
	}
	return c.do(req)
}

func (c *Client) patch(ctx context.Context, url *url.URL, body []byte) (httpResponse, error) {
	req, err := c.prepareRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return httpResponse{}, err
	}
	return c.do(req)
}

func (c *Client) delete(ctx context.Context, url *url.URL) (httpResponse, error) {
	req, err := c.prepareRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return httpResponse{}, err
	}
	return c.do(req)
}

func (c *Client) prepareRequest(ctx context.Context, method string, url *url.URL, body []byte) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Host", url.Host)
	req.Header.Set("Date", time.Now().UTC().Format(time.RFC1123))
	req.Header.Set("Accept", "application/json")

	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	}

	return req, nil
}

func (c *Client) do(req *http.Request) (httpResponse, error) {

	res, err := c.httpClient.Do(req)
	if err != nil {
		return httpResponse{}, fmt.Errorf("failed communicating with wise api: %w", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return httpResponse{}, err
	}

	return httpResponse{
		res,
		res.StatusCode,
		body,
	}, nil
}

func (c *Client) buildApiError(unsuccessfulResponse httpResponse) error {

	if unsuccessfulResponse.body == nil && len(unsuccessfulResponse.body) == 0 {
		return fmt.Errorf("empty error body. status code: %d", unsuccessfulResponse.statusCode)
	}

	var unexpectedError error
	switch unsuccessfulResponse.statusCode {
	case 400, 422:
		var validationError ValidationError
		unexpectedError = json.Unmarshal(unsuccessfulResponse.body, &validationError)
		if unexpectedError == nil {
			return validationError
		}
	case 401, 403:
		var authenticationError AuthenticationError
		unexpectedError = json.Unmarshal(unsuccessfulResponse.body, &authenticationError)
		if unexpectedError == nil {
			return authenticationError
		}
	case 500:
		var systemError SystemError
		unexpectedError = json.Unmarshal(unsuccessfulResponse.body, &systemError)
		if unexpectedError == nil {
			return systemError
		}
	default:
		unexpectedError = fmt.Errorf("unexpected error. status code: %d. error: %s", unsuccessfulResponse.statusCode, unsuccessfulResponse.body)
		return unexpectedError
	}

	return fmt.Errorf("unexpected error. status code: %d", unsuccessfulResponse.statusCode)
}
