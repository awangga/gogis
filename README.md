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
    fmt.Println(GetPublicIP())
    var db MongoGeometry
	db.MongoString = "mongodb+srv://gogis:gogis@cluster0.wghp85v.mongodb.net/?retryWrites=true&w=majority"
	db.DBName = "location"
	db.CollectionName = "villages"
	db.LocationField = "border"
    fmt.Println(GetLocation(db, 107.57297533576039, -6.872079914985439))
}
```

## Testing

```
$ go test
```

## Tagging
develop and publish new version of gogis

```
$ git tag v0.1.2
$ git push origin --tags
$ go list -m github.com/awangga/gogis@v0.1.2
```
## Environment
Setting up environment

```
GOPROXY=proxy.golang.org
```