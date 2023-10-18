package network

import (
	"encoding/json"
	"fmt"
	api "github.com/json-m/udm-api"
)

//@GetMapping({ "/api/site/{siteName}/isp/status" })
//@GetMapping({ "/api/site/{siteName}/isp/health" })
//@GetMapping({ "/api/site/{siteName}/isp/health/compact" })

// TODO: should any of these sub-structs be their own types?

type IspControllerIspStatusStruct struct {
	InternetAlerts struct {
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
			} `json:"parameters"`
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
	} `json:"internet_alerts"`
	LatencyMax          int    `json:"latency_max"`
	PingServer          string `json:"ping_server"`
	SpeedtestHistorical []struct {
		DownloadMbps  int    `json:"download_mbps"`
		ID            string `json:"id"`
		InterfaceName string `json:"interface_name"`
		LatencyMs     int    `json:"latency_ms"`
		Time          int64  `json:"time"`
		UploadMbps    int    `json:"upload_mbps"`
	} `json:"speedtest_historical"`
	UplinkStatus struct {
		LatencyThreshold int   `json:"latency_threshold"`
		ReceivedBytes    int64 `json:"received_bytes"`
		Statistics       []struct {
			Latency                 int   `json:"latency"`
			LatencyMax              int   `json:"latency_max"`
			ReceivedBytesRateAvg    int   `json:"received_bytes_rate_avg"`
			Timestamp               int64 `json:"timestamp"`
			TransmittedBytesRateAvg int   `json:"transmitted_bytes_rate_avg"`
		} `json:"statistics"`
		TransmittedBytes int64 `json:"transmitted_bytes"`
	} `json:"uplink_status"`
	Wan2Status struct {
		LatencyThreshold int `json:"latency_threshold"`
		ReceivedBytes    int `json:"received_bytes"`
		Statistics       []struct {
			Latency                 int   `json:"latency"`
			LatencyMax              int   `json:"latency_max"`
			ReceivedBytesRateAvg    int   `json:"received_bytes_rate_avg"`
			Timestamp               int64 `json:"timestamp"`
			TransmittedBytesRateAvg int   `json:"transmitted_bytes_rate_avg"`
		} `json:"statistics"`
		TransmittedBytes int `json:"transmitted_bytes"`
	} `json:"wan2_status"`
	WanStatus struct {
		LatencyThreshold int   `json:"latency_threshold"`
		ReceivedBytes    int64 `json:"received_bytes"`
		Statistics       []struct {
			Latency                 int   `json:"latency"`
			LatencyMax              int   `json:"latency_max"`
			ReceivedBytesRateAvg    int   `json:"received_bytes_rate_avg"`
			Timestamp               int64 `json:"timestamp"`
			TransmittedBytesRateAvg int   `json:"transmitted_bytes_rate_avg"`
		} `json:"statistics"`
		TransmittedBytes int64 `json:"transmitted_bytes"`
	} `json:"wan_status"`
}

type IspControllerIspHealthStruct struct {
	HealthStats []struct {
		HighLatency        bool  `json:"high_latency,omitempty"`
		PacketLoss         bool  `json:"packet_loss,omitempty"`
		Timestamp          int64 `json:"timestamp"`
		Wan2FailoverActive bool  `json:"wan2_failover_active,omitempty"`
		WanDowntime        bool  `json:"wan_downtime,omitempty"`
		NotReported        bool  `json:"not_reported,omitempty"`
	} `json:"health_stats"`
}

type IspControllerIspHealthCompactStruct struct {
	Periods []struct {
		Index       int  `json:"index"`
		PacketLoss  bool `json:"packet_loss,omitempty"`
		NotReported bool `json:"not_reported,omitempty"`
	} `json:"periods"`
}

// IspControllerIspStatus gets the ISP status
func IspControllerIspStatus(c api.Client, site string) (IspControllerIspStatusStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/isp/status", c.Host, site)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return IspControllerIspStatusStruct{}, err
	}

	// decode
	var isp IspControllerIspStatusStruct
	err = json.Unmarshal(resp, &isp)
	if err != nil {
		return IspControllerIspStatusStruct{}, err
	}

	return isp, nil
}

// IspControllerIspHealth gets the ISP health
func IspControllerIspHealth(c api.Client, site string) (IspControllerIspHealthStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/isp/health", c.Host, site)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return IspControllerIspHealthStruct{}, err
	}

	// decode
	var isp IspControllerIspHealthStruct
	err = json.Unmarshal(resp, &isp)
	if err != nil {
		return IspControllerIspHealthStruct{}, err
	}

	return isp, nil
}

// IspControllerIspHealthCompact gets the compact version of ISP health.
// returns some indexed periods to be used... elsewhere?
func IspControllerIspHealthCompact(c api.Client, site string) (IspControllerIspHealthCompactStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/isp/health/compact", c.Host, site)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return IspControllerIspHealthCompactStruct{}, err
	}

	// decode
	var isp IspControllerIspHealthCompactStruct
	err = json.Unmarshal(resp, &isp)
	if err != nil {
		return IspControllerIspHealthCompactStruct{}, err
	}

	return isp, nil
}
