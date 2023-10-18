package network

import (
	"encoding/json"
	"fmt"
	api "github.com/json-m/udm-api"
)

//@GetMapping({ "/api/site/{siteName}/clients/active" })
//@GetMapping({ "/api/site/{siteName}/clients/active/{clientMac}" })
//@GetMapping({ "/api/site/{siteName}/clients/{clientMac}/power" })
//@GetMapping({ "/api/site/{siteName}/clients/local/{clientMac}" })
//@GetMapping({ "/api/site/{siteName}/clients/history" })

// /proxy/network/v2/api/site/default/clients/active?includeTrafficUsage=true&includeUnifiDevices=true

// ClientStruct struct for all ClientController communication
type ClientStruct struct {
	Anomalies                           int                     `json:"anomalies,omitempty"`
	ApMac                               string                  `json:"ap_mac,omitempty"`
	AssocTime                           int                     `json:"assoc_time,omitempty"`
	Authorized                          bool                    `json:"authorized,omitempty"`
	Blocked                             bool                    `json:"blocked,omitempty"`
	Bssid                               string                  `json:"bssid,omitempty"`
	Ccq                                 int                     `json:"ccq,omitempty"`
	Channel                             int                     `json:"channel,omitempty"`
	ChannelWidth                        string                  `json:"channel_width,omitempty"`
	DhcpendTime                         int                     `json:"dhcpend_time,omitempty"`
	DisplayName                         string                  `json:"display_name,omitempty"`
	Essid                               string                  `json:"essid,omitempty"`
	Fingerprint                         ClientFingerprintStruct `json:"fingerprint,omitempty"`
	FirstSeen                           int                     `json:"first_seen,omitempty"`
	FixedApEnabled                      bool                    `json:"fixed_ap_enabled,omitempty"`
	FixedIP                             string                  `json:"fixed_ip,omitempty"`
	GwMac                               string                  `json:"gw_mac,omitempty"`
	Hostname                            string                  `json:"hostname,omitempty"`
	ID                                  string                  `json:"id,omitempty"`
	Idletime                            int                     `json:"idletime,omitempty"`
	IP                                  string                  `json:"ip,omitempty"`
	Ipv4LeaseExpirationTimestampSeconds int                     `json:"ipv4_lease_expiration_timestamp_seconds,omitempty"`
	IsGuest                             bool                    `json:"is_guest,omitempty"`
	IsWired                             bool                    `json:"is_wired,omitempty"`
	LastSeen                            int                     `json:"last_seen,omitempty"`
	LatestAssocTime                     int                     `json:"latest_assoc_time,omitempty"`
	LocalDNSRecord                      string                  `json:"local_dns_record,omitempty,omitempty"`
	LocalDNSRecordEnabled               bool                    `json:"local_dns_record_enabled,omitempty"`
	Mac                                 string                  `json:"mac,omitempty"`
	Mimo                                string                  `json:"mimo,omitempty,omitempty"`
	Name                                string                  `json:"name,omitempty,omitempty"`
	NetworkID                           string                  `json:"network_id,omitempty"`
	NetworkName                         string                  `json:"network_name,omitempty"`
	Noise                               int                     `json:"noise,omitempty,omitempty"`
	Noted                               bool                    `json:"noted,omitempty"`
	Oui                                 string                  `json:"oui,omitempty"`
	PowersaveEnabled                    bool                    `json:"powersave_enabled,omitempty"`
	Radio                               string                  `json:"radio,omitempty"`
	RadioName                           string                  `json:"radio_name,omitempty"`
	RadioProto                          string                  `json:"radio_proto,omitempty"`
	RateImbalance                       int                     `json:"rate_imbalance,omitempty"`
	Rssi                                int                     `json:"rssi,omitempty"`
	RxBytes                             int64                   `json:"rx_bytes,omitempty"`
	RxBytesR                            int                     `json:"rx_bytes-r,omitempty"`
	RxPackets                           int                     `json:"rx_packets,omitempty"`
	RxRate                              int                     `json:"rx_rate,omitempty,omitempty"`
	Signal                              int                     `json:"signal,omitempty,omitempty"`
	SiteID                              string                  `json:"site_id,omitempty"`
	Status                              string                  `json:"status,omitempty"`
	TxBytes                             int                     `json:"tx_bytes,omitempty"`
	TxBytesR                            int                     `json:"tx_bytes-r,omitempty"`
	TxMcsIndex                          int                     `json:"tx_mcs_index,omitempty,omitempty"`
	TxPackets                           int                     `json:"tx_packets,omitempty"`
	TxRate                              int                     `json:"tx_rate,omitempty,omitempty"`
	Type                                string                  `json:"type,omitempty"`
	UnifiDevice                         bool                    `json:"unifi_device,omitempty"`
	UplinkMac                           string                  `json:"uplink_mac,omitempty"`
	Uptime                              int                     `json:"uptime,omitempty"`
	UseFixedip                          bool                    `json:"use_fixedip,omitempty"`
	UserID                              string                  `json:"user_id,omitempty"`
	UsergroupID                         string                  `json:"usergroup_id,omitempty"`
	VirtualNetworkOverrideEnabled       bool                    `json:"virtual_network_override_enabled,omitempty"`
	VirtualNetworkOverrideID            string                  `json:"virtual_network_override_id,omitempty"`
	WifiExperienceAverage               int                     `json:"wifi_experience_average,omitempty"`
	WifiExperienceScore                 int                     `json:"wifi_experience_score,omitempty"`
	WifiTxAttempts                      int                     `json:"wifi_tx_attempts,omitempty"`
	WlanconfID                          string                  `json:"wlanconf_id,omitempty"`
	SwPort                              int                     `json:"sw_port,omitempty"`
	WiredRateMbps                       int                     `json:"wired_rate_mbps,omitempty"`
	FixedApMac                          string                  `json:"fixed_ap_mac,omitempty"`
	LastConnectionNetworkName           string                  `json:"last_connection_network_name,omitempty"`
	LastIP                              string                  `json:"last_ip,omitempty"`
	LastConnectionNetworkID             string                  `json:"last_connection_network_id,omitempty"`
	UnifiDeviceInfo                     struct {
		IconFilename    string  `json:"icon_filename,omitempty"`
		IconResolutions [][]int `json:"icon_resolutions,omitempty"`
	} `json:"unifi_device_info,omitempty"`
}

type ClientFingerprintStruct struct {
	ComputedDevID  int  `json:"computed_dev_id"`
	ComputedEngine int  `json:"computed_engine"`
	Confidence     int  `json:"confidence"`
	DevCat         int  `json:"dev_cat"`
	DevFamily      int  `json:"dev_family"`
	DevID          int  `json:"dev_id"`
	DevVendor      int  `json:"dev_vendor"`
	HasOverride    bool `json:"has_override"`
	OsName         int  `json:"os_name"`
}

type ClientPowerStruct struct {
	CurrentAmperes float64 `json:"current_amperes,omitempty"`
	PowerWatts     float64 `json:"power_watts,omitempty"`
	VoltageVolts   float64 `json:"voltage_volts,omitempty"`
}

// ClientController_Active gets all active clients
func ClientController_Active(c api.Client, site string) ([]ClientStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/clients/active", c.Host, site)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return []ClientStruct{}, err
	}

	// decode response
	var cs []ClientStruct
	err = json.Unmarshal(resp, &cs)
	if err != nil {
		return []ClientStruct{}, err
	}

	return cs, nil
}

// ClientController_ActiveMac gets an active client by mac address
func ClientController_ActiveMac(c api.Client, site string, mac string) (ClientStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/clients/active/%s", c.Host, site, mac)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return ClientStruct{}, err
	}

	// decode response
	var cs ClientStruct
	err = json.Unmarshal(resp, &cs)
	if err != nil {
		return ClientStruct{}, err
	}

	return cs, nil
}

// ClientController_ClientPower gets the PoE information of a client (restricted to UniFi PoE devices that expose this information)
func ClientController_ClientPower(c api.Client, site string, mac string) (ClientPowerStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/clients/%s/power", c.Host, site, mac)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return ClientPowerStruct{}, err
	}

	// decode response
	var cs ClientPowerStruct
	err = json.Unmarshal(resp, &cs)
	if err != nil {
		return ClientPowerStruct{}, err
	}

	return cs, nil
}

// ClientController_Local gets a client by mac address, similar output to ClientController_Active, but for a specific device.
// i don't really know how this is any different than ClientController_ActiveMac.
// this could definitely use a better name
func ClientController_Local(c api.Client, site string, mac string) (ClientStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/clients/local/%s", c.Host, site, mac)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return ClientStruct{}, err
	}

	// decode response
	var cs ClientStruct
	err = json.Unmarshal(resp, &cs)
	if err != nil {
		return ClientStruct{}, err
	}

	return cs, nil
}

// ClientController_History gets all clients within the last N hours, even offline ones
func ClientController_History(c api.Client, site string, hours int) ([]ClientStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/clients/history?onlyNonBlocked=true&includeUnifiDevices=true&withinHours=%d", c.Host, site, hours)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return []ClientStruct{}, err
	}

	// decode response
	var cs []ClientStruct
	err = json.Unmarshal(resp, &cs)
	if err != nil {
		return []ClientStruct{}, err
	}

	return cs, nil
}
