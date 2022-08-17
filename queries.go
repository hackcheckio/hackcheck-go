package hackcheck

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func (h *Hackcheck) doLookup(lookupType, data string) ([]Result, error) {
	req, err := http.NewRequest("GET", h.getUrl(lookupType, data), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		return nil, ErrInvalidApikey
	}

	h.AllowedRatelimit, _ = strconv.Atoi(resp.Header.Get("hc-allowed-rate"))
	h.CurrentRatelimit, _ = strconv.Atoi(resp.Header.Get("hc-current-rate"))

	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)

	var r Response

	if err := json.Unmarshal(b, &r); err != nil {
		return nil, ErrServerError
	}

	if !r.Success {
		return nil, errors.New(r.Message)
	}

	return r.Results, nil
}

func (h *Hackcheck) LookupEmail(email string) ([]Result, error) {
	return h.doLookup("email", email)
}

func (h *Hackcheck) LookupUsername(username string) ([]Result, error) {
	return h.doLookup("username", username)
}

func (h *Hackcheck) LookupName(name string) ([]Result, error) {
	return h.doLookup("name", name)
}

func (h *Hackcheck) LookupPassword(password string) ([]Result, error) {
	return h.doLookup("password", password)
}

func (h *Hackcheck) LookupIP(ip string) ([]Result, error) {
	return h.doLookup("ip", ip)
}

func (h *Hackcheck) LookupPhone(phone string) ([]Result, error) {
	return h.doLookup("phone", phone)
}

func (h *Hackcheck) LookupDomain(domain string) ([]Result, error) {
	return h.doLookup("domain", domain)
}
