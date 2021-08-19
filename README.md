[![GoDoc](https://godoc.org/github.com/golobby/cast?status.svg)](https://godoc.org/github.com/golobby/cast)
[![CI](https://github.com/golobby/cast/actions/workflows/ci.yml/badge.svg)](https://github.com/golobby/cast/actions/workflows/ci.yml)
[![CodeQL](https://github.com/golobby/cast/workflows/CodeQL/badge.svg)](https://github.com/golobby/cast/actions?query=workflow%3ACodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/golobby/cast)](https://goreportcard.com/report/github.com/golobby/cast)
[![Coverage Status](https://coveralls.io/repos/github/golobby/cast/badge.svg)](https://coveralls.io/github/golobby/cast?branch=master)

# Cast
GoLobby Cast is lightweight casting package for Go projects.

## Documentation

### Required Go Versions
It requires Go `v1.11` or newer versions.

### Installation
To install this package, run the following command in your project directory.

```bash
go get github.com/golobby/cast
```

### Sample Usage
The following example demonstrates how to use GoLobby Cast package.

```go
val, err := cast.FromString("666", "int32")
// `val` type: int32
// `val` value: 666

val, err := cast.FromString("1", cast.Bool)
// `val` type: bool
// `val` value: true
```

## License
GoLobby Cast is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
