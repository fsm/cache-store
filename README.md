<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>

# Cache Store

[![Version](https://img.shields.io/github/tag/fsm/cache-store.svg)](https://github.com/fsm/cache-store/releases)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/fsm/cache-store/blob/master/LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/fsm/cache-store)](https://goreportcard.com/report/github.com/fsm/cache-store)
[![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/fsm/Lobby)

This package is a in-memory cache implementation of a [fsm.Store](https://github.com/fsm/fsm/blob/master/fsm.go#L33-L38).

> Note: This store is NOT recommended for use in production, but is the quickest way to get started in development.

## Getting Started

```go
package main

import "github.com/fsm/cache-store"

func main() {
    store := cachestore.New()
    // ...
}
```
