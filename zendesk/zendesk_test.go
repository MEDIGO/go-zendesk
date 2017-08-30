package zendesk_test

import (
  "log"

  "github.com/MEDIGO/go-zendesk/zendesk"
)

func Example() {
    client, err := zendesk.NewClient("domain", "username", "password")
    if err != nil {
        log.Fatal(err)
    }
    ticket, err := client.ShowTicket(1)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Requester ID is: %d", *ticket.RequesterID)
}
