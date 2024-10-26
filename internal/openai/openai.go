package openai

import (
	"log"

	"github.com/randomtoy/go-smallgptbot/internal/resty"
)

type OpenAi struct {
	Token  string
	System string
	User   string
	Model  string
}

func New(token string) *OpenAi {
	return &OpenAi{
		Token: token,
	}
}

func (o *OpenAi) Send() (string, error) {
	rest := resty.New(o.Token)
	rest.RequestBody = map[string]string{
		"model":    o.Model,
		"messages": "{\"role\": \"system\",\"content\":\"" + o.System + "\"},{\"role\":\"user\",\"content\": \"" + o.User + "\"}",
	}
	log.Printf("sending request: %+v", rest)
	answer, err := rest.SendRequest()

	if err != nil {
		return "", err
	}
	log.Printf("getting answer: %+v", answer)
	return answer, nil
}
