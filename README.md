# cache
A simple in-memory cache package

Add in go.mod file 
```
module github.com/viamarvin/cache
```

## Example
```
package main

import (
	"fmt"
	"github.com/viamarvin/cache"
)

func main() {
	c := cache.New()
	c.Set("RUB", 1)
	c.Set("KZT", 6.20)
	c.Set("USD", 0.014)
	
    c.Show()

    c.Delete("KZT")
    fmt.Println("RUB", c.Get("RUB"))
    fmt.Println("USD", c.GET("USD"))
}
```