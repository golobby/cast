[![GoDoc](https://godoc.org/github.com/golobby/cast?status.svg)](https://godoc.org/github.com/golobby/cast)
[![CI](https://github.com/golobby/cast/actions/workflows/ci.yml/badge.svg)](https://github.com/golobby/cast/actions/workflows/ci.yml)
[![CodeQL](https://github.com/golobby/cast/workflows/CodeQL/badge.svg)](https://github.com/golobby/cast/actions?query=workflow%3ACodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/golobby/cast)](https://goreportcard.com/report/github.com/golobby/cast)
[![Coverage Status](https://coveralls.io/repos/github/golobby/cast/badge.svg?v=1.2)](https://coveralls.io/github/golobby/cast?branch=master)

# Cast
GoLobby Cast is a lightweight casting package for Go projects.

## Documentation

### Required Go Versions
It requires Go `v1.16` or newer versions.

### Installation
To install this package, run the following command in your project directory.

```bash
go get github.com/golobby/cast
```

### Usage Example
The following examples demonstrate how to use the GoLobby Cast package.

```go
val, err := cast.FromString("666", "int32")
// `val` type: int32
// `val` value: 666
```

```go
val, err := cast.FromString("1", cast.Bool)
// `val` type: bool
// `val` value: true
```

### Supported Types
Currently, the GoLobby Cast supports the following types and related array types:
* Int (Int8 .. Int64)
* Uint (Uint8 .. Uint64)
* Float32 and Float64
* Bool
* String

## License
GoLobby Cast is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
