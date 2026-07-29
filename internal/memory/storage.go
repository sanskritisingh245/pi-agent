package memory

import (
	"encoding/json"
	"os"
)

func Save(m *Memory) error {
	data, err := json.MarshalIndent(m.All(), "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("memory.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Load () (*Memory, error){
	data, err := os.ReadFile("memory.json")

	if os.IsNotExist(err){
		return New(), nil
	}

	if err != nil {
		return nil, err
	}

	facts := make(map[string]string)

	err = json.Unmarshal(data, &facts)
	if err != nil {
		return nil, err
	}

	return &Memory{
		facts: facts,
	}, nil
}