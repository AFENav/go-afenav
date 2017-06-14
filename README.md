# go-afenav
Go wrappers for AFE Navigator APIs


```go
package main

import (
	"github.com/AFENav/go-afenav"
)

func main() {
	service := afenav.New("https://afenav.company.com/")
	service.Login("API User", "123456")
    defer service.Logout()
}

```