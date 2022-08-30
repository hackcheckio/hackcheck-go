package hackcheck

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (h *Hackcheck) getUrl(q, data string) string {
	return baseUrl + fmt.Sprintf("%s?key=%v&%s=%s", baseUrl, h.Apikey, q, data)
}

// LookupEmail preforms an email lookup
func (h *Hackcheck) LookupEmail(email string) ([]Result, error) {
	return h.Lookup(FieldEmail, email)
}

// LookupUsername preforms a username lookup
func (h *Hackcheck) LookupUsername(username string) ([]Result, error) {
	return h.Lookup(FieldUsername, username)
}

// LookupName preforms a name lookup
func (h *Hackcheck) LookupName(name string) ([]Result, error) {
	return h.Lookup(FieldName, name)
}

// LookupPassword preforms a password lookup
func (h *Hackcheck) LookupPassword(password string) ([]Result, error) {
	return h.Lookup(FieldPassword, password)
}

// LookupIP preforms an IP lookup
func (h *Hackcheck) LookupIP(ip string) ([]Result, error) {
	return h.Lookup(FieldIP, ip)
}

// LookupPhone preforms a phone lookup
func (h *Hackcheck) LookupPhone(phone string) ([]Result, error) {
	return h.Lookup(FieldPhone, phone)
}

// LookupDomain preforms a domain lookup
func (h *Hackcheck) LookupDomain(domain string) ([]Result, error) {
	return h.Lookup(FieldDomain, domain)
}

// Lookup returns all data breaches for the query
// Does not return an error if there are no data breaches
func (h *Hackcheck) Lookup(field Field, query string) ([]Result, error) {
	req, err := http.NewRequest("GET", h.getUrl(field, query), nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.http.Do(req)
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
