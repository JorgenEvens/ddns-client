package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Api struct {
	Endpoint string
	User     *User
	Client   http.Client
}

func (api *Api) Request(method string, path string, data url.Values, v interface{}) error {
	data_buf := bytes.NewBufferString(data.Encode())

	url := api.Endpoint + path

	if api.User != nil && api.User.AccessToken != "" {
		url = url + "?accessToken=" + api.User.AccessToken
	}
	req, err := http.NewRequest(method, url, data_buf)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := api.Client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(v)

	if err != nil {
		return err
	}

	return nil
}

func (api *Api) Post(path string, data url.Values, v interface{}) error {
	return api.Request("POST", path, data, v)
}
