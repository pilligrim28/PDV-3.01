package dr600

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)



type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) SetFrequency(freq string) error {
	body := map[string]string{"frequency": freq}
	jsonBody, _ := json.Marshal(body)
	
	req, _ := http.NewRequest("POST", c.BaseURL+"/api/frequency", 
		bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	return nil
}