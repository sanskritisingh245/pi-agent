package memory

type Memory struct {
	facts map[string]string
}

func New() *Memory {
	memory: memory.New(),
	return &Memory{
		facts: make(map[string]string),
	}
}

func (m *Memory) Remember(key, value string){
	m.facts[key] = value
}

func (m *Memory) Recall(key string) (string, bool){
	value, ok := m.facts[key]
	return value, ok
}

func (m *Memory) All() map[string]string {
	return m.facts
}