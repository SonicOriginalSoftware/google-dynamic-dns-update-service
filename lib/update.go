package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type updateResult struct {
	status string
	err    error
}

func updateService(
	client *http.Client,
	r chan updateResult,
	outlog *log.Logger,
	newIP, hostname, ipURL, username, password string,
) {
	var request *http.Request
	var err error

	if request, err = http.NewRequest("GET", "https://domains.google.com/nic/update", nil); err != nil {
		r <- updateResult{"Could not generate request", err}
		return
	}

	request.SetBasicAuth(username, password)

	query := request.URL.Query()
	query.Set("hostname", hostname)
	query.Set("myip", newIP)

	request.URL.RawQuery = query.Encode()

	outlog.Printf("\nRequest: %v\n", request.URL)

	response, err := client.Do(request)
	if err != nil {
		r <- updateResult{"Could not fetch update request", err}
		return
	}

	update, err := ioutil.ReadAll(response.Body)
	if err != nil {
		r <- updateResult{"Could not read update response", err}
		return
	}

	defer response.Body.Close()

	updateStatus := string(update)

	if updateStatus != fmt.Sprintf("good %v", newIP) {
		err = fmt.Errorf("Could not update dynamic DNS")
	}

	r <- updateResult{fmt.Sprintf("  Status: %v : %v\n", response.Status, updateStatus), err}
}

// Update sends requests to update all Google Dynamic DNS hostnames with the current IP
func Update(
	client *http.Client,
	outlog, errlog *log.Logger,
	ipURL, username, password string,
	hostnames []string,
) {
	newIP, err := Get(ipURL)
	if err != nil {
		errlog.Printf("%v:  %v\n", "Could not get IP address for updating", err.Error())
		return
	}

	result := make(chan updateResult, len(hostnames))

	for _, hostname := range hostnames {

		go updateService(
			client,
			result,
			outlog,
			string(newIP),
			hostname,
			ipURL,
			username,
			password,
		)
	}

	var r updateResult
	select {
	case result <- r:
		if r.err != nil {
			errlog.Printf("%v:  %v\n", r.status, err.Error())
		}
	}
}
