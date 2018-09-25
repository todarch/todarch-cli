package tclient

import (
	"errors"
	"fmt"
	"net/http"
)

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
