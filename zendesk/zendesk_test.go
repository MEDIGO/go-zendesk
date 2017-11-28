package zendesk_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MEDIGO/go-zendesk/zendesk"
	"github.com/stretchr/testify/require"
)

func TestClientHeaders(t *testing.T) {
	ch := make(chan string, 1)

	handler := func(w http.ResponseWriter, r *http.Request) {
		ch <- r.Header.Get("foo")
		http.NotFound(w, r)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	require.NotNil(t, server)
	defer server.Close()

	client1, err := zendesk.NewURLClient(server.URL, "", "")
	require.NoError(t, err)

	client2 := client1.WithHeader("foo", "bar")
	require.NotNil(t, client2)

	client1.ShowLocale(0)
	require.Equal(t, "", <-ch, "expected no header")

	client2.ShowLocale(0)
	require.Equal(t, "bar", <-ch, "expected header")
}

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
