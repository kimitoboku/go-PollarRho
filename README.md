### go-PollarRho

a Go package to Pollard-Rho-factor Alg

### Importing

    import "bitbucket.org/kimitoboku/go-PollarRho"

### How to

```go
num := 323911866737
list := pollarrho.Factor(num)
fmt.Println(list) 
```

```
$ go build
$ ./hoge
[5477 6311 9371]
```
