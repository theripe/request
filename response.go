package request

// Copyright 2023 TheRipe(theripe666@gmail.com). All rights reserved.
// license that can be found in the LICENSE file.

type Response struct {
	StatusCode      int
	ResponseJsonMap map[string]interface{}
}
