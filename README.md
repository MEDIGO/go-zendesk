# Go-Zendesk

[![CircleCI](https://circleci.com/gh/MEDIGO/go-zendesk.svg?style=shield)](https://circleci.com/gh/MEDIGO/go-zendesk)

Go-Zendesk is a [Zendesk Core API](https://developer.zendesk.com/rest_api/docs/core/introduction) client library for Go.

This library is used internally at MEDIGO and the API resources are being implemented as needed.
**It's work in progress. Use with caution.**

## Usage

```go
package main

import "github.com/MEDIGO/go-zendesk/zendesk"

func main() {
  zendeskcl, err := zendesk.NewClient("your-zendesk-domain", "your-username", "your-api-password")

  if err != nil {
    // I can now use zendesk client...
  }
}
```

## Development

### Linting

To lint the source code, use the command:

```
$ make lint
```

### Testing

The project contains integration tests that uses the Zendesk API. To execute them we must have access to a it and configure the following environment variables:

| Name             | Description
| ---------------- | ----------------------------------
| ZENDESK_DOMAIN   | The Zendesk API domain.
| ZENDESK_USERNAME | The Zendesk API username.
| ZENDESK_PASSWORD | The Zendesk API password.


Then, to run the test, use the command:

```
$ make test
```

## Copyright and license

Copyright © 2016 MEDIGO GmbH. go-zendesk is licensed under the MIT License. See LICENSE for the full license text.
