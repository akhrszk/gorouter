# gorouter

A lightweight powerful HTTP router (or mux) in Go.

## Features

- 100% compatible with net/http
- [Trie](https://en.wikipedia.org/wiki/Trie) based structure
- Only exact path matches
- Support path parameters, can use regexp

## Getting Started

### Installing

Install go route package with go get

```
go get -u github.com/akhrszk/gorouter
```

Start your first server. Create main.go file and add:

### Examples

```go
package main

import (
  "fmt"
  "net/http"

  "github.com/akhrszk/gorouter"
)

func hello(w http.ResponseWriter, r *http.Request, params gorouter.Params) {
  name := params["name"]
  fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
  rt := gorouter.New()
  rt.Get("/hello/:name", hello)
  http.ListenAndServe(":3000", rt)
}
```

## Named Parameters

You can capture path segments. The path parameters is given to 3rd **Handler** function parameter.
You can get the value of `:name` by `params["name"]`.

```
Pattern: /hello/:name

/hello/suzuki            match
/hello/suzuki/           match
/hello/suzuki/welcome    no match
/hello/                  no match
```

### Using regex pattern

```
Pattern: /users/:id(\\d+)

/hello/users/12          match
/hello/users/abc         no match
/hello/users/            no match
```

## License

**[MIT](LICENSE)** License.

# Authors

- **Akihiro Suzuki**
  - GitHub: [akhrszk](https://github.com/akhrszk)
  - Twitter: [@akhr_s](https://twitter.com/akhr_s)
  - Blog: [akihr.io](https://akihr.io/)
