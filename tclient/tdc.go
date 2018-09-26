package tclient

func IsTdUp() bool {
	_, err := doReq(requestOptions{URL: tdUp})
	if err != nil {
		return false
	}
	return true
}

func IsUmUp() bool {
	_, err := doReq(requestOptions{URL: umUp})
	if err != nil {
		return false
	}
	return true
}
