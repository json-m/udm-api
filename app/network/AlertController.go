package network

import (
	"encoding/json"
	"fmt"
	api "github.com/json-m/udm-api"
)

//@GetMapping({ "/api/site/{siteName}/alert" })
//@PostMapping({ "/api/site/{siteName}/system-log/system-critical-alert" })
//@PostMapping({ "/api/site/{siteName}/system-log/update-alert" })
//@PostMapping({ "/api/site/{siteName}/system-log/client-alert" })
//@GetMapping({ "/api/site/{siteName}/alert/dashboard" })
//@PutMapping({ "/api/site/{siteName}/alert/dashboard/mark-all-as-read" })
//@PutMapping({ "/api/site/{siteName}/alert/dashboard/{alertId}/mark-as-read" })
//@PutMapping(value = { "/api/site/{siteName}/alert/mark-as-read" }, consumes = { "application/json" })
//@GetMapping({ "/api/site/{siteName}/alert/setting" })
//@GetMapping({ "/api/site/{siteName}/system-log/setting" })
//@PutMapping(value = { "/api/site/{siteName}/alert/setting" }, consumes = { "application/json" })
//@PutMapping(value = { "/api/site/{siteName}/system-log/setting" }, consumes = { "application/json" })

type AlertControllerSystemCriticalAlertRequestStruct struct {
	TimestampFrom        int      `json:"timestampFrom"`        // default 0
	TimestampTo          int64    `json:"timestampTo"`          // default 1696831199999
	PageSize             int      `json:"pageSize"`             // default 100
	Categories           []string `json:"categories"`           // "INTERNET", "POWER", "DEVICES", "SYSTEM"
	PageNumber           int      `json:"pageNumber"`           // default 0
	SystemLogDeviceTypes []string `json:"systemLogDeviceTypes"` // "GATEWAYS", "SWITCHES", "ACCESS_POINT", "SMART_POWER", "BUILDING_TO_BUILDING_BRIDGES", "UNIFI_LTE", "NON_NETWORK_DEVICES"
}

type AlertControllerAlertStruct struct {
	Data []struct {
		Category   string `json:"category"`
		ID         string `json:"id"`
		Key        string `json:"key"`
		Message    string `json:"message"`
		MessageRaw string `json:"message_raw"`
		Parameters struct {
			ConsoleName struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"CONSOLE_NAME"`
			ConsoleWithDeviceName struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"CONSOLE_WITH_DEVICE_NAME"`
			Device struct {
				DeviceFingerprintID int    `json:"device_fingerprint_id"`
				FingerprintSource   int    `json:"fingerprint_source"`
				ID                  string `json:"id"`
				Model               string `json:"model"`
				ModelName           string `json:"model_name"`
				Name                string `json:"name"`
			} `json:"DEVICE"`
		} `json:"parameters,omitempty"`
		Severity        string `json:"severity"`
		ShowOnDashboard bool   `json:"show_on_dashboard"`
		Status          string `json:"status"`
		Target          string `json:"target"`
		Timestamp       int64  `json:"timestamp"`
		TitleRaw        string `json:"title_raw"`
		Type            string `json:"type"`
	} `json:"data"`
	PageNumber        int `json:"page_number"`
	TotalElementCount int `json:"total_element_count"`
	TotalPageCount    int `json:"total_page_count"`
}

func AlertController_Alert(c api.Client, site string) ([]byte, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/alert", c.Host, site)

	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AlertController_SystemCriticalAlert gets all system critical alerts
func AlertController_SystemCriticalAlert(c api.Client, site string, data AlertControllerSystemCriticalAlertRequestStruct) (AlertControllerAlertStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/system-log/system-critical-alert", c.Host, site)

	// encode request
	payload, err := json.Marshal(data)
	if err != nil {
		return AlertControllerAlertStruct{}, err
	}

	// create request
	resp, err := c.Api("POST", url, payload)
	if err != nil {
		return AlertControllerAlertStruct{}, err
	}

	// decode
	var alert AlertControllerAlertStruct
	err = json.Unmarshal(resp, &alert)
	if err != nil {
		return AlertControllerAlertStruct{}, err
	}

	return alert, nil
}
