package tclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/todarch/todarch-cli/util"
	"io/ioutil"
	"net/http"
)

type requestOptions struct {
	Header http.Header
	Method string
	URL    string
	Body   interface{}
}

func doReq(passedReqOps requestOptions) (interface{}, error) {
	constructedReq, err := newRequest(passedReqOps)
	if err != nil {
		return nil, err
	}
	util.Log(constructedReq)
	client := &http.Client{}
	res, err := client.Do(constructedReq)
	if err != nil {
		fmt.Println("Something happened:", err)
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return "", nil
	}
	util.Log(res.StatusCode)
	f, err := ioutil.ReadAll(res.Body)
	if err != nil {
		util.Log(err)
	}
	res.Body.Close()
	if err != nil {
		util.Log(err)
	}
	util.Log(string(f))
	return nil, errors.New("Status not 200, ")
}

func newRequest(passedOps requestOptions) (req *http.Request, err error) {
	method := "GET"
	if passedOps.Method != "" {
		util.Log("Request method is set" + passedOps.Method)
		method = passedOps.Method
	}
	util.Log(passedOps.Body)
	jsonReq, err := json.Marshal(passedOps.Body)
	util.Log(string(jsonReq))
	req, err = http.NewRequest(
		method,
		passedOps.URL,
		bytes.NewBuffer(jsonReq),
	)

	if err != nil {
		return nil, err
	}

	if passedOps.Header != nil {
		req.Header = passedOps.Header
	} else {
		req.Header = http.Header{}
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
