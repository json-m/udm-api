package network

import (
	"encoding/json"
	"fmt"
	api "github.com/json-m/udm-api"
)

//@GetMapping({ "/api/site/{siteName}/trafficroutes" })
//@PostMapping(value = { "/api/site/{siteName}/trafficroutes" }, consumes = { "application/json" })
//@PutMapping(value = { "/api/site/{siteName}/trafficroutes/{routeId}" }, consumes = { "application/json" })
//@DeleteMapping({ "/api/site/{siteName}/trafficroutes/{routeId}" })

// TrafficRouteStruct json expected from the API for req/resp
type TrafficRouteStruct struct {
	ID             string                            `json:"_id,omitempty"`
	Description    string                            `json:"description,omitempty"`
	Domains        []any                             `json:"domains,omitempty"`
	Enabled        bool                              `json:"enabled,omitempty"`
	IPAddresses    []TrafficRouteIPAddressesStruct   `json:"ip_addresses,omitempty"`
	IPRanges       []any                             `json:"ip_ranges,omitempty"`
	MatchingTarget string                            `json:"matching_target,omitempty"` // "IP", "INTERNET", "DOMAIN", "REGION"
	NetworkID      string                            `json:"network_id,omitempty"`
	NextHop        string                            `json:"next_hop,omitempty"`
	Regions        []string                          `json:"regions,omitempty"` // string slice of 2 letter region codes
	TargetDevices  []TrafficRouteTargetDevicesStruct `json:"target_devices,omitempty"`
}

// TrafficRouteIPAddressesStruct for slices of IP targets
type TrafficRouteIPAddressesStruct struct {
	IPOrSubnet string `json:"ip_or_subnet,omitempty"`
	IPVersion  string `json:"ip_version,omitempty"` // "v4" or "v6"
	PortRanges []any  `json:"port_ranges,omitempty"`
	Ports      []any  `json:"ports,omitempty"`
}

// TrafficRouteTargetDevicesStruct for slices of devices to target
type TrafficRouteTargetDevicesStruct struct {
	ClientMac string `json:"client_mac,omitempty"`
	Type      string `json:"type,omitempty"` // "ALL_CLIENTS", "CLIENT"
}

// TrafficRouteController_GetRoutes gets all traffic routes
func TrafficRouteController_GetRoutes(c api.Client, site string) ([]TrafficRouteStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/trafficroutes", c.Host, site)

	// create request
	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return []TrafficRouteStruct{}, err
	}

	// decode response
	var trr []TrafficRouteStruct
	err = json.Unmarshal(resp, &trr)
	if err != nil {
		return []TrafficRouteStruct{}, err
	}

	return trr, nil
}

// TrafficRouteController_CreateRoute creates a traffic route rule, returns the ID of the new rule
func TrafficRouteController_CreateRoute(c api.Client, site string, data TrafficRouteStruct) (string, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/trafficroutes", c.Host, site)

	// encode data
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// create request
	resp, err := c.Api("POST", url, payload)
	if err != nil {
		return "", err
	}

	// decode response
	var trr TrafficRouteStruct
	err = json.Unmarshal(resp, &trr)
	if err != nil {
		return "", err
	}

	return trr.ID, nil
}

// TrafficRouteController_UpdateRoute updates a traffic route rule
func TrafficRouteController_UpdateRoute(c api.Client, site, id string, data TrafficRouteStruct) error {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/trafficroutes/%s", c.Host, site, id)

	// encode data
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// create request
	_, err = c.Api("PUT", url, payload)
	if err != nil {
		return err
	}

	return nil
}

// TrafficRouteController_DeleteRoute deletes a traffic route rule
func TrafficRouteController_DeleteRoute(c api.Client, site, id string) error {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/trafficroutes/%s", c.Host, site, id)

	// create request
	_, err := c.Api("DELETE", url, nil)
	if err != nil {
		return err
	}

	return nil
}
