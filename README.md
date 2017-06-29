Go Service Container
====================

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
