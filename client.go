package gofreegeoipclient

import (
	"encoding/json"
	"net/http"
	"time"
)

type Location struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

type Locator struct {
	Timeout time.Duration
}

func (locator Locator) Locate(ip string) (Location, error) {
	client := http.Client{
		Timeout: locator.Timeout,
	}
	res, err := client.Get("https://freegeoip.net/json/" + ip)
	if err != nil {
		return Location{}, err
	}
	var ret Location
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&ret); err != nil {
		return Location{}, err
	}
	return ret, nil
}
