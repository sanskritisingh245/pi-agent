package agent 

import (
	"github.com/sanskritisingh/pi-agent/internal/chat"
	"github.com/sanskritisingh/pi-agent/internal/llm"
)

type Agent struct {
	llm     *llm.Client
	history []chat.Message
}


func New(llm *llm.Client) *Agent {
	return &Agent{
		llm: llm,
		history: []chat.Message{},
	}
}

func (a *Agent) Respond(message string) (string, error) {
	a.history = append(a.history , chat.Message{
		Role: "user",
		Content: message,
	})

	reply, err := a.llm.Generate(a.history)
	if err != nil{
		return "",err
	}

	a.history = append(a.history, chat.Message {
		Role: "assistant",
		Content: reply,
	})

	return reply, nil 
}