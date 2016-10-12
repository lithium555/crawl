package intel

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/dennwc/gcrawl"
	"github.com/lithium555/crawl/parser"
	"log"
	"strconv"
	"strings"
)

//type Characteristic struct {
//	Id    string
//	Key   string
//	Value string
//}

func GetSpecification(id string) (*gcrawl.Object /*[]Characteristic*/, error) {
	NewURL := baseUrl + "/products/" + id
	doc, err := goquery.NewDocument(NewURL)
	if err != nil {
		return nil, err
	}
	var Slice []gcrawl.Property //Characteristic
	doc.Find("table.specs.infoTable>tbody>tr[id]").Each(func(i int, s *goquery.Selection) {
		key := strings.TrimSpace(s.Find(".lc").Text())
		key = strings.Trim(key, " ‡")
		value := strings.TrimSpace(s.Find(".rc").Text())
		link, ok := s.Attr("id")
		if !ok {
			return
		}
		//valWithLink := s.Find("td>a")
		//l, _ := valWithLink.Attr("href")
		//if l == ""{
		//	value = value
		//}else{
		//	value = l
		//}
		//Закоммиченый код можно написать в двух вариантах, вариант первый:
		//valWithLink := s.Find("td>a")
		//if l, _ := valWithLink.Attr("href"); l != ""{
		//	value = l
		//}
		// Вариант второй:
		var data gcrawl.Value
		if l, _ := s.Find("td>a").Attr("href"); l != "" {
			value = l
			data = gcrawl.URL(value)
		} else if value == "Yes" {
			data = gcrawl.Bool(true)
		} else if value == "No" {
			data = gcrawl.Bool(false)
		} else if intValue, err := strconv.ParseInt(value, 0, 64); err == nil {
			data = gcrawl.Int(intValue)
		} else if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			data = gcrawl.Float(floatValue)
		} else if unitValue, err := parser.ParseUnits(value); err == nil {
			data = *unitValue
			/*
				 type Unit struct {
						Value float64
						Unit  string
					}
			*/
		} else {
			data = gcrawl.String(value)
		}
		Slice = append(Slice, gcrawl.Property{
			ID:    strings.TrimSpace(link),
			Name:  key,
			Value: data,
		})
	})
	log.Println(Slice)
	//var data gcrawl.Object
	//data.Properties = append(data.Properties, Slice...)
	return &gcrawl.Object{
		URL:        NewURL,
		Properties: Slice,
	}, nil
}

func ListProducts(id string) ([]string, error) {
	NEWurl := baseUrl + "/products/family/" + id
	doc, err := goquery.NewDocument(NEWurl)
	if err != nil {
		return nil, err
	}
	var Slice []string
	doc.Find("table.infoTable>tbody>tr>td>a").Each(func(i int, s *goquery.Selection) {
		Link, ok := s.Attr("href")
		if !ok {
			return
		}
		REZ := strings.TrimSpace(Link)
		if !strings.Contains(REZ, "/products/") {
			return
		}
		REZ = strings.TrimPrefix(Link, "/products/")
		Index := strings.Index(REZ, "/")
		if Index < 0 {
			return
		}
		Slice = append(Slice, REZ[:Index])
	})
	//log.Println(Slice)
	return Slice, nil
}

const baseUrl = "http://ark.intel.com"

func GetFamiliesId() ([]string, error) {
	var Slice []string
	doc, err := goquery.NewDocument(baseUrl)
	if err != nil {
		return nil, err
	}
	doc.Find("table.col2.infoTable.infoColumns>tbody>tr>td.columns-2>a").Each(func(i int, s *goquery.Selection) {
		Link, ok := s.Attr("href")
		if !ok {
			return
		}
		Result := strings.TrimSpace(Link)
		//log.Printf("THE RESULT:   '%v'", Result)
		if !strings.Contains(Result, "/products/") {
			return
		}
		//href="/products/family/59133/2nd-Generation-Intel-Core-i3-Processors#@Desktop"
		if !strings.Contains(Result, "/products/family/") {
			//Result = strings.TrimPrefix(Link, "/products/series/")
			return
		}
		Result = strings.TrimPrefix(Link, "/products/family/")
		//fmt.Printf("WITHOUT PREFIX:   '%v' \n", Result)
		//href="59133/2nd-Generation-Intel-Core-i3-Processors#@Desktop"
		Index := strings.Index(Result, "/")
		if Index < 0 {
			return
		}
		Slice = append(Slice, Result[:Index])
		//log.Printf("SLICE: '%v'", Slice)
	})
	return Slice, nil
}
