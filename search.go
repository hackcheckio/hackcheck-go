package hackcheck

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const searchUrl string = "https://api.hackcheck.io/search"

func (h *HackCheckClient) Search(options *SearchOptions) (*SearchResponse, error) {
	req, err := http.NewRequestWithContext(h.ctx, "GET", h.getUrl(options), nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r searchResponse

	if err := json.Unmarshal(b, &r); err != nil {
		return nil, ErrServerError
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			switch r.Error {
			case "Invalid API key.":
				return nil, ErrInvalidAPIKey
			case "Unauthorized IP address.":
				return nil, ErrUnauthorizedIPAddress
			}
		}

		if resp.StatusCode == 429 {
			limit, lerr := strconv.Atoi(resp.Header.Get("X-HackCheck-Limit"))
			remaining, rerr := strconv.Atoi(resp.Header.Get("X-HackCheck-Remaining"))

			if lerr != nil || rerr != nil {
				return nil, ErrServerError
			}

			return nil, newRateLimitError(limit, remaining, "rate limit reached")
		}

		return nil, ErrServerError
	}

	return &SearchResponse{Results: r.Results, Pagination: r.Pagination, Databases: r.Databases}, nil
}

func (h *HackCheckClient) getUrl(options *SearchOptions) string {
	thyUrl := fmt.Sprintf("%s/%s/%s/%s", searchUrl, h.Apikey, options.Field, options.Query)

	query := url.Values{}

	if options.Filter != nil {
		query.Add("filter", options.Filter.Type)
		query.Add("databases", strings.Join(options.Filter.Databases, ","))
	}

	if options.Pagination != nil {
		query.Add("offset", strconv.Itoa(options.Pagination.Offset))
		query.Add("limit", strconv.Itoa(options.Pagination.Limit))
	}

	if encoded := query.Encode(); encoded != "" {
		thyUrl += "?" + encoded
	}

	return thyUrl
}
