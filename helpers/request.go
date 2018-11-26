package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"misteraladin.com/jasmine/payment-gateway/requests"
)

var defHTTPTimeout = 80 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

func IsSuccess(status int) bool {
	if status >= 400 && status <= 505 {
		return false
	}

	return true
}

func CallRequest(method, path string, headers []requests.Header, body io.Reader, v interface{}) error {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		fmt.Println("Request creation failed : ", err)
		return err
	}

	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	start := time.Now()

	fmt.Println("Request ", method, ": ", path)
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Cannot send request : ", err)
		return err
	}
	fmt.Println("Completed in ", time.Since(start))
	defer res.Body.Close()

	if !IsSuccess(res.StatusCode) {
		return errors.New("Request not success with status code " + strconv.Itoa(res.StatusCode))
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Cannot read response body: ", err)
		return err
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}
