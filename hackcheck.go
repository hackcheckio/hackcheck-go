package hackcheck

import "net/http"

const baseUrl string = "https://api.hackcheck.io/api/v2/lookup/"

type Hackcheck struct {
	Apikey           string
	CurrentRatelimit int
	AllowedRatelimit int
	http             *http.Client
}

func New(apikey string) *Hackcheck {
	hc := Hackcheck{
		Apikey:           apikey,
		CurrentRatelimit: 0,
		AllowedRatelimit: 0,
		http:             &http.Client{},
	}

	return &hc
}
