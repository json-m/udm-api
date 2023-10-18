package network

import (
	"encoding/json"
	"fmt"
	api "github.com/json-m/udm-api"
)

//@GetMapping({ "/api/site/{siteName}/poe-info" })

type PoeControllerStruct struct {
	PoeHostDeviceInfo struct {
		AvgPoeUsage             float64 `json:"avg_poe_usage"`
		DisplayName             string  `json:"display_name"`
		Mac                     string  `json:"mac"`
		MaxPoeUsage             float64 `json:"max_poe_usage"`
		NearPowerLimitDownlinks []any   `json:"near_power_limit_downlinks"`
	} `json:"poe_host_device_info"`
}

// PoeController_PoeInfo gets PoE info for a device
func PoeController_PoeInfo(c api.Client, site, mac string) (PoeControllerStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/poe-info?mac=%s", c.Host, site, mac)

	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return PoeControllerStruct{}, err
	}

	var p PoeControllerStruct
	err = json.Unmarshal(resp, &p)
	if err != nil {
		return PoeControllerStruct{}, err
	}
	return p, nil
}
