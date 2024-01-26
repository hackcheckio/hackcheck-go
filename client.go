package hackcheck

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type HackCheckClient struct {
	Apikey string
	ctx    context.Context
	http   *http.Client
}

type HackCheckClientOption func(*HackCheckClient)

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

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *HackCheckClient) request(method, url string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(h.ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := h.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(respData))

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusUnauthorized:
		var x ErrorResponse
		if err := json.Unmarshal(respData, &x); err != nil {
			return nil, err
		}

		if x.Error == "Invalid API key." {
			return nil, ErrInvalidAPIKey
		} else if x.Error == "Unauthorized IP address." {
			return nil, ErrInvalidAPIKey
		} else {
			return nil, ErrServerError
		}
	case http.StatusTooManyRequests:
		limit, lerr := strconv.Atoi(resp.Header.Get("X-HackCheck-Limit"))
		remaining, rerr := strconv.Atoi(resp.Header.Get("X-HackCheck-Remaining"))

		if lerr != nil || rerr != nil {
			return nil, ErrServerError
		}

		return nil, newRateLimitError(limit, remaining, "rate limit reached")
	case http.StatusBadRequest:
		var x ErrorResponse
		if err := json.Unmarshal(respData, &x); err != nil {
			return nil, err
		}

		return nil, errors.New(x.Error)
	default:
		return nil, ErrServerError
	}

	return respData, nil
}
