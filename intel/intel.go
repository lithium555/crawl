package intel

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Characteristic struct {
	Id    string
	Key   string
	Value string
}

func GetSpecification(id string) ([]Characteristic, error) {
	NewURL := "http://ark.intel.com/products/" + id
	doc, err := goquery.NewDocument(NewURL)
	if err != nil {
		return nil, err
	}
	var Slice []Characteristic
	doc.Find("table.specs.infoTable>tbody>tr[id]").Each(func(i int, s *goquery.Selection) {
		Key := strings.TrimSpace(s.Find(".lc").Text())
		Value := strings.TrimSpace(s.Find(".rc").Text())
		Link, ok := s.Attr("id")
		if !ok {
			return
		}
		Slice = append(Slice, Characteristic{
			Id:    strings.TrimSpace(Link),
			Key:   Key,
			Value: Value})
	})
	return Slice, nil
}

func ListProducts(id string) ([]string, error) {
	NEWurl := "http://ark.intel.com/products/family/" + id
	doc2, err := goquery.NewDocument(NEWurl)
	if err != nil {
		return nil, err
	}
	var Slice []string
	doc2.Find("table.infoTable>tbody>tr>td>a").Each(func(i int, s *goquery.Selection) {
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
	return Slice, nil
}

func GetFamiliesId() ([]string, error) {
	var Slice []string
	const url = "http://ark.intel.com/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	doc.Find("table.col2.infoTable.infoColumns>tbody>tr>td.columns-2>a").Each(func(i int, s *goquery.Selection) {
		Link, ok := s.Attr("href")
		if !ok {
			return
		}
		Result := strings.TrimSpace(Link)
		log.Printf("THE RESULT:   '%v'", Result)
		if !strings.Contains(Result, "/products/") {
			return
		}
		//href="/products/family/59133/2nd-Generation-Intel-Core-i3-Processors#@Desktop"
		if !strings.Contains(Result, "/products/family/") {
			//Result = strings.TrimPrefix(Link, "/products/series/")
			return
		}
		Result = strings.TrimPrefix(Link, "/products/family/")
		fmt.Printf("WITHOUT PREFIX:   '%v' \n", Result)
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
