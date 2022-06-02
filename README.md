# Wargaming API Client
![maintained](https://img.shields.io/badge/maintained-yes-brightgreen.svg)
![Programming Language](https://img.shields.io/badge/language-Go-orange.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/IceflowRE/go-wargaming/blob/master/LICENSE.md)

[![Go report card](https://goreportcard.com/badge/github.com/IceflowRE/go-wargaming)](https://goreportcard.com/report/github.com/IceflowRE/go-wargaming)
[![Go Reference](https://pkg.go.dev/badge/github.com/IceflowRE/go-wargaming.svg)](https://pkg.go.dev/github.com/IceflowRE/go-wargaming)
![API Coverage](https://img.shields.io/badge/API%20coverage-188%20%2F%20188-green.svg)

---

Go client for accessing [Wargaming.net Public API](https://developers.wargaming.net/documentation/guide/getting-started/).

This api client is generated automatically based on the provided documentation of Wargaming.
The documentation is often inaccurate, especially the return types (object, list, map), so manual fixes afterwards are required. If something was not corrected, please open an issue.

## Installation

go-wargaming is compatible with modern Go releases in module mode, with Go installed:

```shell
go get github.com/IceflowRE/go-wargaming/v2.0.0
```

will resolve and add the package to the current development module, along with its dependencies.

Or just import it and run `go get` afterwards

```go
import "github.com/IceflowRE/go-wargaming/v2.0.0/wargaming"
```

## Usage

```go
import "github.com/IceflowRE/go-wargaming/v2.0.0/wargaming" // with go modules
import "github.com/IceflowRE/go-wargaming/wargaming"        // without go modules
```

Create a new Wargaming Client and use the different services to access the different sections of the Wargaming API.
The services of a client are the same as the Wargaming API sections in their documentation.
```go
client := wargaming.NewClient("a7f838650dcb008552966db063eeeb35", nil)
// get World of Tanks EU accounts starting with "Yzne"
res, err := client.Wot.AccountList(context.Background(), wargaming.RealmEu, "Yzne", nil)
// print them to console
if err != nil {
	fmt.Println(err.Error())
} else {
	for _, value := range res {
		fmt.Println(value)
	}
}
```

Also, the client can be customized.
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
	Type_: wargaming.String("exact"),
})
```

All structs for Wargaming resources use pointer values for all non-repeated fields. This allows distinguishing between unset fields and those set to a zero-value. Helper functions have been provided to easily create these pointers for string, bool, and int values. For example:
```go
options := &wot.AccountListOptions{
	Type_: wargaming.String("exact"),
}
```

NOTE: Using the context package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request. In case there is no context available, then `context.Background()` can be used as a starting point.

## Development

To run the generation of the API code and check for their unit tests:

It will also edit `wargaming/client.go` and `codecov.yml`.

```shell
python ./gen_api.py
go fmt ./wargaming/
```

The client test requires an environment variable `WARGAMING_API_ID` with the API ID.

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