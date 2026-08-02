package llm

import (
	"context"
	"fmt"
	"encoding/json"
	"strings"

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

func (c *Client) ExtractFacts(message string) (map[string]string, error) {
	systemPrompt := `You are an information extraction system.

Your job is to extract ONLY long-term facts about the user.

Return ONLY a valid JSON object.
Do NOT return markdown.
Do NOT wrap the JSON in triple backticks.
Do NOT include explanations or extra text.

Rules:
- Return a flat JSON object.
- Keys must be lowercase snake_case.
- Values must always be strings.
- If no long-term facts are present, return {}.

Use ONLY these keys when applicable:
- name
- age
- city
- occupation
- education
- learning
- project
- likes
- favorite_language
- favorite_database
- hobby
- goal

Remember information that is likely to be useful in future conversations.

Do NOT remember:
- Greetings
- Temporary emotions
- One-time requests
- Questions
- Weather
- Short-lived information

Examples:

User: My name is Sans.
Output:
{"name":"Sans"}

User: I'm studying Computer Science.
Output:
{"education":"Computer Science"}

User: I'm learning Go.
Output:
{"learning":"Go"}

User: I love coffee.
Output:
{"likes":"coffee"}

User: My favourite database is PostgreSQL.
Output:
{"favorite_database":"PostgreSQL"}

User: I live in Bangalore.
Output:
{"city":"Bangalore"}

User: What's the weather today?
Output:
{}`

	input := responses.ResponseInputParam{
		responses.ResponseInputItemParamOfMessage(
			systemPrompt,
			responses.EasyInputMessageRole("system"),
		),
		responses.ResponseInputItemParamOfMessage(
			message,
			responses.EasyInputMessageRole("user"),
		),
	}

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
		return nil, err
	}

	output := strings.TrimSpace(resp.OutputText())

	output = strings.TrimPrefix(output, "```json")
	output = strings.TrimPrefix(output, "```")
	output = strings.TrimSuffix(output, "```")
	output = strings.TrimSpace(output)

	facts := make(map[string]string)

	err = json.Unmarshal([]byte(output), &facts)
	if err != nil {
		return nil, err
	}

	return facts, nil
}