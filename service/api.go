package service

import (
	"fmt"
	"strings"
)

// func GetAPI(uri string) (response *http.Response) {
// 	response, err := http.Get(uri)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

func ListCommand(botname string) string {
	var str strings.Builder
	str.WriteString("Fajar Bot command list :\n")
	str.WriteString(fmt.Sprint("- `", botname, " jokes` = Get one random joke. \n"))
	str.WriteString(fmt.Sprint("- `", botname, " joktod` = Get today's joke. \n"))
	str.WriteString(fmt.Sprint("- `", botname, " search=<text you want to search>` = Data search using AI. \n"))
	// str.WriteString(fmt.Sprint("- `", botname, " rcelist` = List Programming Language\n"))
	str.WriteString(fmt.Sprint("- `", botname, " env`= Check environment\n"))
	str.WriteString(fmt.Sprint("- `", botname, " sholat` = COMING SOON!!\n"))
	str.WriteString(fmt.Sprint("- `", botname, " search` = Search Engine use google.com \n"))
	str.WriteString(fmt.Sprint("- `", botname, " translate-langlist` = List supported language \n"))
	str.WriteString(fmt.Sprint("- `", botname, " translate-codelang <language>` = Examine the language code \n"))
	str.WriteString(fmt.Sprint("- `", botname, " translate-detectlang` = COMING SOON!!\n"))
	str.WriteString(fmt.Sprint("- `", botname, " translate` = COMING SOON!!\n"))
	str.WriteString(fmt.Sprint("- `", botname, " ping` = test ping\n"))
	str.WriteString(fmt.Sprint("- `", botname, " pong` = test ping\n"))
	str.WriteString(fmt.Sprint("- `", botname, " intro` = About this bot\n"))
	str.WriteString(fmt.Sprint("- `", botname, " intro` = About this bot\n"))
	str.WriteString(fmt.Sprint("- `", botname, " contribute` = Link Repository\n"))

	return str.String()
}
