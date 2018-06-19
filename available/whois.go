package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WhoisAPI struct {
	APIKey string
}

type searchResult struct {
	WhoisRecord *detail `json:"WhoisRecord"`
}

type detail struct {
	DataError string `json:"dataError"`
}

func (whois *WhoisAPI) Available(domain string) (bool, error) {
	resp, err := http.Get("https://www.whoisxmlapi.com/whoisserver/WhoisService?apiKey=" + whois.APIKey + "&outputFormat=JSON&domainName=" + domain)
	defer resp.Body.Close()
	if err != nil {
		return false, fmt.Errorf("whois: Failed while querying whois api: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("bighuge: Failed with status: %s", resp.Status)
	}
	var data searchResult
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return false, err
	}
	if data.WhoisRecord == nil {
		return false, fmt.Errorf("Missing WhoisRecord")
	}
	return data.WhoisRecord.DataError != "", nil
}
