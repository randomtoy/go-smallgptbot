package resty

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

var url string = "https://api.openai.com/v1/chat/completions"

type Resty struct {
	client      *resty.Client
	Url         string
	RequestBody map[string]string `json:"requestbody"`
	Headers     map[string]string
}

func New(token string) *Resty {

	r := resty.New()

	return &Resty{
		Url:    url,
		client: r,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + token,
		},
	}
}

func (r *Resty) SendRequest() (string, error) {
	client := r.client.R()
	client.SetHeaders(r.Headers)
	client.SetBody(r.RequestBody)
	resp, err := client.Post(r.Url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		err = fmt.Errorf("error: %s", resp.String())
		return "", err
	}

	return resp.String(), nil
}
