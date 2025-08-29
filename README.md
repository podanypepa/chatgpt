# chatgpt
go chatgpt package

```go
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

	res, err := client.SimpleSend(context.TODO(), "Hello, how are you?")
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Choices[0].Message.Content)
}

```
