package hackcheck

import (
	"encoding/json"
	"net/http"
	"time"
)

type MonitorStatus int

const (
	MonitorStatusRunning = iota
	MonitorStatusPaused
	MonitorStatusExpired
)

type AssetMonitor struct {
	ID                string        `json:"id"`
	Status            MonitorStatus `json:"status"`
	Type              SearchField   `json:"type"`
	Asset             string        `json:"asset"`
	NotificationEmail string        `json:"notification_email"`
	ExpiresSoon       bool          `json:"expires_soon"`
	CreatedAt         time.Time     `json:"created_at"`
	EndsAt            time.Time     `json:"ends_at"`
}

type DomainMonitor struct {
	ID                string        `json:"id"`
	Status            MonitorStatus `json:"status"`
	Domain            string        `json:"domain"`
	NotificationEmail string        `json:"notification_email"`
	ExpiresSoon       bool          `json:"expires_soon"`
	CreatedAt         time.Time     `json:"created_at"`
	EndsAt            time.Time     `json:"ends_at"`
}

type GetMonitorsResponse struct {
	AssetMonitors  []AssetMonitor  `json:"asset_monitors"`
	DomainMonitors []DomainMonitor `json:"domain_monitors"`
}

type UpdateAssetMonitorParams struct {
	Asset             string `json:"asset"`
	Type              string `json:"asset_type"`
	NotificationEmail string `json:"notification_email"`
}

type UpdateDomainMonitorParams struct {
	Domain            string
	NotificationEmail string `json:"notification_email"`
}

func (h *HackCheckClient) GetMonitors() (*GetMonitorsResponse, error) {
	resp, err := h.request(http.MethodGet, EndpointGetMonitors(h.Apikey), nil)
	if err != nil {
		return nil, err
	}

	var r GetMonitorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HackCheckClient) GetAssetMonitor(monitorID string) (*AssetMonitor, error) {
	resp, err := h.request(http.MethodGet, EndpointGetMonitor(h.Apikey, monitorID), nil)
	if err != nil {
		return nil, err
	}

	var r AssetMonitor
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HackCheckClient) GetDomainMonitor(monitorID string) (*DomainMonitor, error) {
	resp, err := h.request(http.MethodGet, EndpointGetMonitor(h.Apikey, monitorID), nil)
	if err != nil {
		return nil, err
	}

	var r DomainMonitor
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HackCheckClient) GetAssetMonitorSources(monitorID string) ([]Source, error) {
	resp, err := h.request(http.MethodGet, EndpointGetAssetMonitorSources(h.Apikey, monitorID), nil)
	if err != nil {
		return nil, err
	}

	var r []Source
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return r, nil
}
func (h *HackCheckClient) GetDomainMonitorSources(monitorID string) ([]Source, error) {
	resp, err := h.request(http.MethodGet, EndpointGetDomainMonitorSources(h.Apikey, monitorID), nil)
	if err != nil {
		return nil, err
	}

	var r []Source
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (h *HackCheckClient) UpdateAssetMonitor(monitorID string, updateParams *UpdateAssetMonitorParams) (*AssetMonitor, error) {
	data, err := json.Marshal(updateParams)
	if err != nil {
		return nil, err
	}

	resp, err := h.request(http.MethodPost, EndpointUpdateAssetMonitor(h.Apikey, monitorID), data)
	if err != nil {
		return nil, err
	}

	var r AssetMonitor
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HackCheckClient) UpdateDomainMonitor(monitorID string, updateParams *UpdateDomainMonitorParams) (*DomainMonitor, error) {
	data, err := json.Marshal(updateParams)
	if err != nil {
		return nil, err
	}

	resp, err := h.request(http.MethodPost, EndpointUpdateDomainMonitor(h.Apikey, monitorID), data)
	if err != nil {
		return nil, err
	}

	var r DomainMonitor
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HackCheckClient) TogglePauseAssetMonitor(monitorID string) (*AssetMonitor, error) {
	resp, err := h.request(http.MethodGet, EndpointTogglePauseAssetMonitor(h.Apikey, monitorID), nil)
	if err != nil {
		return nil, err
	}

	var r AssetMonitor
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (h *HackCheckClient) TogglePauseDomainMonitor(monitorID string) (*DomainMonitor, error) {
	resp, err := h.request(http.MethodGet, EndpointTogglePauseDomainMonitor(h.Apikey, monitorID), nil)
	if err != nil {
		return nil, err
	}

	var r DomainMonitor
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
