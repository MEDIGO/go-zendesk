// Package zendesk provides a client for using the Zendesk Core API.
//
// Usage:
//
//    package main
//
//    import (
//      "log"
//
//      "github.com/MEDIGO/go-zendesk/zendesk"
//    )
//
//    func main() {
//        client, err := zendesk.NewClient("domain", "username", "password")
//        if err != nil {
//            log.Fatal(err)
//        }
//        ticket, err := client.ShowTicket(1)
//        if err != nil {
//            log.Fatal(err)
//        }
//        log.Printf("Requester ID is: %d", *ticket.RequesterID)
//    }
package zendesk
