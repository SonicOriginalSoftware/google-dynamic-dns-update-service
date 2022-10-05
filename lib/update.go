package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type updateResult struct {
	err    error
	status string
}

func updateService(
	client *http.Client,
	r chan updateResult,
	outlog *log.Logger,
	newIP, hostname, username, password string,
) {
	var request *http.Request
	var err error

	if request, err = http.NewRequest("GET", "https://domains.google.com/nic/update", nil); err != nil {
		r <- updateResult{status: "Could not generate request", err: err}
		return
	}

	request.SetBasicAuth(username, password)

	query := request.URL.Query()
	query.Set("hostname", hostname)
	query.Set("myip", newIP)

	request.URL.RawQuery = query.Encode()

	outlog.Printf("Request\n  %v\n", request.URL)

	response, err := client.Do(request)
	if err != nil {
		r <- updateResult{status: "Could not fetch update request", err: err}
		return
	}

	update, err := ioutil.ReadAll(response.Body)
	if err != nil {
		r <- updateResult{status: "Could not read update response", err: err}
		return
	}

	defer response.Body.Close()

	updateStatus := string(update)

	if updateStatus != fmt.Sprintf("good %v", newIP) {
		err = fmt.Errorf("Could not update dynamic DNS")
	}

	r <- updateResult{
		status: fmt.Sprintf("[%v] %v\n", response.Status, updateStatus),
		err:    err,
	}
}

// Update sends requests to update all Google Dynamic DNS hostnames with the current IP
func Update(
	client *http.Client,
	outlog, errlog *log.Logger,
	ipURL string,
	hostnames, usernames, passwords []string,
) {
	newIP, err := Get(ipURL)
	if err != nil {
		errlog.Printf("%v:  %v\n", "Could not get IP address for updating", err.Error())
		return
	}

	result := make(chan updateResult, len(hostnames))

	for eachIndex := range hostnames {
		go updateService(
			client,
			result,
			outlog,
			string(newIP),
			hostnames[eachIndex],
			usernames[eachIndex],
			passwords[eachIndex],
		)

		select {
		case r := <-result:
			if r.err != nil {
				errlog.Printf("\n  %v\n  %v\n", r.err.Error(), r.status)
			} else {
				outlog.Printf("%v\n", r.status)
			}
		}
	}
}
