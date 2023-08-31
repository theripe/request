package request

// Copyright 2023 TheRipe(theripe666@gmail.com). All rights reserved.
// license that can be found in the LICENSE file.

const (
	HTTP  = "http"
	HTTPS = "https"
)

var (
	EmptyHeader = map[string]string{}
	EmptyParams = map[string]string{}
	EmptyData   = map[string]interface{}{}
)
