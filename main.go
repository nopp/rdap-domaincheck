package main

import (
	"flag"
	"os"
	"rdap-expiration/check"
)

func main() {

	option := flag.String("option", "", "expiration - Days between today and expiration date \ndate - Expiration date\nstatus - Status from registro.br")
	domainURL := flag.String("domain", "", "example.com")

	flag.Parse()

	if (*domainURL == "") || (*option == "") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	check.DomainInfo(*domainURL, *option)
	os.Exit(0)
}
