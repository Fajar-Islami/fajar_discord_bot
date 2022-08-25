package service

import (
	"fmt"
	"net/http"
)

func GetAPI(uri string) (response *http.Response) {
	response, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
	}
	return
}
