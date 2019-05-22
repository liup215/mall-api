package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	host string
	c    *http.Client
}

func NewClient(host string) *Client {
	return &Client{
		host: host,
		c:    &http.Client{},
	}
}

func (client *Client) RestfulGET(uri string, params url.Values) error {
	return nil

}

func (client *Client) RestfulPOST(uri string, params map[string]interface{}) ([]byte, error) {
	req, err := client.request(http.MethodPost, uri, params)
	if err != nil {
		return nil, err
	}

	resp, err := client.c.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return client.parseResponse(resp)

}

func (client *Client) request(method, uri string, params map[string]interface{}) (*http.Request, error) {
	uri = client.host + uri

	var r *bytes.Reader
	switch method {
	case http.MethodGet:

		values := url.Values{}
		for k, v := range params {
			values.Set(k, v.(string))
		}
		uri = uri + "?" + values.Encode()
	default:
		method = http.MethodPost
		data, err := json.Marshal(&params)
		if err != nil {
			return nil, err
		}

		fmt.Println(string(data))

		r = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, uri, r)
	if err != nil {
		err = errors.New("请求失败: " + err.Error())
		return nil, err
	}
	const (
		_contentType = "Content-Type"
		_urlencoded  = "application/json"
	)

	if method == http.MethodPost {
		req.Header.Set(_contentType, _urlencoded)
	}
	return req, nil
}

func (client *Client) parseResponse(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("body 解析错误" + err.Error())
		return nil, err
	}
	return body, err

}
