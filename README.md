# Go-Zendesk

Go-Zendesk is a [Zendesk API](https://developer.zendesk.com/rest_api/docs/core/introduction)
client library for Go.

This library is used internally at MEDIGO and the API resources are being implemented as needed.
**It's work in progress. Use with caution.**

## Development

### Linting

To lint the source code, use the command:

```
$ make lint
```

### Testing

The project contains integration tests that uses the Zendesk API. To execute them we must have access
to a it and configure the following environment variables:

| Name             | Description
| ---------------- | ----------------------------------
| ZENDESK_DOMAIN   | The Zendesk API domain.
| ZENDESK_USERNAME | The Zendesk API username.
| ZENDESK_PASSWORD | The Zendesk API password.


Then, to run the test, use the command:

```
$ make test
```

## License

MIT
