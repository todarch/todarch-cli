package tclient

import (
	"fmt"
)

type loginReq struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func Authenticate(username string, password string) {
	_, err := doReq(requestOptions{
		Method: "POST",
		URL:    loginURL,
		Body:   loginReq{username, password},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully logged in.")
	}
}
