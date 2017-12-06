Go Service Container ![Last release](https://img.shields.io/github/release/euskadi31/go-service.svg)
====================

[![Go Report Card](https://goreportcard.com/badge/github.com/euskadi31/go-service)](https://goreportcard.com/report/github.com/euskadi31/go-service)

| Branch  | Status | Coverage |
|---------|--------|----------|
| master  | [![Build Status](https://img.shields.io/travis/euskadi31/go-service/master.svg)](https://travis-ci.org/euskadi31/go-service) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-service/master.svg)](https://coveralls.io/github/euskadi31/go-service?branch=master) |
| develop | [![Build Status](https://img.shields.io/travis/euskadi31/go-service/develop.svg)](https://travis-ci.org/euskadi31/go-service) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-service/develop.svg)](https://coveralls.io/github/euskadi31/go-service?branch=develop) |


Simple Dependency Injection Container for Golang

## Example

```go
import "github.com/euskadi31/go-service"

type MyService struct {}

sc := service.New()

// Define service
sc.Set("my.service", func(c *service.Container) interface{} {
    return &MyService{}
})

// Call service 
myService := sc.Get("my.service").(*MyService)

```


## License

go-service is licensed under [the MIT license](LICENSE.md).
