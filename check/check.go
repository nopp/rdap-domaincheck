package check

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type domain struct {
	Handle string `json:"handle"`
	Events []struct {
		EventAction string    `json:"eventAction"`
		EventDate   time.Time `json:"eventDate,omitempty"`
	} `json:"events"`
	Status []string `json:"status"`
}

type domainStatus struct {
	Name       string
	Expiration time.Time
	Status     string
}

// Status - Status from registro.br
func Status(domainURL string) {

	var info domain
	var status domainStatus

	resp, err := http.Get("https://rdap.registro.br/domain/" + domainURL)

	if err != nil {
		panic("RDAP error.")
	}

	_ = json.NewDecoder(resp.Body).Decode(&info)

	for _, event := range info.Events {
		if event.EventAction == "expiration" {
			status.Expiration = event.EventDate
		}
	}

	status.Name = domainURL
	status.Status = info.Status[0]

	fmt.Println(status.Status)
}

// DiffDays - Days between today and expiration date
func DiffDays(domainURL, option string) {

	now := time.Now()
	var info domain
	var status domainStatus

	resp, err := http.Get("https://rdap.registro.br/domain/" + domainURL)

	if err != nil {
		panic("RDAP error.")
	}

	_ = json.NewDecoder(resp.Body).Decode(&info)

	for _, event := range info.Events {
		if event.EventAction == "expiration" {
			status.Expiration = event.EventDate
		}
	}

	status.Name = domainURL
	status.Status = info.Status[0]

	diffDays := int(status.Expiration.Sub(now).Hours() / 24)
	fmt.Println(diffDays)
}

// Date - Expiration date
func Date(domainURL, option string) {

	var info domain
	var status domainStatus

	resp, err := http.Get("https://rdap.registro.br/domain/" + domainURL)

	if err != nil {
		panic("RDAP error.")
	}

	_ = json.NewDecoder(resp.Body).Decode(&info)

	for _, event := range info.Events {
		if event.EventAction == "expiration" {
			status.Expiration = event.EventDate
		}
	}

	fmt.Println(status.Expiration)
}
