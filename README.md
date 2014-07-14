#Erply API Wrapper in Golang

## To Use:

```
    import (
    "github.com/kshaw/erply"
    )

    func getProducts() {
      config := map[string]string {
        "username": "username"
        "password": password",
        "clientCode": "clientCode",
      }

      queryParams := map[string]string{}
      res, err := erply.Get("getProducts", queryParams)
    }

```
