package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"request/data"
	"strconv"
	"time"
)

// Copyright 2023 TheRipe(theripe666@gmail.com). All rights reserved.
// license that can be found in the LICENSE file.

type Request struct {
	client *http.Client
}

// New Providing an init function to users who can use the "request.New()" build a request struct fast.
func New() *Request {
	return &Request{
		client: &http.Client{Timeout: 1 * time.Second},
	}
}

// HttpPost Providing generic post http requests
func (r *Request) HttpPost() {

}

// HttpGet Providing generic get http requests
func (r *Request) HttpGet(address *Address, params map[string]string, headers map[string]interface{}) (response *Response, err error) {
	// Create get request
	if address.Ip == "" || address.Port == 0 || address.uri == "" {
		return nil, errors.New("address params error")
	}
	url := data.HTTP + "://" + address.Ip + ":" + strconv.FormatInt(address.Port, 10) + address.uri
	if len(params) != 0 {
		url += "?"
		for key, value := range params {
			url += key + ":" + value
		}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Begin get Request
	res, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("read response message error")
	}
	defer res.Body.Close()

	response = &Response{
		StatusCode: res.StatusCode,
	}
	responseJsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &responseJsonMap)
	if err != nil {
		return response, nil
	}
	response.ResponseJsonMap = responseJsonMap
	return response, nil
}
