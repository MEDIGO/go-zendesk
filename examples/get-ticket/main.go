package main

import (
	"encoding/json"
	"fmt"

	"github.com/medigo/go-zendesk/zendesk"
)

func main() {
	client, err := zendesk.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ticket, err := client.Tickets.Get(123)
	if err != nil {
		panic(err)
	}

	raw, err := json.MarshalIndent(ticket, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(raw))
}
