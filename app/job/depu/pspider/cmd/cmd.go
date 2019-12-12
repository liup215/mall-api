package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// Url := "https://papers.xtremepape.rs/CAIE/IGCSE/Biology%20(0610)/"
	// Url := "https://papers.xtremepape.rs/CAIE/AS%20and%20A%20Level/Biology%20(9700)/"
	Url := "https://papers.xtremepape.rs/Edexcel/Advanced%20Level/Biology"

	resp, err := http.Get(Url)

	if err != nil {
		fmt.Println("网页爬取失败")
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Println("网页解析失败")
	}

	downLoad(doc, Url, "Edexcel/Alevel/")

}

func downLoad(doc *goquery.Document, url, path string) {
	doc.Find("#top .xb-canvas-menuActive .xb-content-wrapper .p-body .p-body-inner .xpp-fb .xpp-files table tbody tr").Each(func(i int, s *goquery.Selection) {
		fileName, exist1 := s.Find("td a").Attr("href")

		if fileName == "../" {
			return
		}

		fileType, exist2 := s.Find("td a img").Attr("alt")

		u := url + "/" + fileName
		patho := path + fileName

		if exist1 && exist2 {
			if fileType != "[dir]" {
				_, e := os.Stat(patho)

				if os.IsNotExist(e) {
					file, err := os.OpenFile(patho, os.O_RDWR|os.O_CREATE, 0644)

					if err != nil {
						fmt.Println("file err, " + err.Error())
					}
					defer file.Close()

					resp, _ := http.Get(u)
					cot, _ := ioutil.ReadAll(resp.Body)
					file.Write(cot)
				}
			} else {
				_, e := os.Stat(patho)

				if os.IsNotExist(e) {
					err := os.Mkdir(patho, os.ModePerm)
					if err != nil {
						fmt.Println("文件夹创建失败")
						return
					}
				}

				fmt.Println(u)
				resp, err := http.Get(u)

				if err != nil {
					fmt.Println("网页爬取失败")
				}

				doc, err := goquery.NewDocumentFromResponse(resp)
				if err != nil {
					fmt.Println("网页解析失败")
				}

				downLoad(doc, u, patho)
			}

		}
	})
}
