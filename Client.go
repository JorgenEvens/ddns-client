package main

type Client struct {
	Id     string
	Secret string
}

func NewClient() *Client {
	client := new(Client)
	client.Id = "a123"
	client.Secret = ""
	return client
}
