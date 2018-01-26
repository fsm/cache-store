<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>

# Cache Store

This package is a in-memory cache implementation of a [fsm](https://github.com/fsm/fsm).[Store](https://github.com/fsm/fsm/blob/master/fsm.go#L26-L29).

Not recommended for use in production, but useful in development.

## Getting Started

```go
package main

import "github.com/fsm/cache-store"

func main() {
    store := cachestore.New()
    // ...
}
```

# License

[MIT](LICENSE.md)
