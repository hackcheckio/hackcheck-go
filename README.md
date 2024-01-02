# HackCheck-go

<p>
    <a href="https://pkg.go.dev/github.com/hackcheckio/hackcheck-go?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/hackcheckio/hackcheck-go"><img alt="Go report", src="https://goreportcard.com/badge/github.com/hackcheckio/hackcheck-go"></a>
</p>

Official go library for the [hackcheck.io](https://hackcheck.io) API

- [HackCheck-go](#hackcheck-go)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [Lookup methods](#lookup-methods)
  - [Other options](#other-options)

## Installation

```sh-session
go get -u github.com/hackcheckio/hackcheck-go
```

## Quick start

Example usage

```go
import (
    "fmt"

	"github.com/hackcheckio/hackcheck-go"
)

func main() {
    hc := hackcheck.New("MY_API_KEY")

    resp, err := hc.Search(
        &SearchOptions{
            Field: hackcheck.SearchFieldEmail,
            Query: "example@example.com",
        },
    )

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(resp.Results)
}
```

## Getting your api key

1. Visit https://hackcheck.io/profile
2. Add your IP address in the "Authorized IP Addresses" text area and click Update
3. Copy your API key
