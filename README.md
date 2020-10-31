## rdap-domaincheck

This tool use rdap.registro.br

### Install
```
$ git clone https://github.com/nopp/rdap-domaincheck.git
$ cd rdap-domaincheck
$ go build -o domaincheck
```

### Usage
```
$ ./domaincheck
  -domain string
    	example.com
  -option string
    	expiration - days between today and expiration date
    	status - status from registro.br
```
Examples:
```
$ ./domaincheck -domain uol.com.br -option status
active
```

```
./domaincheck -domain uol.com.br -option expiration
905
```