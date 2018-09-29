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

func doReq(passedReqOps requestOptions) (string, error) {
	constructedReq, err := newRequest(passedReqOps)
	if err != nil {
		return "", err
	}
	util.Log(constructedReq)

	client := &http.Client{}

	res, err := client.Do(constructedReq)
	//defer res.Body.Close()

	if err != nil {
		fmt.Println("Something happened:", err)
		return "", err
	}

	if res.StatusCode == http.StatusOK {
		res, _ := ioutil.ReadAll(res.Body)
		return string(res), nil
	}
	if res.StatusCode == http.StatusCreated {
		res, _ := ioutil.ReadAll(res.Body)
		return string(res), nil
	}
	if res.StatusCode == http.StatusNoContent {
		if loginURL == passedReqOps.URL {
			saveToken(res.Header.Get("Authorization"))
		}
		return "", nil
	}

	if res.StatusCode == http.StatusForbidden {
		clearToken()
		util.SayLastWords("you need to login")
	}

	if res.StatusCode == http.StatusInternalServerError {
		f, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		util.Debug(string(f))
		util.SayLastWords("Service is unavailable. Try later again.")
	}

	return "", errors.New("unexpected things happened")
}

func newRequest(passedOps requestOptions) (req *http.Request, err error) {
	method := "GET"
	if passedOps.Method != "" {
		util.Debug("Http method is set to " + passedOps.Method)
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
	if token := readToken(); token != "" {
		req.Header.Set("Authorization", token)
		util.Debug("Added token to header")
	}

	return req, nil
}
