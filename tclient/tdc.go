package tclient

import (
  "fmt"
  "net/http"
  "errors"
)

func req(url string) (interface{}, error){
  res, err := http.Get(url)
  if err != nil {
    fmt.Println("Something happened:", err)
    return nil, err
  }
  if res.StatusCode != http.StatusOK {
    return nil, errors.New("Status not 200, ")
  }
  return "", nil
}

func IsTdUp() bool {
  _, err := req(tdUp)
  if err != nil {
    return false
  }
  return true
}

func IsUmUp() bool {
  _, err := req(umUp)
  if err != nil {
    return false
  }
  return true
}
