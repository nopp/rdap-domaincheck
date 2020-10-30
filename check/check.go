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
		EventDate   time.Time `json:"eventDate"`
	} `json:"events"`
	Status []string `json:"status"`
}

// Check - Check expiration date
func Check(domainURL, option string) {

	var domainInfo domain

	resp, err := http.Get("https://rdap.registro.br/domain/" + domainURL)

	if err != nil {
		panic("Can't create new request.")
	}

	_ = json.NewDecoder(resp.Body).Decode(&domainInfo)

	fmt.Println(domainInfo.Events)
}
