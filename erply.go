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
  sessionKey = _getSessionKey()
  c = newClient()
  return c != nil
}

func Get(request string, queryParams map[string]string) (res *http.Response, err error) {
  queryString := _makeQueryString(queryParams, request)
  fullPath    := _combinePath(queryString, request)

  return c._get(fullPath)
}

func Serialize(res *http.Response)(response Response){
  response = _serializeResponse(res)
  return response
}

func _getSessionKey() (string){
  var queryParams = make(map[string]string)
  res, err := Get("verifyUser", queryParams)
  if err != nil{
    fmt.Println(err)
  }
  sessionKey := _parseSessionKey(res)
  return sessionKey
}
