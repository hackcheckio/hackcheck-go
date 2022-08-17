package hackcheck

import (
	"fmt"
)

var baseUrl string = "https://api.hackcheck.io/api/v2/lookup/"

func (h *Hackcheck) getUrl(q, data string) string {
	return baseUrl + fmt.Sprintf("%s?key=%v&%s=%s", baseUrl, h.Apikey, q, data)
}

type Hackcheck struct {
	Apikey           string
	CurrentRatelimit int
	AllowedRatelimit int
}

func New(apiKey string) *Hackcheck {
	hc := Hackcheck{Apikey: apiKey, CurrentRatelimit: 0, AllowedRatelimit: 0}
	return &hc
}
