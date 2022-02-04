package main

import (
	"google-dynamic-dns-update-service/lib"
	"log"
	"os"
)

func main() {
	outlog := log.New(os.Stdout, "", log.LstdFlags)
	errlog := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	hostnames, apiURL, frequencyTime, username, password, client, err := lib.Initialize()
	if err != nil {
		errlog.Fatalf("%v\n", err)
	}

	interrupt := lib.RegisterInterruptHandler(frequencyTime)
	ticker := lib.RegisterTicker(frequencyTime)

	interrupted := false
	for !interrupted {
		select {
		case <-ticker.C:
			lib.Update(client, outlog, errlog, apiURL, username, password, hostnames)
		case <-interrupt:
			outlog.Printf("\n - Service stop requested!\n")
			interrupted = true
		}
	}

	outlog.Printf("Service stopped!\n")
}
