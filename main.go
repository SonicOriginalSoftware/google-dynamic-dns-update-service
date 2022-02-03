package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func initialize() (
	hostname string,
	ipURL string,
	frequencyTime int,
	request *http.Request,
	err error,
) {
	found := false

	if hostname, found = os.LookupEnv("GDDNS_HOSTNAME"); !found {
		err = fmt.Errorf("No hostname given for updating")
		return
	}

	if ipURL, found = os.LookupEnv("GDDNS_IP_URL"); !found {
		err = fmt.Errorf("No IP URL given for updating")
		return
	}

	var username, password, frequency string

	if username, found = os.LookupEnv("GDDNS_USERNAME"); !found {
		err = fmt.Errorf("No username given for updating")
		return
	}

	if password, found = os.LookupEnv("GDDNS_PASSWORD"); !found {
		err = fmt.Errorf("No password given for updating")
		return
	}

	if frequency, found = os.LookupEnv("GDDNS_FREQUENCY"); !found {
		err = fmt.Errorf("No update frequency given for updating")
		return
	}

	if frequencyTime, err = strconv.Atoi(frequency); err != nil {
		return
	}

	if request, err = http.NewRequest("GET", "https://domains.google.com/nic/update", nil); err != nil {
		return
	}

	request.SetBasicAuth(username, password)

	return
}

func registerInterruptHandler(frequencyTime int, ticker *time.Ticker) chan os.Signal {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	return c
}

func getIP(url string) ([]byte, error) {
	// response, err := http.Get("https://domains.google.com/checkip")
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func updateService(request *http.Request, hostname string, ipURL string) {
	newIP, err := getIP(ipURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IP address for updating:\n%v", err)
		return
	}
	ip := string(newIP)

	query := request.URL.Query()
	query.Set("hostname", hostname)
	query.Set("myip", ip)

	request.URL.RawQuery = query.Encode()

	fmt.Fprintf(os.Stdout, "\nRequest: %v\n", request.URL)

	client := http.Client{}

	response, err := client.Do(request)

	update, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read update response:\n%v", err)
		return
	}

	defer response.Body.Close()

	updateStatus := string(update)

	if updateStatus != fmt.Sprintf("good %v", ip) {
		fmt.Fprintf(os.Stderr, "\n**Could not update dynamic DNS**\n")
	}

	fmt.Fprintf(
		os.Stdout,
		"  Status: %v\n  Status Code: %v\n  Update Response: %v\n",
		response.Status,
		response.StatusCode,
		updateStatus,
	)
}

func main() {
	hostname, ipURL, frequencyTime, request, err := initialize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	ticker := time.NewTicker(time.Duration(frequencyTime) * time.Second)

	interrupt := registerInterruptHandler(frequencyTime, ticker)
	interrupted := false

	defer close(interrupt)
	defer ticker.Stop()

	for !interrupted {
		select {
		case <-ticker.C:
			updateService(request, hostname, ipURL)
		case <-interrupt:
			fmt.Fprintf(os.Stdout, "\n - Service stop requested!\n")
			interrupted = true
		}
	}

	fmt.Fprintf(os.Stdout, "Service stopped!\n")
}
