# Adobe Analytics API 2.0 Client

This project provides a client library for accessing the Adobe Analytics API 2.0.

## Installation

Use the standard Go tool chain to use this library in your project.

Example:

```shell
go get -u github.com/adobe/aa-client-go
```

## Usage

Create a new Analytics client with a `http.DefaultClient`.

```go
client, err := analytics.NewClient(&analytics.Config{
    BaseURL:     "https://analytics.adobe.io/api",
    ClientID:    "<CLIENT-ID>",
    OrgID:       "<ORG-ID>",
    AccessToken: "<ACCESS-TOKEN>",
    CompanyID:   "<COMPANY-ID>",
})
```

Use specific API endpoint. For example, get all available metrics.

```go
metrics, err := client.Metrics.GetAll("<ReportSuiteID>", "en_US", false, []string{})
```

See [examples](./examples) for more information about using the various API endpoints.

### Authentication

Use [ims-go](github.com/adobe/ims-go) to get an authentication token.

```go
client, err := ims.NewClient(&ims.ClientConfig{
    URL: "https://ims-na1.adobelogin.com",
})

resp, err := client.ExchangeJWT(&ims.ExchangeJWTRequest{
    Expiration:   time.Now().Add(12 * time.Hour),
    PrivateKey:   []byte("<PRIVATE-KEY>"),
    Issuer:       "<ORG-ID>",
    Subject:      "<TECHNICAL-ACCOUNT-ID>",
    ClientID:     "<CLIENT-ID>",
    ClientSecret: "<CLIENT-SECRET>",
    MetaScope:    []ims.MetaScope{ims.MetaScopeAnalyticsBulkIngest},
})

accessToken := resp.AccessToken
```

### Handling 429 status codes

To handle `429` response status codes (returned if the API rate limit is hit), a HTTP client with retry/backoff like [go-retryablehttp](github.com/hashicorp/go-retryablehttp) can be used.

See the [examples/utils/utils.go](./examples/utils/utils.go) for an example of a custom rate limit policy.

```go
retryClient := retryablehttp.NewClient()
retryClient.RetryMax = 10
retryClient.CheckRetry = rateLimitPolicy

client, err := analytics.NewClient(&analytics.Config{
    HTTPClient:  retryClient.StandardClient(),
    BaseURL:     "https://analytics.adobe.io/api",
    ClientID:    "<CLIENT-ID>",
    OrgID:       "<ORG-ID>",
    AccessToken: "<ACCESS-TOKEN>",
    CompanyID:   "<COMPANY-ID>",
})
```

## Development

The module provides a `Makefile` with following targets.

* `all` - Runs `fmt`, `vet`, `lint` and `test`.
* `fmt` - Formats all code.  
    Runs `go fmt ./...`
* `vet` - Vets all code.  
    Runs `go vet -all ./...`
* `lint` - Lints all code.  
    Runs `golint ./...`
* `test` - Runs the test of the `analytics` package.  
    Runs `go test ./analytics -cover`
* `coverage` - Runs the tests of the `analytics` package and opens the coverage report.  
    Runs `go test -coverprofile=coverage.out ./analytics & go tool cover -html=coverage.out`

A specific target can be executed by running the following command (Linux, macOS).

```shell
make [target]
```

Windows users can simply bypass `make` and run the plain commands as indicated above or install `make` via [GNUWin32](http://gnuwin32.sourceforge.net/packages/make.htm), [Chocolatey](https://chocolatey.org/install) or [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/install-win10).

## Contributing

Contributions are welcome! Read the [Contributing Guide](./CONTRIBUTING.md) for more information.

## Licensing

This project is licensed under the Apache 2.0 License. See [LICENSE](./LICENSE) for more information.
