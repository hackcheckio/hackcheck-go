package hackcheck

import (
	"context"
	"net/http"
)

type HackCheckClient struct {
	Apikey string
	ctx    context.Context
	http   *http.Client
}

type HackCheckClientOption func(*HackCheckClient)

func New(apikey string, options ...HackCheckClientOption) *HackCheckClient {
	hc := &HackCheckClient{
		Apikey: apikey,
		http:   &http.Client{},
		ctx:    context.Background(),
	}

	for _, o := range options {
		o(hc)
	}

	return hc
}

// WithHttp will set the HackCheckClients http client
func WithHTTP(client *http.Client) func(hc *HackCheckClient) {
	return func(hc *HackCheckClient) {
		hc.http = client
	}
}

// WithHttp will set the HackCheckClients context
func WithContext(ctx context.Context) func(hc *HackCheckClient) {
	return func(hc *HackCheckClient) {
		hc.ctx = ctx
	}
}
