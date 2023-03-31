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
	log.Println("Searching ", text)
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

	// jsonString := `{
	// 	"message": "Google adalah sebuah perusahaan multinasional Amerika Serikat yang berfokus pada jasa dan produk Internet, terutama mesin pencari. Google didirikan pada tahun 1998 oleh Larry Page dan Sergey Brin. Misi utama dari Google adalah untuk mengumpulkan informasi dunia dan membuatnya dapat diakses dan bermanfaat oleh semua orang. Google menyediakan berbagai layanan dan produk online seperti email, browser web, perangkat lunak produktivitas, ponsel dan aplikasi, alat pemetaan, e-book, iklan internet, serta berbagai video dan situs jejaring sosial. Google adalah perusahaan mesin pencari terbesar di dunia dan mengoperasikan lebih dari satu juta server di beberapa pusat data di seluruh dunia [1][2][3].<br/><br/><b>References:</b><br/><span>[1] <a href='https://id.wikipedia.org/wiki/Google' target='_blank' class='text-purple-1 underline'>Google - Wikipedia bahasa Indonesia, ensiklopedia bebas</a></span><br/><span>[2] <a href='https://www.idn.id/pengertian-sejarah-dan-fungsi-google-sebagai-search-engine-terbesar/' target='_blank' class='text-purple-1 underline'>Pengertian, Sejarah, dan Fungsi Google sebagai Search ...</a></span><br/><span>[3] <a href='https://kumparan.com/berita-update/pengertian-sejarah-dan-fungsi-google-sebagai-search-engine-terbesar-1wwPN1SYZo1' target='_blank' class='text-purple-1 underline'>Pengertian, Sejarah, dan Fungsi Google sebagai Search ...</a></span><br/> [Link text](https://www.idn.id/pengertian-sejarah-dan-fungsi-google-sebagai-search-engine-terbesar/)",
	// 	"image_urls": []
	// }`
	// var body = []byte(jsonString)

	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println(err)
		log.Println("body = ", string(body))
		return fmt.Sprintf("Error UnMarshal SearchBot AI Res Body : %s \n\n", err)
	}

	// return r["message"].(string)

	replacer := strings.NewReplacer("<br/>", "\n", "<b>", "**", "</b>", "**", "<span>", "", "</span>", "")
	r.Message = replacer.Replace(r.Message)

	strSplit := strings.Split(r.Message, "\n\n")

	var builder strings.Builder
	builder.WriteString(strSplit[0] + "\n\n")

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
