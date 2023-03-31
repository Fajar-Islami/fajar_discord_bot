package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Fajar-Islami/fajar_discord_bot/model"
	"github.com/Fajar-Islami/fajar_discord_bot/service"
	"github.com/PuerkitoBio/goquery"
)

type AIBot interface {
	SearchBot(text string) string
}

type AIBotImpl struct {
	APIKEY string
}

func NewAIBot(APIKEY string) AIBot {
	return &AIBotImpl{
		APIKEY: APIKEY,
	}
}

func (ai *AIBotImpl) SearchBot(text string) string {
	reqBody := model.AIReq{
		EnableGoogleResults: true,
		EnableMemory:        true,
		InputText:           text,
	}
	payload, err := json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
		return "Error Marshal SearchBot AI"
	}

	post := service.Post("https://api.writesonic.com/v2/business/content/chatsonic?engine=premium", bytes.NewBuffer(payload))
	post.Header("accept", "application/json")
	post.Header("content-type", "application/json")
	post.Header("X-API-Key", ai.APIKEY)

	resp, err := post.Do()

	if err != nil {
		log.Println(err)
		return "Error Get Search Bot"
	}

	var r model.AIRes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "Error Read Search Bot Body"
	}

	fmt.Println("response body : ", string(body))

	// @@TODO : for response code
	// if strings.Contains(fmt.Sprint(body), "\\n\\n`") {
	// 	var resp map[string]any
	// 	err = json.Unmarshal(body, &resp)
	// 	if err != nil {
	// 		log.Println(err)
	// 		log.Println("body = ", string(body))
	// 		return fmt.Sprintf("Error UnMarshal SearchBot AI Res Body : %s \n\n", err)
	// 	}

	// 	return resp["message"].(string)

	// }

	// jsonString := `{
	//     "message": "Berikut adalah contoh program \"Hello, World!\" yang sederhana menggunakan bahasa pemrograman Go:\n\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n```\n\nPenjelasan kode di atas:\n\n1. Baris pertama merupakan sebuah deklarasi package. `package main` menyatakan bahwa file tersebut merupakan sebuah executable program, bukan sebuah library.\n2. Baris kedua mengimpor package `fmt`, yang menyediakan fungsi-fungsi untuk formatting dan output.\n3. Baris ketiga mendefinisikan fungsi `main`. Setiap program Go selalu dimulai dari fungsi `main`.\n4. Baris keempat mencetak string \"Hello, World!\" ke console menggunakan fungsi `Println` dari package `fmt`.\n\nUntuk menjalankan program di atas, Anda dapat melakukan hal berikut ini:\n\n1. Simpan kode di atas ke dalam file dengan nama `main.go`.\n2. Buka command prompt atau terminal di direktori yang sama dengan file `main.go`.\n3. Ketikkan perintah `go run main.go` untuk menjalankan program.\n\nProgram akan menampilkan output \"Hello, World!\" di console.",
	//     "image_urls": []
	// }`
	// jsonString := `{
	// 	"message": "Berikut adalah contoh program \"Hello, World!\" yang sederhana menggunakan bahasa pemrograman Go:\n\n\"\"\"go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n\"\"\"\n\nPenjelasan kode di atas:\n\n1. Baris pertama merupakan sebuah deklarasi package. `package main` menyatakan bahwa file tersebut merupakan sebuah executable program, bukan sebuah library.\n2. Baris kedua mengimpor package `fmt`, yang menyediakan fungsi-fungsi untuk formatting dan output.\n3. Baris ketiga mendefinisikan fungsi `main`. Setiap program Go selalu dimulai dari fungsi `main`.\n4. Baris keempat mencetak string \"Hello, World!\" ke console menggunakan fungsi `Println` dari package `fmt`.\n\nUntuk menjalankan program di atas, Anda dapat melakukan hal berikut ini:\n\n1. Simpan kode di atas ke dalam file dengan nama `main.go`.\n2. Buka command prompt atau terminal di direktori yang sama dengan file `main.go`.\n3. Ketikkan perintah `go run main.go` untuk menjalankan program.\n\nProgram akan menampilkan output \"Hello, World!\" di console.",
	// 	"image_urls": []
	// }`
	// var body = []byte(jsonString)

	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println(err)
		log.Println("body = ", string(body))
		return fmt.Sprintf("Error UnMarshal SearchBot AI Res Body : %s \n\n", err)
	}

	if r.Detail != "" {
		return r.Detail
	}

	// return r["message"].(string)

	replacer := strings.NewReplacer("<br/>", "\n", "<b>", "**", "</b>", "**", "<span>", "", "</span>", "")
	r.Message = replacer.Replace(r.Message)

	strSplit := strings.Split(r.Message, "\n\n")

	var builder strings.Builder
	builder.WriteString(strSplit[0] + "\n\n")
	// builder.WriteString(strSplit[0] + "\n\n")
	// // str := "```PHP\n// Contoh kode PHP Swoole\n<?php\n$server = new Swoole\\Http\\Server(\"127.0.0.1\", 9501);\n\n$server->on(\"Start\", function ($server) {\n echo \"Server started\\n\";\n});\n\n$server->on(\"Request\", function ($request, $response) {\n $response->header(\"Content-Type\", \"text/plain\");\n $response->end(\"Hello World\\n\");\n});\n\necho \"Server is starting...\\n\";\n\n$server->start();\n```"

	// str := "```PHP\n// Contoh kode PHP Swoole\n<?php\n$server = new Swoole\\Http\\Server(\"127.0.0.1\", 9501);\n\n$server->on(\"Start\", function ($server) {\n echo \"Server started\\n\";\n});\n\n$server->on(\"Request\", function ($request, $response) {\n $response->header(\"Content-Type\", \"text/plain\");\n $response->end(\"Hello World\\n\");\n});\n\necho \"Server is starting...\\n\";\n\n$server->start();\n```"

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(r.Message))
	if err != nil {
		log.Println("Error when parsing html : ", err)
		return err.Error()
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			href = ""
		}
		builder.WriteString(fmt.Sprintf("[%d] %s. ( %s ) \n", i+1, s.Text(), href))
	})

	return builder.String()
}
