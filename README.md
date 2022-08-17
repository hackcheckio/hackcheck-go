# Hackcheck-go
A go wrapper for [hackcheck.io](https://hackcheck.io)'s API

## Installation
```bash
go get -u github.com/hackcheckio/hackcheck-go
```

## Usage

You need an API key to utilize the Hackcheck API, which you can get here:
https://hackcheck.io/profile

NOTE: The API is only available to users with a Developer plan

```go
import (
	"github.com/hackcheckio/hackcheck-go"
)

func main() {
    hc := hackcheck.New("your hackcheck api key")

    result, err := hc.LookupEmail("your@email.com")
    if err != nil {
        panic(err)
    }

    for _, r := range result {
        fmt.Println(r)
    }
}
```

## Lookup Methods

```go
LookupEmail
LookupUsername
LookupName
LookupPassword
LookupIP
LookupPhone
LookupDomain
```


## Result types

```go
Email    string
Password string
Username string
IP       string
Phone    string
Source   struct {
    Name        string
    Date        string
    Description string
}
```