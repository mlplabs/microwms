package client

import (
	"context"
	"fmt"
	httpClient "github.com/mlplabs/app-utils/pkg/http/client"
	"github.com/mlplabs/mwms-core/whs/model"
	"net/http"
)

type HttpClient struct {
	client *httpClient.Client
}

func NewHttpClient(clientName string, cfg *ServiceConfig) *HttpClient {
	return &HttpClient{
		client: httpClient.NewClient(clientName, cfg.Name, cfg.BaseURL),
	}
}

func (c *HttpClient) GetWarehouses(ctx context.Context) ([]model.Warehouse, error) {
	responseData := make([]model.Warehouse, 0)
	_, err := c.client.Get(ctx, fmt.Sprintf("warehouses"), &httpClient.RequestParams{
		ProxyToken:   true,
		ResponseBody: &responseData,
		RequestHandler: func(request *http.Request) *http.Request {
			return request
		},
	})

	if err != nil {
		return nil, err
	}
	return responseData, nil
}
