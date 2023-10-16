package udm_api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func init() {
	log.Println("UDM_API init()")
}

// Client is a struct that holds internal state for the API client
type Client struct {
	username   string
	password   string
	Host       string
	cookie     string
	csrfToken  string
	httpClient *http.Client
}

// loginRequestData is a struct that holds the data for a login request
type loginRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateClient creates a new API client
func CreateClient(u, p, h string, timeout int, skip bool) (*Client, error) {
	log.Println("UDM_API CreateClient()")

	// TODO: test if all vars set are valid

	// default timeout
	if timeout == 0 {
		timeout = 10
	}

	// create client
	c := &Client{
		username: u,
		password: p,
		Host:     h,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: skip},
			},
			Timeout: time.Duration(timeout) * time.Second, // TODO: is this right?
		},
	}

	// initial login
	err := c.login()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Api hands requests to the API, just moves bytes from here to there
func (c *Client) Api(method, url string, data []byte) ([]byte, error) {
	// build request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", c.cookie)
	req.Header.Set("X-CSRF-Token", c.csrfToken)

	// send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	// read response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// update auth tokens
	if c.cookie != resp.Header.Get("Set-Cookie") || c.csrfToken != resp.Header.Get("x-csrf-token") {
		c.cookie = resp.Header.Get("Set-Cookie")
		c.csrfToken = resp.Header.Get("x-csrf-token")
	}

	// check if response is 2xx
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprint("API returned non-200 status code: ", resp.StatusCode, " "+string(respBytes)))
	}

	return respBytes, nil
}

// login to the UDM and stores the auth info
func (c *Client) login() error {
	url := fmt.Sprintf("%s/api/auth/login", c.Host)

	// login to bytes
	loginBytes, err := json.Marshal(loginRequestData{
		Username: c.username,
		Password: c.password,
	})
	if err != nil {
		return err
	}

	// create request
	_, err = c.Api("POST", url, loginBytes)
	if err != nil {
		return err
	}

	// background routine for refresh
	go func() {
		for {
			// wait 5 minutes
			rerr := c.refreshToken()
			if rerr != nil {
				log.Fatalln("Failed to refresh token:", rerr)
			}
			time.Sleep(5 * time.Minute)
		}
	}()

	return nil
}

// refreshToken refreshes the authentication token
func (c *Client) refreshToken() error {
	url := fmt.Sprintf("%s/api/users/self", c.Host)

	// create request
	_, err := c.Api("GET", url, nil)
	if err != nil {
		return err
	}

	return nil
}
