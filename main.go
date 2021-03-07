package main

import (
	"flag"
	"os"
	"sync"

	"bgs/models"
	"bgs/output"
	"bgs/resolve"
	"bgs/wordlist"
)

func main() {
	var wg sync.WaitGroup

	var resolversList string
	flag.StringVar(&resolversList, "r", "", "resolvers list path")

	var clean bool
	flag.BoolVar(&clean, "v", false, "verbose mode.")

	flag.Parse()

	var toResolve []string
	var resolvers []string

	toResolve = wordlist.LoadFromSTDIN()
	resolvers = wordlist.LoadWords(resolversList)

	for _, toTest := range toResolve {
		wg.Add(1)
		for _, dnsResolverIp := range resolvers {
			defer wg.Done()
			go func(t, i string) {
				object := models.Mapper{
					Target: t,
					DnsIp:  i,
				}

				switch clean {
				case false:
					output.LogClean(resolve.Resolve(&object))
				case true:
					output.LogIt(resolve.Resolve(&object))
				default:
					os.Exit(1)
				}
			}(toTest, dnsResolverIp)
		}
	}
	wg.Wait()
}
