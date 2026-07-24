package llm

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"github.com/sanskritisingh/pi-agent/internal/chat"
)
	

type Client struct {
	client  openai.Client
	model string
}

func NewClient(apiKey string, model string) *Client {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &Client{
		client: client,
		model: model,
	}
}

func convertMessages(messages []chat.Message) responses.ResponseInputParam {
	var input responses.ResponseInputParam

	for _, message := range messages {
		input = append(
			input,
			responses.ResponseInputItemParamOfMessage(
				message.Content,
				responses.EasyInputMessageRole(message.Role),
			),
		)
	}

	return input
}

func (c *Client) Generate(messages [] chat.Message) (string, error) {
	fmt.Println("Message received by llm")
	fmt.Println(messages)

	input := convertMessages(messages)

	resp, err := c.client.Responses.New(
		context.Background(),
		responses.ResponseNewParams{
			Model: openai.ChatModel(c.model),
			Input: responses.ResponseNewParamsInputUnion{
				OfInputItemList: input,
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.OutputText(), nil


	// return resp.OutputText(), nil
	//return "TODO", nil
}