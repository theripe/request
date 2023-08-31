# request
'request' providing  convenient methods of net request

# example
get:

	request := New()
	address := &Address{
		Ip:   "localhost",
		Port: 8080,
	}
	request.SetAddress(address)
	response, err := request.Get("/test", EmptyHeader, EmptyParams)

post:

    request := New()
    address := &Address{
        Ip:   "localhost",
        Port: 8080,
    }
    request.SetAddress(address)
    response, err := request.Post("/test", EmptyHeader, map[string]interface{}{"a": 1})