package main

import (
	"net"
	"sync"

	goTLD "github.com/jpillora/go-tld"
)

// concurrencyLimit limits the maximum number of concurrent reading process tasks
const concurrencyLimit = 50

// Row represents a row
type Row struct {
	Subdomain string
	Domain    string
	TLD       string
	IPRecords []net.IP
}

// Process processes all rows
// code is from https://github.com/google/gops/blob/5d514cabbb21de80cb27f529555ee090b60bbb80/goprocess/gp.go#L29
func Process() []Row {
	domains := Domains()

	input := make(chan string, len(domains))
	output := make(chan Row, len(domains))

	for _, domain := range domains {
		input <- domain
	}
	close(input)

	var wg sync.WaitGroup
	wg.Add(concurrencyLimit)

	// Run concurrencyLimit of workers until there
	// is no more processes to be checked in the input channel.
	for i := 0; i < concurrencyLimit; i++ {
		go func() {
			defer wg.Done()

			for row := range input {
				subdomain, domain, tld, IPRecords, ok := processDomain(row)
				if !ok {
					continue
				}
				output <- Row{
					Subdomain: subdomain,
					Domain:    domain,
					TLD:       tld,
					IPRecords: IPRecords,
				}
			}
		}()
	}

	wg.Wait()     // wait until all workers are finished
	close(output) // no more results to be waited for

	var results []Row
	for p := range output {
		results = append(results, p)
	}
	return results
}

// processDomain processes a row
func processDomain(row string) (subdomain, domain, tld string, IPRecords []net.IP, ok bool) {
	u, err := goTLD.Parse("http://" + row)
	if err != nil {
		ok = false
		return
	}

	var records []net.IP

	ipRecords, err := net.LookupIP(row)
	if err == nil {
		for _, ip := range ipRecords {
			records = append(records, ip)
		}
	}

	return u.Subdomain, u.Domain, u.TLD, records, true
}
