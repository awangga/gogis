/**
 * Author: Rolly Maulana Awangga
 * File: gogis.go
 */

package gogis

import (
	"encoding/json"
	"io"
	"net/http"
)

type IpApi struct {
	Status      string `json:"status,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Region      string `json:"region,omitempty"`
	RegionName  string `json:"regionName,omitempty"`
	City        string `json:"city,omitempty"`
	Lat         string `json:"lat,omitempty"`
	Lon         string `json:"lon,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	Isp         string `json:"isp,omitempty"`
	Org         string `json:"org,omitempty"`
	As          string `json:"as,omitempty"`
	Query       string `json:"query,omitempty"`
}

// Returns the sum of two numbers
func Add(a int, b int) int {
	return a + b
}

// Returns the difference between two numbers
func Subtract(a int, b int) int {
	return a - b
}

func GetPublicIP() IpApi {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		println(err.Error())
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		println(err.Error())
	}

	var ip IpApi
	json.Unmarshal(body, &ip)
	return ip
}
