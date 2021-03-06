# LetsCloud: A Go library for the LetsCloud Cloud API

The library’s documentation is available at [GoDoc](https://godoc.org/github.com/letscloud-community/letscloud-go),
the public API documentation is available at [developers.letscloud.io](https://developers.letscloud.io/).

## Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/letscloud-community/letscloud-go"
)

func main() {
    client, err := letscloud.New(token) // or use `export LETSCLOUD_API_KEY=your_api_key`
    if err != nil {
        log.Fatal(err)
    }

    profile, err := client.Profile()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(profile)
}
```

## License

MIT license