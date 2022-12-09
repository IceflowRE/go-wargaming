# Wargaming API Client
![maintained](https://img.shields.io/badge/maintained-yes-brightgreen.svg)
![Programming Language](https://img.shields.io/badge/language-Go-orange.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/IceflowRE/go-wargaming/blob/master/LICENSE.md)

[![Go report card](https://goreportcard.com/badge/github.com/IceflowRE/go-wargaming/v2)](https://goreportcard.com/report/github.com/IceflowRE/go-wargaming/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/IceflowRE/go-wargaming/v2.svg)](https://pkg.go.dev/github.com/IceflowRE/go-wargaming/v2)

---

Go client for accessing [Wargaming.net Public API](https://developers.wargaming.net/documentation/guide/getting-started/).

This API client is generated automatically based on the documentation provided by Wargaming. All endpoints are available.  
The documentation is often inaccurate, especially the return types (object, list, map), so fixes afterwards are required. If something was not corrected, please open an issue.

## Installation

go-wargaming is compatible with modern Go releases and modules enabled.

```shell
go get github.com/IceflowRE/go-wargaming/v3
```

will resolve and add the package to the current development module, along with its dependencies.

Or just import it and run `go get` afterwards

```go
import "github.com/IceflowRE/go-wargaming/v3/wargaming"
```

## Usage

```go
import "github.com/IceflowRE/go-wargaming/v3/wargaming" // with go modules
import "github.com/IceflowRE/go-wargaming/wargaming"    // without go modules
```

Create a new Wargaming Client and use the different services to access the different sections of the Wargaming API.
The services of a client are the same as the Wargaming API sections in their documentation.
```go
client := wargaming.NewClient("a7f838650dcb008552966db063eeeb35", nil)
// get World of Tanks EU accounts starting with "Yzne"
res, err := client.Wot.AccountList(context.Background(), wargaming.RealmEu, "Yzne", nil)
if err == nil {
    // print them to console
	for _, value := range res {
		fmt.Println(value)
	}
}
```

Also, the http client can be customized.
```go
client := wargaming.NewClient("a7f838650dcb008552966db063eeeb35", &wargaming.ClientOptions{
	HTTPClient: &http.Client{
		Timeout: 10 * time.Second,
	},
})
```

Some API methods have options which can be passed.
```go
client := wargaming.NewClient("a7f838650dcb008552966db063eeeb35", nil)
// get World of Tanks EU account named exact "Lichtgeschwindigkeit" and return only the 'account_id' field
res, err = client.Wot.AccountList(context.Background(), wargaming.RealmEu, "Lichtgeschwindigkeit", &wot.AccountListOptions{
	Fields: []string{"account_id"},
	Type: wargaming.String("exact"),
})
```

All structs for Wargaming resources use pointer values for all non-repeated fields. This allows distinguishing between unset fields and those set to a zero-value. Helper functions have been provided to easily create these pointers for string, bool, and int values. For example:
```go
options := &wot.AccountListOptions{
	Type: wargaming.String("exact"),
}
```

To retrieve the error returned by the Wargaming API, use `errors.As`.

```go
_, err := client.Wot.AccountList(context.Background(), wargaming.RealmEu, "Yzne", nil)
if err != nil {
	// handle a Wargaming Response Error, returned by the API itself
	var respErr *wargaming.ResponseError
	if errors.As(err, &respErr) {
		fmt.Println(respErr.Error())
	} else {
		// handle client or other errors
		fmt.Println(err.Error())
	}
}
```

NOTE: Using the context package, one can easily pass cancellation signals and deadlines to various services of the client for handling a request. In case there is no context available, then `context.Background()` can be used as a starting point.

## Development

The client library is generated automatically. The generator is located in [tools/generator/main.go](tools/generator/main.go).

Endpoint patches can be found in [patches.go](tools/generator/internal/patches.go).

To update the library just run the generator from project root.
```shell
go build -o generator github.com/IceflowRE/go-wargaming/v3/tools/generator
./generator
```

### Testing

Client tests require an environment variable `WARGAMING_API_ID` with the API ID as value.

Not all endpoints are tested either they are complicated to test or method who required a clan membership, etc.
Those tests are explicitly skipped.

Also, the tests are not covering edge cases or all possibilities, only the availability of the endpoint is tested and if the default return data can be parsed.

## Contributing

Every contribution and talk about the structure and organization of the project are always welcome.

## Thanks

This client library is heavily influenced by

- [go-github](https://github.com/google/go-github)
- [go-gitlab](https://github.com/xanzy/go-gitlab)
- [cloudflare-go](https://github.com/cloudflare/cloudflare-go)

## MIT License

Copyright 2018-present Iceflower S (iceflower@gmx.de)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.