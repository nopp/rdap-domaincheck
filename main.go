package main

import (
	"flag"
	"os"
	"rdap-expiration/check"
)

func main() {

	option := flag.String("option", "", "expiration - days between today and expiration date \nstatus - status from registro.br")
	domainURL := flag.String("domain", "", "example.com")
	flag.Parse()

	if (*domainURL == "") || (*option == "") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch *option {
	case "status":
		check.CheckStatus(*domainURL)
	case "expiration":
		check.CheckDiffDays(*domainURL, *option)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
