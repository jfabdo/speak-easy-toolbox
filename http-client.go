package toolbox

import (
	"io/ioutil"
	"net/http"
	"time"
)

//HTTPClient will return an http
func HTTPClient(timeout time.Duration) http.Client {
	client := http.Client{
		Timeout: timeout * time.Second,
	}

	return client
}

func getURL(url string) string {
	response, err := http.Get("")
	if err != nil {
		//do something
	}
	retVal, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//do something
	}
	return string(retVal)
}
