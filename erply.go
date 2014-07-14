package erply

import(
  "net/http"
  "fmt"
)

var c *client
var sessionKey string
var config map[string]string

func SetConfigValues (configValues map[string]string) (bool){
  config = configValues
  sessionKey = getSessionKey()
  c = newClient()
  return c != nil
}

func Get(request string, queryParams map[string]string) (res *http.Response, err error) {
  queryString := _makeQueryString(queryParams)
  fullPath    := _combinePath(queryString, request)
  fmt.Println(fullPath)

  return c._get(fullPath)
}

func getSessionKey() (string){
  var queryParams = make(map[string]string)
  res, err := Get("verifyUser", queryParams)
  if err != nil{
    fmt.Println(err)
  }
  sessionKey := _parseSessionKey(res)
  fmt.Println(sessionKey)
  return sessionKey
}
