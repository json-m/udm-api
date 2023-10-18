package network

import (
	"encoding/json"
	"fmt"
	api "github.com/json-m/udm-api"
)

//@GetMapping({ "/api/site/{siteName}/network_status" })

type NetworkStatusControllerStruct struct {
	AverageSatisfaction    int    `json:"average_satisfaction"`
	Health                 string `json:"health"`
	HistoricalSatisfaction []int  `json:"historical_satisfaction"`
	Reasons                []any  `json:"reasons"`
}

// NetworkStatusController_NetworkStatus gets network status
func NetworkStatusController_NetworkStatus(c api.Client, site string) (NetworkStatusControllerStruct, error) {
	url := fmt.Sprintf("%s/proxy/network/v2/api/site/%s/network_status", c.Host, site)

	resp, err := c.Api("GET", url, nil)
	if err != nil {
		return NetworkStatusControllerStruct{}, err
	}

	var n NetworkStatusControllerStruct
	err = json.Unmarshal(resp, &n)
	if err != nil {
		return NetworkStatusControllerStruct{}, err
	}

	return n, nil
}
