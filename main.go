package main

import (
	"ProductClient/src/client"
	"context"
)

func main() {
	client.Start(context.Background())
}
