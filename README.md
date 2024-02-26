# HackCheck-go

<p>
    <a href="https://pkg.go.dev/github.com/hackcheckio/hackcheck-go?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
</p>

Official go library for the [hackcheck.io](https://hackcheck.io) API

- [HackCheck-go](#hackcheck-go)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [Getting your api key](#getting-your-api-key)
  - [Other examples](#other-examples)

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
		panic(err)
	}

	fmt.Println(resp.Results)
}
```

## Getting your api key

1. Visit https://hackcheck.io/profile
2. Add your IP address in the "Authorized IP Addresses" text area and click Update
3. Copy your API key

## Other examples

<details>
<summary>Breach Monitors</summary>

```go
import (
	"fmt"

	"github.com/hackcheckio/hackcheck-go"
)

func main() {
	hc := hackcheck.New("MY_API_KEY")

	// Listing breach monitors
	monitors, err := hc.GetMonitors()
	if err != nil {
		panic(err)
	}


	fmt.Println(monitors.AssetMonitors)
	fmt.Println(monitors.DomainMonitors)

	// Getting a monitor
	myAssetMonitor, err := hc.GetAssetMonitor("...") // or hc.GetDomainMonitor
	if err != nil {
		panic(err)
	}

	fmt.Println(myAssetMonitor.Status)
	fmt.Println(myAssetMonitor.Asset)

	// Updating a monitor
	domainMonitor, err := hc.UpdateDomainMonitor("id123123123", &hackcheck.UpdateDomainMonitorParams{
		Domain: "example.com",
		NotificationEmail: "notifications@example.com",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(domainMonitor.Domain)
}
```

</details>

<details>
<summary>Filtering databases</summary>

```go
import (
	"fmt"

	"github.com/hackcheckio/hackcheck-go"
)

func main() {
	hc := hackcheck.New("MY_API_KEY")

	// This will only yield results from "website.com" and "website.org"
	// Use hackcheck.SearchFilterTypeIgnore if you want to ignore the databases
	resp, err := hc.Search(
		&SearchOptions{
			Field: hackcheck.SearchFieldEmail,
			Query: "example@example.com",
			Filter: &SearchFilterOptions{
				Type: hackcheck.SearchFilterTypeUse,
				Databases: []string{"website.com", "website.org"},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Results)
}
```

</details>

<details>
<summary>Checking if a query exists</summary>

```go
import (
	"fmt"

	"github.com/hackcheckio/hackcheck-go"
)

func main() {
	hc := hackcheck.New("MY_API_KEY")

	// Returns true if the query is found
	exists, err := hc.Check(
		&CheckOptions{
			Field: hackcheck.SearchFieldEmail,
			Query: "example@example.com",
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(exists)
}
```

</details>

<details>
<summary>Custom http client & context</summary>

```go
import (
	"fmt"

	"github.com/hackcheckio/hackcheck-go"
)

func main() {
	customCtx := context.Background()
	customHttpClient := &http.Client{}

	hc := hackcheck.New("MY_API_KEY", hackcheck.WithContext(customCtx), hackcheck.WithHTTP(customHttpClient))
	// ...
}
```

</details>
