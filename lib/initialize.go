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
	splitHostnames []string,
	ipURL string,
	frequencyTime int,
	username string,
	password string,
	client *http.Client,
	err error,
) {
	client = &http.Client{}

	hostnames := ""
	found := false

	if hostnames, found = os.LookupEnv("GDDNS_HOSTNAME"); !found {
		err = fmt.Errorf("No hostname given for updating")
		return
	}

	splitHostnames = strings.Split(hostnames, ",")

	if ipURL, found = os.LookupEnv("GDDNS_IP_URL"); !found {
		err = fmt.Errorf("No IP URL given for updating")
		return
	}

	var frequency string

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

	return
}
