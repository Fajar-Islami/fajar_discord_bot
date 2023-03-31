package service

import (
	"io"
	"log"
	"net/http"
)

// Client hold http object
type Client struct {
	req *http.Request
	err error
}

func Get(uri string) *Client {
	r := new(Client)

	return r.Request("GET", uri, nil)
}

func Post(uri string, body io.Reader) *Client {
	r := new(Client)

	return r.Request("POST", uri, body)
}

// Header add header
func (c *Client) Header(key, value string) *Client {
	c.req.Header.Add(key, value)
	return c
}

// Headers set multiple header at once
func (c *Client) Headers(headers http.Header) *Client {
	c.req.Header = headers
	return c
}

// Request wrapper of http NewRequest
func (c *Client) Request(method string, url string, body io.Reader) *Client {
	c.req, c.err = http.NewRequest(method, url, body)
	return c
}

// Do execute request
func (c *Client) Do() (response *http.Response, err error) {

	httpClient := http.DefaultClient

	log.Printf("Address : %s \nMethod : %s \nURL : %s \nRequest Header : %v \nRequest Body : %s \n", c.req.RemoteAddr, c.req.Method, c.req.URL, c.req.Header, c.req.Body)

	// var mapInterface map[string]any
	// err = json.Unmarshal( c.req.Body, mapInterface)
	// if err != nil {
	// 	log.Println(err)
	// 	return "Error Marshal SearchBot AI"
	// }

	response, err = httpClient.Do(c.req)

	return response, err
}
