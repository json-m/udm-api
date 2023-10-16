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

// https://172.16.69.1/proxy/network/v2/api/site/default/clients/active?includeTrafficUsage=true&includeUnifiDevices=true

type ClientStruct struct {
	Anomalies                           int                     `json:"anomalies"`
	ApMac                               string                  `json:"ap_mac,omitempty"`
	AssocTime                           int                     `json:"assoc_time"`
	Authorized                          bool                    `json:"authorized"`
	Blocked                             bool                    `json:"blocked"`
	Bssid                               string                  `json:"bssid,omitempty"`
	Ccq                                 int                     `json:"ccq,omitempty"`
	Channel                             int                     `json:"channel,omitempty"`
	ChannelWidth                        string                  `json:"channel_width,omitempty"`
	DhcpendTime                         int                     `json:"dhcpend_time,omitempty"`
	DisplayName                         string                  `json:"display_name"`
	Essid                               string                  `json:"essid,omitempty"`
	Fingerprint                         ClientFingerprintStruct `json:"fingerprint,omitempty"`
	FirstSeen                           int                     `json:"first_seen"`
	FixedApEnabled                      bool                    `json:"fixed_ap_enabled,omitempty"`
	FixedIP                             string                  `json:"fixed_ip,omitempty"`
	GwMac                               string                  `json:"gw_mac"`
	Hostname                            string                  `json:"hostname"`
	ID                                  string                  `json:"id"`
	Idletime                            int                     `json:"idletime,omitempty"`
	IP                                  string                  `json:"ip"`
	Ipv4LeaseExpirationTimestampSeconds int                     `json:"ipv4_lease_expiration_timestamp_seconds,omitempty"`
	IsGuest                             bool                    `json:"is_guest"`
	IsWired                             bool                    `json:"is_wired"`
	LastSeen                            int                     `json:"last_seen"`
	LatestAssocTime                     int                     `json:"latest_assoc_time"`
	LocalDNSRecord                      string                  `json:"local_dns_record,omitempty"`
	LocalDNSRecordEnabled               bool                    `json:"local_dns_record_enabled"`
	Mac                                 string                  `json:"mac"`
	Mimo                                string                  `json:"mimo,omitempty"`
	Name                                string                  `json:"name,omitempty"`
	NetworkID                           string                  `json:"network_id"`
	NetworkName                         string                  `json:"network_name"`
	Noise                               int                     `json:"noise,omitempty"`
	Noted                               bool                    `json:"noted"`
	Oui                                 string                  `json:"oui"`
	PowersaveEnabled                    bool                    `json:"powersave_enabled,omitempty"`
	Radio                               string                  `json:"radio,omitempty"`
	RadioName                           string                  `json:"radio_name,omitempty"`
	RadioProto                          string                  `json:"radio_proto,omitempty"`
	RateImbalance                       int                     `json:"rate_imbalance,omitempty"`
	Rssi                                int                     `json:"rssi,omitempty"`
	RxBytes                             int64                   `json:"rx_bytes"`
	RxBytesR                            int                     `json:"rx_bytes-r"`
	RxPackets                           int                     `json:"rx_packets"`
	RxRate                              int                     `json:"rx_rate,omitempty"`
	Signal                              int                     `json:"signal,omitempty"`
	SiteID                              string                  `json:"site_id"`
	Status                              string                  `json:"status"`
	TxBytes                             int                     `json:"tx_bytes"`
	TxBytesR                            int                     `json:"tx_bytes-r"`
	TxMcsIndex                          int                     `json:"tx_mcs_index,omitempty"`
	TxPackets                           int                     `json:"tx_packets"`
	TxRate                              int                     `json:"tx_rate,omitempty"`
	Type                                string                  `json:"type"`
	UnifiDevice                         bool                    `json:"unifi_device"`
	UplinkMac                           string                  `json:"uplink_mac"`
	Uptime                              int                     `json:"uptime"`
	UseFixedip                          bool                    `json:"use_fixedip"`
	UserID                              string                  `json:"user_id"`
	UsergroupID                         string                  `json:"usergroup_id"`
	VirtualNetworkOverrideEnabled       bool                    `json:"virtual_network_override_enabled"`
	VirtualNetworkOverrideID            string                  `json:"virtual_network_override_id,omitempty"`
	WifiExperienceAverage               int                     `json:"wifi_experience_average,omitempty"`
	WifiExperienceScore                 int                     `json:"wifi_experience_score,omitempty"`
	WifiTxAttempts                      int                     `json:"wifi_tx_attempts,omitempty"`
	WlanconfID                          string                  `json:"wlanconf_id,omitempty"`
	SwPort                              int                     `json:"sw_port,omitempty"`
	WiredRateMbps                       int                     `json:"wired_rate_mbps,omitempty"`
	FixedApMac                          string                  `json:"fixed_ap_mac,omitempty"`
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

func ClientController_History(c api.Client) {
	return
}
