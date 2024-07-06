package model

type RequestCalculation struct {
	Model           LLM
	EinsteinInput   int
	EinsteintOutput int
}

type LLM struct {
	Name       string
	Multiplier int
}

func NewLLM(name string) *LLM {

	var llm LLM
	llm.Name = name

	switch name {
	case "own":
		llm.Multiplier = 7
	case "foundation":
		llm.Multiplier = 10
	}

	return &llm
}
