package hackcheck

import (
	"context"
	"net/http"

	lru "github.com/hashicorp/golang-lru"
)

const baseUrl string = "https://api.hackcheck.io/v3/lookup"

type Hackcheck struct {
	Apikey           string
	CurrentRatelimit int
	AllowedRatelimit int
	limit            int
	ctx              context.Context
	http             *http.Client
	cache            *lru.Cache
}

type HackcheckOption func(*Hackcheck)

func New(apikey string, options ...HackcheckOption) *Hackcheck {
	hc := &Hackcheck{
		Apikey:           apikey,
		CurrentRatelimit: 0,
		AllowedRatelimit: 0,
		http:             &http.Client{},
		ctx:              context.Background(),
	}

	for _, o := range options {
		o(hc)
	}

	return hc
}

// WithHttp will set the Hackcheck http client
func WithHTTP(client *http.Client) func(hc *Hackcheck) {
	return func(hc *Hackcheck) {
		hc.http = client
	}
}

// WithHttp will set the Hackcheck context
func WithContext(ctx context.Context) func(hc *Hackcheck) {
	return func(hc *Hackcheck) {
		hc.ctx = ctx
	}
}

// WithCache will set the Hackcheck cache
func WithCache(size int) func(hc *Hackcheck) {
	return func(hc *Hackcheck) {
		cache, _ := lru.New(size)
		hc.cache = cache
	}
}

// WithLimit will set the Hackcheck result limit
// Should not be used, since its not fully implemented yet.
func WithLimit(limit int) func(hc *Hackcheck) {
	return func(hc *Hackcheck) {
		hc.limit = limit
	}
}
