# github.com/awangga/gogis

## Usage

### Initialize your module

```
$ go mod init example.com/my-gogis-demo
```

### Get the gogis module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/awangga/gogis@v0.1.0
```

```go
package main

import (
    "fmt"

    "github.com/awangga/gogis"
)

func main() {
    fmt.Println(gogis.Add(2,3))
    fmt.Println(gogis.Subtract(2,3))
}
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v0.1.0
$ git push origin --tags
```
