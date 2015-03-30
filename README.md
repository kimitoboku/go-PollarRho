### go-PollarRho

a Go package to Pollard-Rho-factor algorithm

### Getting

    go get -u github.com/kimitoboku/go-PollarRho

### Importing

    import "github.com/kimitoboku/go-PollarRho"

### How to

```go
package main

import (
"github.com/kimitoboku/go-PollarRho"
"fmt"  
)
func main(){
num := 323911866737
list := pollarrho.Factor(num)
fmt.Println(list)
}
```

```
$ go build
$ ./hoge
[5477 6311 9371]
```
