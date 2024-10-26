package resty

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

var url string = "https://api.openai.com/v1/chat/completions"

type Resty struct {
	client      *resty.Client
	Url         string
	RequestBody map[string]interface{}
	Headers     map[string]string
	Response    Response
}

type Response struct {
	Id      string `json:"id"`
	Choises []struct {
		Id      int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choises"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
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

func (r *Resty) SendRequest() (*Response, error) {
	client := r.client.R()
	client.SetHeaders(r.Headers)
	client.SetBody(r.RequestBody)
	resp, err := client.Post(r.Url)
	if err != nil {
		return &Response{}, err
	}
	log.Printf("Resty responce: %+v", resp)
	if resp.StatusCode() != 200 {
		err = fmt.Errorf("error: %s", resp.String())
		return &Response{}, err
	}
	err = json.Unmarshal(resp.Body(), &r.Response)
	if err != nil {
		return &Response{}, err
	}
	return &r.Response, nil
}
