[![doc](https://godoc.org/github.com/golobby/cast?status.svg)](https://godoc.org/github.com/golobby/cast)
[![build](https://travis-ci.com/golobby/cast.svg?branch=master)](https://travis-ci.com/golobby/cast)
[![report](https://goreportcard.com/badge/github.com/golobby/cast)](https://goreportcard.com/report/github.com/golobby/cast)
[![coverage](https://coveralls.io/repos/github/golobby/cast/badge.svg?branch=master)](https://coveralls.io/github/golobby/cast?branch=master)
# GoLobby Cast
A lightweight yet powerful casting package for Go projects

## Documentation

### Required Go Versions
It requires Go `v1.11` or newer versions.

### Installation
To install this package, run the following command in your project directory.

```bash
go get github.com/golobby/cast
```

### Sample Usage
```go
val, err := cast.FromString("666", "int32")
// `val` type: int32
// `val` value: 666
```

## License
GoLobby Cast is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
