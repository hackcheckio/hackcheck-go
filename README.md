# Hackcheck-go

<p>
    <a href="https://pkg.go.dev/github.com/hackcheckio/hackcheck-go?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/hackcheckio/hackcheck-go"><img alt="Go report", src="https://goreportcard.com/badge/github.com/hackcheckio/hackcheck-go"></a>
</p>

A go wrapper for [hackcheck.io](https://hackcheck.io)'s public API


## Installation
```sh-session
go get -u github.com/hackcheckio/hackcheck-go
```

## Example usage

Simple email lookup that prints leaked passwords:

```go
import (
    "fmt"

	"github.com/hackcheckio/hackcheck-go"
)

func main() {
    hc := hackcheck.New("your hackcheck api key")

    result, err := hc.LookupEmail("your@email.com")
    if err != nil {
        panic(err)
    }

    for _, r := range result {
        fmt.Printf("Password found: '%s'", r.Password)
    }
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