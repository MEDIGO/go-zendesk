# Go-Zendesk

[![CircleCI](https://circleci.com/gh/MEDIGO/go-zendesk.svg?style=shield)](https://circleci.com/gh/MEDIGO/go-zendesk)
[![GoDoc](http://godoc.org/github.com/MEDIGO/go-zendesk/zendesk?status.png)](http://godoc.org/github.com/MEDIGO/go-zendesk/zendesk)

Go-Zendesk is a [Zendesk Core API](https://developer.zendesk.com/rest_api/docs/core/introduction) client library for Go.

This library is used internally at MEDIGO and the API resources are being implemented as needed.

**It's work in progress. Use with caution.**

## Usage

```go
package main

import (
  "log"

  "github.com/MEDIGO/go-zendesk/zendesk"
)

func main() {
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
```

## Development

### Linting

To lint the source code, use the command:

```
$ make lint
```

### Testing

The project contains integration tests that uses the Zendesk API. To execute them the following environment variables must be available:

```
ZENDESK_DOMAIN=<your-zendesk-domain>
ZENDESK_USERNAME=<your-zendesk-api-username>
ZENDESK_PASSWORD=<your-zendesk-api-password>
```

Then, to run the test, use the command:

```
$ make test
```

Please note that integration tests will create and alter entities in the configured Zendesk instance.
You most likely want to run them against a [Zendesk Sandbox](https://support.zendesk.com/hc/en-us/articles/203661826-Testing-changes-in-your-sandbox-Enterprise-) instance.

## Copyright and license

Copyright © 2017 MEDIGO GmbH. go-zendesk is licensed under the MIT License. See LICENSE for the full license text.
