package hackcheck

var (
	EndpointBase = "https://api.hackcheck.io/"

	EndpointSearch = func(apiKey, field, query string) string {
		return EndpointBase + "search/" + apiKey + "/" + field + "/" + query
	}
	EndpointCheck = func(apiKey, field, query string) string {
		return EndpointBase + "search/check/" + apiKey + "/" + field + "/" + query
	}

	EndpointGetMonitors = func(apiKey string) string {
		return EndpointBase + "monitors/" + apiKey + "/list"
	}
	EndpointGetMonitor = func(apiKey, monitorID string) string {
		return EndpointBase + "monitors/" + apiKey + "/list/" + monitorID
	}

	EndpointGetAssetMonitorSources = func(apiKey, monitorID string) string {
		return EndpointBase + "monitors/" + apiKey + "/sources/asset/" + monitorID
	}
	EndpointGetDomainMonitorSources = func(apiKey, monitorID string) string {
		return EndpointBase + "monitors/" + apiKey + "/sources/domain/" + monitorID
	}

	EndpointUpdateAssetMonitor = func(apiKey, monitorID string) string {
		return EndpointBase + "monitors/" + apiKey + "/update-asset/" + monitorID
	}
	EndpointUpdateDomainMonitor = func(apiKey, monitorID string) string {
		return EndpointBase + "monitors/" + apiKey + "/update-domain/" + monitorID
	}

	EndpointTogglePauseAssetMonitor = func(apiKey, monitorID string) string {
		return EndpointBase + "monitors/" + apiKey + "/pause-asset/" + monitorID
	}
	EndpointTogglePauseDomainMonitor = func(apiKey, monitorID string) string {
		return EndpointBase + "monitors/" + apiKey + "/pause-domain/" + monitorID
	}
)
