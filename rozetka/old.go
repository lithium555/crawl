package rozetka

//
//import (
//	"github.com/PuerkitoBio/goquery"
//	"fmt"
//	"strings"
//	"github.com/dennwc/gcrawl"
//)
//
//const baseURL = "http://rozetka.com.ua/"
//
//func FamiliesId()([]string, error){
//	var slice []string
//	doc, err := goquery.NewDocument(baseURL)
//	if err != nil{
//		return nil, err
//	}
//	//.f-menu-l-i-link.f-menu-l-i-link-arrow.sprite-side.novisited
//	doc.Find(".f-menu-l-i a").Each(func (i int, s *goquery.Selection){
//		link, ok :=  s.Attr("href")
//		if !ok{
//			return
//		}
//		fmt.Printf("link = '%v'\n", link)
//		slice = append(slice, link)
//	})
//	//log.Printf("slice = '%v'\n", slice)
//	return slice, nil
//}
//
//func ListProduct(id string)([]string, error){
//	var slice []string
//	doc, err := goquery.NewDocument(id)
//	if err != nil{
//		return nil, err
//	}
//	doc.Find(".pab-h3-link").Each(func(i int, s *goquery.Selection){
//		link, ok := s.Attr("href")
//		if !ok{
//			return
//		}
//		//	fmt.Printf("Here is link = '%v'\n", link)
//		slice = append(slice, link)
//	})
//	//	log.Printf("LINKSLICE = '%v'\n", slice)
//	return slice, nil
//}
//
//func Categories(url string)([]string, error){  // выбрать категорию, нуотбуки. нетбуки и прочее
//	var slice []string
//	doc, err := goquery.NewDocument(url)
//	if err != nil{
//		return nil, err
//	}
//	doc.Find(".pab-h4-link").Each(func(i int, s *goquery.Selection){
//		rez, ok := s.Attr("href")
//		if !ok{
//			return
//		}
//		//log.Printf("rez = '%v'\n", rez)
//		slice = append(slice, rez)
//	})
//	return slice, nil // в итог еуже страница с ноутбуками
//}
//
//func GetLinkOnProduct(url string)([]string, string, error){
//	var slice []string
//	doc, err := goquery.NewDocument(url)
//	if err != nil{
//		return nil, "", err
//	}
//	doc.Find(".g-i-tile-i-title.clearfix a").Each(func(i int, s *goquery.Selection){
//		link, ok := s.Attr("href")
//		if !ok{
//			return
//		}
//		//	log.Printf("HERE = '%v'\n", link)
//		slice = append(slice, link)
//	})
//	var nextPage string
//	/*.novisited*/
//	//doc.Find(".novisited.paginator-catalog-l-link").Each(func(i int, s *goquery.Selection){
//	doc.Find("a.novisited.paginator-catalog-l-link").First().Each(func(i int, s *goquery.Selection){
//		link, ok := s.Attr("href")
//		if !ok{
//			return
//		}
//		nextPage = link//strings.TrimSpace(link)
//		//log.Printf("nextPage = '%v'\n", nextPage)
//	})
//	return slice, nextPage, nil// получил слайсов ссылок на ноутбуки
//}
//
//func AllCharacteristics(url string)(string, error){
//	var newUrl string
//	doc, err := goquery.NewDocument(url)
//	if err != nil{
//		return "", err
//	}
//	//[name=characteristics]
//	doc.Find(".m-tabs-i[name=characteristics] a").Each(func(i int, s *goquery.Selection){
//		//ret, _ := s.Html()
//		//log.Printf("AllCharacteristics = '%v'\n", ret)
//		link, ok := s.Attr("href")
//		if !ok{
//			return
//		}
//		newUrl = link
//		//log.Printf("AllCharacteristics = '%v'\n", newUrl)
//	})
//
//	return newUrl, nil
//}
//
//func GetSpecification(url string)(*gcrawl.Object , error){
//	doc, err := goquery.NewDocument(url)
//	if err != nil{
//		return nil, err
//	}
//	var slice []gcrawl.Property
//	//.pp-characteristics-tab-l
//	//doc.Find("#pp-characteristics-tab-i.pp-characteristics-tab-i").Each(func(i int, s *goquery.Selection){
//	//	key := strings.TrimSpace(s.Find(".glossary-term").Text())
//	//	//value := strings.TrimSpace(s.Find(".pp-characteristics-tab-i-field").Text())
//	//	log.Printf("key = '%v'\n", key)
//	//	//log.Printf("value = '%v'\n", value)
//	//})
//	//var slice []gcrawl.Property
//	doc.Find("#pp-characteristics-tab-i.pp-characteristics-tab-i").Each(func(i int, s *goquery.Selection) {
//		var key string
//		key = strings.TrimSpace(s.Find(".glossary-term").Text())
//		if key == ""{
//			key = s.Find(".pp-characteristics-tab-i-title").Text()
//		}else{
//			key = strings.TrimSpace(s.Find(".glossary-term").Text())
//		}
//		//		value := strings.TrimSpace(s.Find(".pp-characteristics-tab-i-field").Text())
//		//т.е. если нету s.Find(".glossary-term").Text() то нужно искать s.Find(".pp-characteristics-tab-i-title").Text()
//
//		key = strings.Trim(key, " ")
//		key = strings.Trim(key, "\x09 \n\r")
//		//0a это сброс строки
//		//09 - вот это странные штуки
//		//		log.Printf("key = '%v'\n", key)
//		value := strings.TrimSpace(s.Find(".pp-characteristics-tab-i-field").Text())
//		value = strings.TrimSpace(value)
//		//	log.Printf("value = '%v'\n", value)
//		var data gcrawl.Value
//		data = gcrawl.String(value)
//		slice = append(slice, gcrawl.Property{
//			ID: "",
//			Name: key,
//			Value: data,
//		})
//		//	log.Printf("slice = '%q'\n", slice)
//	})
//	return &gcrawl.Object{
//		URL: url,
//		Properties: slice,
//	}, nil
//}
