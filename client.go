package erply

import(
  "net/http"
)

type Response struct {
  Status map[string]interface{}
  Records []map[string]interface{}
}

type client struct {
  sessionKey string
  config map[string]string
}

func newClient() (c *client) {
  return &client{sessionKey:sessionKey, config: config}
}

func (*client) _get(fullPath string) (res *http.Response, err error) {
  return http.Get(fullPath)
}
