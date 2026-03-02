package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct{}

func (c Client) sendGet() error {
	resp, err := http.Get("http://localhost:8080/echo")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Server response: " + string(body))
	return nil
}

func (c Client) sendPost() error {
	payload := "TestMessage"

	resp, err := http.Post("http://localhost:8080/echo", "", strings.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Server response: " + string(body))
	return nil
}

func main() {
	var c Client

	for {
		if err := c.sendPost(); err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
