# Hackcheck-go

<p>
    <a href="https://pkg.go.dev/github.com/hackcheckio/hackcheck-go?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/hackcheckio/hackcheck-go"><img alt="Go report", src="https://goreportcard.com/badge/github.com/hackcheckio/hackcheck-go"></a>
</p>

Official go library for the [hackcheck.io](https://hackcheck.io) API

- [Hackcheck-go](#hackcheck-go)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [Lookup methods](#lookup-methods)
  - [Other options](#other-options)


## Installation
```sh-session
go get -u github.com/hackcheckio/hackcheck-go
```

## Quick start

Simple email lookup that prints leaked passwords:

```go
import (
    "fmt"

	"github.com/hackcheckio/hackcheck-go"
)

func main() {
    // Get an api key by purchasing a developer plan https://hackcheck.io/plans
    hc := hackcheck.New("MY_API_KEY")

    result, err := hc.LookupEmail("your@email.com")
    if err != nil {
        panic(err)
    }

    for _, r := range result {
        fmt.Println("Database:", r.Source.Name)
        fmt.Println("Date:", r.Source.Date)
        fmt.Println("Password:", r.Password)
        fmt.Println("Username:", r.Username)
        fmt.Println("IP:", r.IP)
        fmt.Println("-------")
    }
    fmt.Println()

    // Check your ratelimits
	fmt.Println("Current rate limit:", hc.CurrentRatelimit)
	fmt.Println("Allowed rate limit:", hc.AllowedRatelimit)
}
```

## Lookup methods

- `LookupEmail()`
- `LookupUsername()`
- `LookupName()`
- `LookupPassword()`
- `LookupIP()`
- `LookupPhone()`
- `LookupDomain()`

## Other options

```go
// Using with cache
// 10 represents the cache size
hackcheck.New("MY_API_KEY", hackcheck.WithCache(10))

// With custom http client
hackcheck.New("MY_API_KEY", hackcheck.WithHttp(&http.Client{...}))

// With context
hackcheck.New("MY_API_KEY", hackcheck.WithContext(context.Background()))
```