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

      queryParams := map[string]string{
        'abv': '6',
      }
      res, err := erply.Get("getProducts", queryParams)
    }

```
