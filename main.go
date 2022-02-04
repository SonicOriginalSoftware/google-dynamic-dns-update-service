package main

import (
	"google-dynamic-dns-update-service/lib"
	"log"
	"os"
)

func main() {
	outlog := log.New(os.Stdout, "", log.LstdFlags)
	errlog := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	outlog.Printf("Spinning up service...")

	hostnames, usernames, passwords, ipURL, updateFrequency, client, err := lib.Initialize()
	if err != nil {
		errlog.Fatalf("%v\n", err)
	}

	outlog.Printf(
		"\n  Hostnames: %v\n  API URL: %v\n  Update Frequency: %v",
		hostnames,
		ipURL,
		updateFrequency,
	)

	outlog.Printf("Registering service...")

	interrupt := lib.RegisterInterruptHandler(updateFrequency)
	ticker := lib.RegisterTicker(updateFrequency)

	outlog.Printf("Service registered! Starting...")

	interrupted := false
	for !interrupted {
		select {
		case <-ticker.C:
			lib.Update(client, outlog, errlog, ipURL, hostnames, usernames, passwords)
		case s := <-interrupt:
			outlog.Printf("Service stop requested: %v\n", s)
			interrupted = true
		}
	}

	defer ticker.Stop()
	defer close(interrupt)

	outlog.Printf("Service stopped!\n")
}
