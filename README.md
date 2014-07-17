# goDiscoverEtcd
goDiscoverEtcd is a simple library to query `https://discovery.etcd.io/new` and return a single URL which you can use to build your etcd cluster.

## Installation
Just include the library in your go program:

```go
import "github.com/doublerr/goDiscoverEtcd"
```

## Usage
Here is an exmaple app:

```go
package main

import "fmt"
import "github.com/doublerr/goDiscoverEtcd"

func main() {
    url := goDiscoverEtcd.GetURL()
    fmt.Printf(url)
}
```
## Contributing

1. Fork
2. code
3. contribute

## Authors
Ryan Richard (<ryanrichard07@gmail.com>)