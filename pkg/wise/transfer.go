package wise

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"wiseclient/pkg/wise/models"
)

func (c *Client) CreateTransfer(ctx context.Context, createTransferRequestBody models.CreateTransferRequestBody) (models.CreateTransferResponse, error) {
	url := c.baseUrl.ResolveReference(&url.URL{Path: "v1/transfers"})
	body, err := json.Marshal(createTransferRequestBody)
	if err != nil {
		return models.CreateTransferResponse{}, err
	}
	res, err := c.post(ctx, url, body)
	if err != nil {
		return models.CreateTransferResponse{}, err
	}
	if res.statusCode != http.StatusOK {
		return models.CreateTransferResponse{}, c.buildApiError(res)
	}
	var createTransferResponse models.CreateTransferResponse
	err = json.Unmarshal(res.body, &createTransferResponse)
	if err != nil {
		return models.CreateTransferResponse{}, err
	}
	return createTransferResponse, nil
}
