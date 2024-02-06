# returnmagicfinder

Static analysis tool for Go that finds magic number in return statements.

## Example
```golang
package main

import "fmt"

func calculateDiscount(price float64) float64 {
    return price * 0.9 // 0.9 is a magic number
}

func main() {
    price := 1000.0
    discountedPrice := calculateDiscount(price)
    fmt.Println("Discounted Price:", discountedPrice)
}

```

```
./main.go:6:20: magic number used in return statement
```

## Install
```shell
go install github.com/mshr0969/returnmagicfinder/cmd/returnmagicfinder@v.1.0.0
```

## Usage
```sehll
go vet -vettool=$(which returnmagicfinder) main.go
```
