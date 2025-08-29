# chatgpt simplesend example

## simple send example
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

## dependencies
`OPENAI_API_KEY` environment variable is required.

## usage
```bash
git clone https://github.com/podanypepa/chatgpt.git
cd examples/simplesend

# if you have OPENAI_API_KEY in your environment
go run .

# or you can set it inline
OPENAI_API_KEY=your_api_key go run .
```
