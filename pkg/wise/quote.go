package wise

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"wiseclient/pkg/wise/models"
)

func (c *Client) CreateQuote(ctx context.Context, profileId int64, createQuote models.CreateQuoteRequestBody) (CreateQuoteResponse, error) {
	path := fmt.Sprintf("v3/profiles/%d/quotes", profileId)
	url := c.baseUrl.ResolveReference(&url.URL{Path: path})
	body, err := json.Marshal(createQuote)
	if err != nil {
		return CreateQuoteResponse{}, err
	}
	res, err := c.post(ctx, url, body)
	if err != nil {
		return CreateQuoteResponse{}, err
	}
	if res.statusCode != http.StatusOK {
		return CreateQuoteResponse{}, c.buildApiError(res)
	}
	var createQuoteResponse CreateQuoteResponse
	err = json.Unmarshal(res.body, &createQuoteResponse)
	if err != nil {
		return CreateQuoteResponse{}, err
	}
	return createQuoteResponse, nil
}
