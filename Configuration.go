package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Endpoint    string
	AccessToken string
}

func (api *Api) ReadConfiguration(path string) {
	var config Configuration

	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)

	api.Endpoint = config.Endpoint
	api.User = NewUser(config.AccessToken)
}

func (api *Api) SaveConfiguration(path string) error {
	var config Configuration

	config.Endpoint = api.Endpoint

	if api.User != nil {
		config.AccessToken = api.User.AccessToken
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&config)

	return err
}
