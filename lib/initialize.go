package lib

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Initialize required inputs from the service runtime environment
func Initialize() (
	hostnames, usernames, passwords []string,
	ipURL string,
	frequencyTime int,
	client *http.Client,
	err error,
) {
	client = &http.Client{}

	var hosts, users, passes string
	found := false

	if hosts, found = os.LookupEnv("GDDNS_HOSTNAME"); !found {
		err = fmt.Errorf("No hostname given for updating")
		return
	}

	if users, found = os.LookupEnv("GDDNS_USERNAME"); !found {
		err = fmt.Errorf("No username given for updating")
		return
	}

	if passes, found = os.LookupEnv("GDDNS_PASSWORD"); !found {
		err = fmt.Errorf("No password given for updating")
		return
	}

	hostnames = strings.Split(hosts, ",")
	usernames = strings.Split(users, ",")
	passwords = strings.Split(passes, ",")

	if ipURL, found = os.LookupEnv("GDDNS_IP_URL"); !found {
		err = fmt.Errorf("No IP URL given for updating")
		return
	}

	var frequency string

	if frequency, found = os.LookupEnv("GDDNS_FREQUENCY"); !found {
		err = fmt.Errorf("No update frequency given for updating")
		return
	}

	if frequencyTime, err = strconv.Atoi(frequency); err != nil {
		return
	}

	return
}
