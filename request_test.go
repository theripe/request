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
		uri:  "/test",
		Ip:   "localhost",
		Port: 8080,
	}
	response, err := request.HttpGet(address, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(response)
}
