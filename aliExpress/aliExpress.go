package aliExpress

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const baseURL = "https://www.aliexpress.com/?spm=2114.11020108.1000001.7.bA5PdZ"

func GetComputerId() ([]string, error) {
	var data []string
	doc, err := goquery.NewDocument(baseURL)
	if err != nil {
		return nil, err
	}
	doc.Find(".sub-cate-row .sub-cate-items dt a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if !ok {
			return
		}
		rez := strings.TrimSpace(link)
		data = append(data, rez)
	})
	doc.Find(".sub-cate-row .sub-cate-items dd a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if !ok {
			return
		}
		rez := strings.TrimSpace(link)
		data = append(data, rez)
	})
	for i := 0; i < len(data); i++ {
		if strings.HasPrefix(data[i], "//") {
			data[i] = "https:" + data[i]
		}
	}
	return data, nil
}

//func main() {
//	DT, err := GetComputerId()
//	if err != nil {
//		log.Printf("ERR = '%v'\n", err)
//	}
//	log.Printf("DT = '%v'\n", DT)
//}
const testURL = "https://ru.aliexpress.com/all-wholesale-products.html"

func Get() ([]string, error) {
	var data []string
	doc, err := goquery.NewDocument(testURL)
	if err != nil {
		return nil, err
	}
	doc.Find(".sub-item-cont.util-clearfix li a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if !ok {
			return
		}
		rez := strings.TrimSpace(link)
		data = append(data, rez)
	})
	for i := 0; i < len(data); i++ {
		if strings.HasPrefix(data[i], "//") {
			data[i] = "https:" + data[i]
		}
	}
	return data, nil
}

func GetList(v string) ([]string, error) {
	var data []string
	doc, err := goquery.NewDocument(v)
	if err != nil {
		return nil, err
	}
	doc.Find(".detail h3 a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if !ok {
			return
		}
		rez := strings.TrimSpace(link)
		data = append(data, rez)
	})
	for i := 0; i < len(data); i++ {
		if strings.HasPrefix(data[i], "//") {
			data[i] = "https:" + data[i]
		}
	}
	return data, nil
}
