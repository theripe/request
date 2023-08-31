package request

import (
	"fmt"
	"testing"
)

// Copyright 2023 TheRipe(theripe666@gmail.com). All rights reserved.
// license that can be found in the LICENSE file.

func TestAddressStruct(t *testing.T) {
	address := &Address{
		Ip:   "192.168.0.1",
		Port: 8901,
	}
	fmt.Println(address)
}
