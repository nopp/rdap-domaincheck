package check

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type domain struct {
	Handle string `json:"handle"`
	Events []struct {
		EventAction string    `json:"eventAction"`
		EventDate   time.Time `json:"eventDate"`
	} `json:"events"`
	Status []string `json:"status"`
}

type domainStatus struct {
	Name       string
	Expiration time.Time
	Status     string
}

// Check - Check expiration date
func Check(domainURL, option string) {

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

	if option == "status" {
		fmt.Println(status.Status)
		os.Exit(0)
	}

	if option == "expiration" {
		fmt.Println(status.Expiration)
		os.Exit(0)
	}
}
