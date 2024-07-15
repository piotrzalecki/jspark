
package main

import (
	"context"
	"encoding/json"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)


func callOllama(input Input) (completionObject, error) {
	var co completionObject
	llm, err := ollama.New(ollama.WithModel(input.Model))
	if err != nil {
		return co, err
	}

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, "<<" + input.Prompt + ">>")
	if err != nil {
		return co, err
	}

	err = json.Unmarshal([]byte(completion), &co)
	if err != nil {
		return co, err
	}
	return co, nil
}