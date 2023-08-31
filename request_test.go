package request

import (
	"fmt"
	"testing"
)

// Copyright 2023 TheRipe(theripe666@gmail.com). All rights reserved.
// license that can be found in the LICENSE file.

func TestRequestHttpGet(t *testing.T) {
	request := New()
	address := &Address{
		Ip:   "localhost",
		Port: 8080,
	}
	request.SetAddress(address)
	response, err := request.Get("/test", EmptyHeader, EmptyParams)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(response)
}

func TestRequestHTTPPost(t *testing.T) {
	request := New()
	address := &Address{
		Ip:   "localhost",
		Port: 8080,
	}
	request.SetAddress(address)
	response, err := request.Post("/test", EmptyHeader, map[string]interface{}{"a": 1})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(response)
}
