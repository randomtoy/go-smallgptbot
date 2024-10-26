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
	requestBody := map[string]interface{}{
		"model": o.Model, // Или "gpt-3.5-turbo"
		"messages": []map[string]string{
			{"role": "system", "content": o.System},
			{"role": "user", "content": o.User},
		},
	}
	rest.RequestBody = requestBody
	log.Printf("sending request: %+v", rest)
	answer, err := rest.SendRequest()

	if err != nil {
		return "", err
	}
	log.Printf("getting answer: %+v", answer)
	return answer, nil
}
