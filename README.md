# api-client

`api-client` is a reusable Go library for sending external API requests with features like debugging, retries, and request/response logging.

Author: Mohamed Riyad  
Email: m@ryad.dev  
Website: [ryad.dev](https://ryad.dev)

## Installation

To install `api-client`, use `go get`:

```sh
go get github.com/RyadPasha/api-client
```

## Usage Example

See `example.go` for a usage example:

```go
/*
Package apiClient provides a reusable library for sending external API requests with features like debugging, retries, and request/response logging.

Author: Mohamed Riyad
Email: m@ryad.dev
Website: https://ryad.dev
*/

package main

import (
	"fmt"
	"github.com/yourusername/apiClient"
	"net/http"
	"time"
)

func main() {
	client := apiClient.NewClient("https://api.example.com", true, 3, 2*time.Second)

	req := &apiClient.APIRequest{
		Method:   http.MethodGet,
		Endpoint: "/data",
		Headers:  map[string]string{"Content-Type": "application/json"},
	}

	resp, err := client.SendRequest(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", string(resp.Body))
}
```

## Features

-   Send HTTP requests with ease
-   Automatic retries with configurable retry logic
-   Debugging mode to log requests and responses
-   Support for custom headers and request bodies

## License

MIT License
