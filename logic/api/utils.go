package api

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"xframe/log"
)

func SendHttpPostRequest(url string, header http.Header, body io.Reader, timeOut uint32) (res []byte, err error, status_code int) {
	client := &http.Client{Timeout: time.Duration(timeOut) * time.Second}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	req.Header = header
	result, err := client.Do(req)
	if err != nil {
		return
	}
	log.DEBUG(result)
	status_code = result.StatusCode
	defer result.Body.Close()
	res, err = ioutil.ReadAll(result.Body)
	return
}
