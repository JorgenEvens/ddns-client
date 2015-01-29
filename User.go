package main

import (
	"errors"
	"net/url"
)

type User struct {
	AccessToken string
	Username    string
	Password    string
	Client      *Client
}

func NewUser(AccessToken string) *User {
	user := new(User)
	user.AccessToken = AccessToken
	return user
}

func (user *User) Login(api *Api) error {

	if user.Client == nil {
		return errors.New("Client not set.")
	}

	if user.Username == "" || user.Password == "" {
		return errors.New("Username or Password not set.")
	}

	data := url.Values{}
	data.Set("username", user.Username)
	data.Set("password", user.Password)
	data.Set("grant_type", "password")
	data.Set("client_id", user.Client.Id)
	data.Set("client_secret", user.Client.Secret)

	var resp struct {
		// When error
		Code              int
		Error             string
		Error_description string

		// When success
		Access_token string
		Token_type   string
	}

	err := api.Request("POST", "/oauth/token", data, &resp)

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return errors.New(resp.Error_description)
	}

	api.User = user
	user.AccessToken = resp.Access_token

	return nil
}
