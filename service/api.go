package service

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPI(uri string) (response *http.Response) {
	response, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ListCommand(botname string) string {
	var str strings.Builder
	str.WriteString("Fajar Bot command list :\n")
	str.WriteString(fmt.Sprint("- `", botname, ".jokes` = Get single random joke\n"))
	str.WriteString(fmt.Sprint("- `", botname, ".env`= check environment\n"))
	str.WriteString(fmt.Sprint("- `", botname, ".sholat` = COMING SOON!!\n"))
	str.WriteString(fmt.Sprint("- `", botname, ".search` = COMING SOON!!\n"))
	str.WriteString(fmt.Sprint("- `", botname, ".translate` = COMING SOON!!\n"))
	str.WriteString(fmt.Sprint("- `", botname, ".ping` = test ping\n"))
	str.WriteString(fmt.Sprint("- `", botname, ".pong` = test ping\n"))

	return str.String()
}
