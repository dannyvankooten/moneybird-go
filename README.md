Moneybird Go
=============
[![GoDoc](https://godoc.org/github.com/dannyvankooten/moneybird-go?status.svg)](https://godoc.org/github.com/dannyvankooten/moneybird-go)
 [![Build Status](https://travis-ci.org/dannyvankooten/moneybird-go.png?branch=master)](https://travis-ci.org/dannyvankooten/moneybird-go)

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
  HTTPClient: &http.Client{}
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

The MIT License (MIT)

Copyright (c) 2017 Danny van Kooten

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
