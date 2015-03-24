### go-PollarRho

a Go package to Pollard-Rho-factor algorithm

### Getting

    go get -u bitbucket.org/kimitoboku/go-PollarRho

### Importing

    import "bitbucket.org/kimitoboku/go-PollarRho"

### How to

```go
package main

import (
"bitbucket.org/kimitoboku/go-PollarRho"
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
