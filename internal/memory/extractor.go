package memory

import "strings"

func Extract(message string) map[string]string {
	facts := make(map[string]string)

	lower := strings.ToLower(message)

	if strings.HasPrefix(lower, "my name is "){
		name := strings.TrimSpace(message[len("My name is "):])
		facts["name"] = name
	}

	return facts
}