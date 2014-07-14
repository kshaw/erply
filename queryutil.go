package erply
import(
  "encoding/json"
  "io/ioutil"
  "fmt"
  "net/http"
)

func _makeQueryString(queryParams map[string]string) (queryString string) {
  queryParams["username"] = config["username"]
  queryParams["password"] = config["password"]
  queryParams["clientCode"] = config["clientCode"]
  if (request != "verifyUser") {
    queryParams["sessionKey"] = sessionKey
  }
  for k, v := range queryParams {
    queryString = queryString + "&" + k + "=" + v
  }

  return queryString
}

func _combinePath(queryString string, request string) (fullPath string) {
  return "https://" + config["clientCode"] + ".erply.com/api/?request=" + request + queryString
}

func _parseSessionKey(res *http.Response) (sessionKey string){
  var data = _serializeResponse(res)
  sessionKey = data.Records[0]["sessionKey"].(string)
  return sessionKey
}

func _serializeResponse(res *http.Response)(unmarData Response){
  body, err := ioutil.ReadAll(res.Body)
  if err != nil{
    fmt.Println(err)
  }
  res.Body.Close()
  var jsonByte []byte = []byte(body)
  json.Unmarshal(jsonByte, &unmarData)
  return unmarData
}

