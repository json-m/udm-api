package network

import (
	"encoding/json"
	"fmt"
	api "github.com/json-m/udm-api"
)

//@GetMapping({ "/api/site/{siteName}/traffic" })
//@GetMapping({ "/api/site/{siteName}/traffic/{mac}" })
//@GetMapping({ "/api/site/{siteName}/traffic-rate" })
//@PostMapping({ "/api/site/{siteName}/app-traffic-rate" })

// /network/default/securityInsights/trafficIdentification

type TrafficOverviewControllerStruct struct {
	ClientUsageByApp []ClientUsageByAppStruct `json:"client_usage_by_app"`
	TotalUsageByApp  []UsageByAppStruct       `json:"total_usage_by_app"`
}

type ClientUsageByAppStruct struct {
	Client struct {
		Fingerprint ClientFingerprintStruct `json:"fingerprint"`
		Hostname    string                  `json:"hostname"`
		IsWired     bool                    `json:"is_wired"`
		Mac         string                  `json:"mac"`
		Name        string                  `json:"name"`
		Oui         string                  `json:"oui"`
		WlanconfID  string                  `json:"wlanconf_id"`
	} `json:"client,omitempty"`
	UsageByApp []UsageByAppStruct `json:"usage_by_app"`
}

type UsageByAppStruct struct {
	Application      int   `json:"application,omitempty"`
	BytesReceived    int64 `json:"bytes_received,omitempty"`
	BytesTransmitted int64 `json:"bytes_transmitted,omitempty"`
	Category         int   `json:"category,omitempty"`
	ClientCount      int   `json:"client_count,omitempty"`
	TotalBytes       int64 `json:"total_bytes,omitempty"`
	ActivitySeconds  int   `json:"activity_seconds,omitempty"`
}

type AppTrafficRateStruct struct {
	IntervalSeconds int   `json:"interval_seconds"`
	RxByteR         int   `json:"rx_byte-r"`
	Timestamp       int64 `json:"timestamp"`
	TotalBytes      int   `json:"total_bytes"`
	TxByteR         int   `json:"tx_byte-r"`
	TopApp          struct {
		Application int `json:"application"`
		Category    int `json:"category"`
		RxByteR     int `json:"rx_byte-r"`
		TotalBytes  int `json:"total_bytes"`
		TxByteR     int `json:"tx_byte-r"`
	} `json:"top_app,omitempty"`
}

// TrafficOverviewController_Traffic gets traffic overview
// TODO: takes int64 utc time, should it? or would it be more proper to take a time.Time?
func TrafficOverviewController_Traffic(c api.Client, site string, start, end int64) (TrafficOverviewControllerStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/traffic?start=%d&end=%d", c.Host, site, start, end)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return TrafficOverviewControllerStruct{}, err
	}

	// decode response
	var to TrafficOverviewControllerStruct
	err = json.Unmarshal(resp, &to)
	if err != nil {
		return TrafficOverviewControllerStruct{}, err
	}

	return to, nil
}

// TrafficOverviewController_TrafficMac gets traffic overview for a specific mac
func TrafficOverviewController_TrafficMac(c api.Client, site, mac string, start, end int64) (TrafficOverviewControllerStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/traffic/%s?start=%d&end=%d", c.Host, site, mac, start, end)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return TrafficOverviewControllerStruct{}, err
	}

	// decode response
	var to TrafficOverviewControllerStruct
	err = json.Unmarshal(resp, &to)
	if err != nil {
		return TrafficOverviewControllerStruct{}, err
	}

	return to, nil
}

// TrafficOverviewController_TrafficRate gets traffic rate(?).
// i can't actually find an example of this in use in the UI?
func TrafficOverviewController_TrafficRate(c api.Client, site string, start, end int64) ([]AppTrafficRateStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/traffic-rate?start=%d&end=%d&includeUnidentified=true", c.Host, site, start, end)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return []AppTrafficRateStruct{}, err
	}

	// decode response
	var to []AppTrafficRateStruct
	err = json.Unmarshal(resp, &to)
	if err != nil {
		return []AppTrafficRateStruct{}, err
	}

	return to, nil
}

// TrafficOverviewController_AppTrafficRate gets app traffic rate
func TrafficOverviewController_AppTrafficRate(c api.Client, site string, start, end int64) ([]AppTrafficRateStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/app-traffic-rate?start=%d&end=%d&includeUnidentified=true", c.Host, site, start, end)

	// create request
	resp, err := c.Api("POST", url, []byte(`{}`))
	if err != nil {
		return []AppTrafficRateStruct{}, err
	}

	// decode response
	var to []AppTrafficRateStruct
	err = json.Unmarshal(resp, &to)
	if err != nil {
		return []AppTrafficRateStruct{}, err
	}

	return to, nil
}
