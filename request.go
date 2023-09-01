package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Copyright 2023 TheRipe(theripe666@gmail.com). All rights reserved.
// license that can be found in the LICENSE file.

type Request struct {
	client  *http.Client
	address *Address
}

// New Providing an init function to users who can use the "request.New()" build a request struct fast.
func New() *Request {
	return &Request{
		client: &http.Client{Timeout: 1 * time.Second},
	}
}

// SetTimeout Providing a method to set the request timeout duration
// timeout default duration is 1 second
func (r *Request) SetTimeout(duration time.Duration) {
	r.client.Timeout = duration
}

// SetAddress Providing a method to set the request address, including ip,port,protocol ...
func (r *Request) SetAddress(address *Address) {
	r.address = address
	if r.address.Protocol == "" {
		r.address.Protocol = HTTP
	}
}

// Post Providing generic post http requests
func (r *Request) Post(uri string, headers map[string]string, data map[string]interface{}) (response *Response, err error) {
	if r.address.Ip == "" || r.address.Port == 0 || uri == "" {
		return nil, errors.New("address params error")
	}
	url := r.address.Protocol + "://" + r.address.Ip + ":" + strconv.FormatInt(r.address.Port, 10) + uri
	bodyBtye, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("request body format error")
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBtye))
	if err != nil {
		return nil, err
	}
	if headers != nil {
		header := http.Header{}
		for key, value := range headers {
			header.Add(key, value)
		}
		req.Header = header
	}
	res, err := r.client.Do(req)
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

// Get Providing generic get http requests
func (r *Request) Get(uri string, headers map[string]string, params map[string]string) (response *Response, err error) {
	// Create get request
	if r.address.Ip == "" || r.address.Port == 0 || uri == "" {
		return nil, errors.New("address params error")
	}
	url := r.address.Protocol + "://" + r.address.Ip + ":" + strconv.FormatInt(r.address.Port, 10) + uri
	if len(params) != 0 {
		url += "?"
		for key, value := range params {
			url += key + ":" + value + "&"
		}
		url = url[:len(url)-1]
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		header := http.Header{}
		for key, value := range headers {
			header.Add(key, value)
		}
		req.Header = header
	}

	// Begin get Request
	res, err := r.client.Do(req)
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
