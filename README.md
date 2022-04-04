Moneybird Go
=============
[![GoDoc](https://godoc.org/github.com/dannyvankooten/moneybird-go?status.svg)](https://godoc.org/github.com/dannyvankooten/moneybird-go)
 [![Build Status](https://github.com/dannyvankooten/moneybird-go/actions/workflows/go.yml/badge.svg)](https://github.com/dannyvankooten/moneybird-go/actions/workflows/go.yml)
 [![Coverage Status](https://coveralls.io/repos/github/dannyvankooten/moneybird-go/badge.svg?branch=master)](https://coveralls.io/github/dannyvankooten/moneybird-go?branch=master)

An unofficial Go client library for [Moneybird](https://developer.moneybird.com/). This package is still experimental and could be subject to heavy change.

## Usage

```go
import (
  "net/http"
  "github.com/dannyvankooten/moneybird-go"
)

mb := &moneybird.Client{
  Token: "token",
  AdministrationID: "administration-id-here",
  HTTPClient: &http.Client{},
}

contact, _ := mb.Contact().Create(&moneybird.Contact{
  Email: "john@doe.com",
  FirstName: "John",
  LastName: "Doe",
})
```

See the integration tests for some more working examples.

## Testing

In order to run the integration tests, you have to configure a sandbox account in Moneybird. Before running the integration tests with `go test`, make sure the following environment variables are set.

```
export MONEYBIRD_TEST_TOKEN="your-sandbox-token"
export MONEYBIRD_TEST_ADMINISTRATION_ID="your-sandbox-administration-id"
```


## License

MIT Licensed. See the [LICENSE](LICENSE) file for details.
