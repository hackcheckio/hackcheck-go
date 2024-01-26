package hackcheck

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Source struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type SearchResult struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	IPAddress   string `json:"ip_address"`
	PhoneNumber string `json:"phone_number"`
	Hash        string `json:"hash"`
	Source      Source `json:"source"`
}

type SearchResponse struct {
	Databases  int                       `json:"databases"`
	Results    []SearchResult            `json:"results"`
	Pagination *SearchResponsePagination `json:"pagination"`
}

type CheckResponse struct {
	Found bool `json:"found"`
}

type SearchResponsePagination struct {
	DocumentCount int `json:"document_count"`
	Next          *struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"next"`
	Prev *struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"prev"`
}

type SearchFilterType = string

const (
	SearchFilterTypeUse    = "use"
	SearchFilterTypeIgnore = "ignore"
)

type SearchFilterOptions struct {
	Type      SearchFilterType
	Databases []string
}

type SearchPaginationOptions struct {
	Offset int
	Limit  int
}

type SearchOptions struct {
	Field      SearchField
	Query      string
	Filter     *SearchFilterOptions
	Pagination *SearchPaginationOptions
}

type CheckOptions struct {
	Field SearchField
	Query string
}

type SearchField = string

const (
	SearchFieldEmail       SearchField = "email"
	SearchFieldUsername    SearchField = "username"
	SearchFieldFullName    SearchField = "full_name"
	SearchFieldPassword    SearchField = "password"
	SearchFieldIPAddress   SearchField = "ip_address"
	SearchFieldPhoneNumber SearchField = "phone_number"
	SearchFieldDomain      SearchField = "domain"
	SearchFieldHash        SearchField = "hash"
)

func (h *HackCheckClient) Search(options *SearchOptions) (*SearchResponse, error) {
	resp, err := h.request(http.MethodGet, h.getSearchUrl(options), nil)
	if err != nil {
		return nil, err
	}

	var r SearchResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HackCheckClient) Check(options *CheckOptions) (bool, error) {
	resp, err := h.request(http.MethodGet, EndpointCheck(h.Apikey, options.Field, options.Query), nil)
	if err != nil {
		return false, err
	}

	var r CheckResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return false, err
	}

	return r.Found, nil
}

func (h *HackCheckClient) getSearchUrl(options *SearchOptions) string {
	thyUrl := EndpointSearch(h.Apikey, options.Field, options.Query)

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
