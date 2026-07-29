package agent 

import (
	"github.com/sanskritisingh/pi-agent/internal/chat"
	"github.com/sanskritisingh/pi-agent/internal/llm"
	"github.com/sanskritisingh/pi-agent/internal/memory"
)

type Agent struct {
	llm            *llm.Client
	history        []chat.Message
	memory         *memory.Memory
	systemPrompt   string
}


func New(llm *llm.Client) (*Agent, error) {
	mem, err := memory.Load()
	if err != nil {
		return nil, err
	}

	return &Agent{
		llm:     llm,
		history: []chat.Message{},
		memory:  mem,
		systemPrompt: `You are Pi, a warm, thoughtful, and emotionally intelligent AI assistant.

Keep your responses natural and conversational.

Be concise unless the user asks for more detail.`,
	}, nil
}


func (a *Agent) buildSystemPrompt() string {
	prompt := a.systemPrompt

	for key, value := range a.memory.All() {
		prompt += "\n" + key + ": " + value
	}

	return prompt

}

func (a *Agent) Respond(message string) (string, error) {
	facts := memory.Extract(message)

	for key, value := range facts {
		a.memory.Remember(key, value)
	}
	
	if len(facts) > 0 {
		if err := memory.Save(a.memory); err != nil {
			return "", err
		}
	}

	a.history = append(a.history, chat.Message{
		Role:    chat.RoleUser,
		Content: message,
	})

	messages := []chat.Message{
		{
			Role: chat.RoleSystem,
			Content:  a.buildSystemPrompt(),
		},
	}

	messages = append(messages, a.history...)
	
	reply, err := a.llm.Generate(messages)
	if err != nil{
		return "",err
	}

	a.history = append(a.history, chat.Message {
		Role: chat.RoleAssistant,
		Content: reply,
	})

	return reply, nil 
}