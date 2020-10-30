package main

import (
	"flag"
	"os"
	"rdap-expiration/check"
)

func main() {
	option := flag.String("option", "", "expiration|status (Required)")
	domainURL := flag.String("domain", "", "example.com (Required)")
	flag.Parse()

	if (*domainURL == "") || (*option == "") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch *option {
	case "expiration", "status":
		check.Check(*domainURL, *option)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

}
