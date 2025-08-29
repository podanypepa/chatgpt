// Package main sends a message to a ChatGPT.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/chatgpt"
)

func main() {
	client, err := chatgpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &chatgpt.ChatRequest{
		Model: chatgpt.Gpt4OMini,
		Messages: []chatgpt.ChatMessage{
			{Role: "user", Content: "Hello, world!"},
		},
	}

	res, err := client.Send(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println("response:", res.Choices[0].Message.Content)
	fmt.Println("total tokens used:", res.Usage.TotalTokens)
	fmt.Println("input tokens used:", res.Usage.PromptTokens)
	fmt.Println("output tokens used:", res.Usage.CompletionTokens)
}
