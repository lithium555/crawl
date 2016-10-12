package capacitors

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dennwc/gcrawl"
	"github.com/lithium555/crawl/funcParseUnit"
	"log"
	"strings"
)

// "/product-search/en/capacitors"
const CapacitorsURL = "http://www.digikey.com"

func List(id string) (capacitors []string, nextPage string, _ error) {
	//	log.Printf("ID: '%v'", id)
	//	WorkingURL := CapacitorsURL +"/product-search/en/capacitors"+ id
	WorkingURL := id
	//	log.Printf("WorkingURL: '%v'\n", WorkingURL)
	doc, err := goquery.NewDocument(WorkingURL)
	if err != nil {
		return nil, "", err
	}
	var caps []string
	doc.Find("table tbody tr td.tr-dkPartNumber.nowrap-culture>a").Each(func(i int, s *goquery.Selection) {
		Link, ok := s.Attr("href")
		if !ok {
			return
		}
		Result := strings.TrimSpace(Link)
		Result = CapacitorsURL + Result
		caps = append(caps, Result)
	})
	//ТЕПЕРЬ НАДО СКАЧАТЬ В ПОЛЕ СТРУКТУРЫ ССЫЛКУ НА СЛЕДУЮЩУЮ СТРАНИЦУ
	//	log.Printf("WORKINGurl: '%v'\n", WorkingURL)
	var nP string
	doc.Find("a.Next").First().Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if !ok {
			return
		}
		nP = strings.TrimSpace(link)
	})
	capacitors = caps
	nextPage = CapacitorsURL + nP
	//fmt.Printf("NEXTPAGE: '%v'\n", nP)
	//log.Printf("capacitors: '%v'", capacitors)
	return capacitors, nextPage, nil
}

//func main(){
//	//"http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-supercaps/131084"
//	nextLink := "/electric-double-layer-capacitors-supercaps/131084"
//	nextLink = "http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-supercaps/131084"
//	//capacitors, nextPage, err := CapacitorsListProducts33(nextLink)
//	//fmt.Printf("capacitors: '%v'\n", capacitors)
//	//fmt.Printf("nextPage: '%v'\n", nextPage)
//	//fmt.Printf("err: '%v'\n", err)
//
//	url := nextLink
//	for url != "" {
//		caps, next, err := CapacitorsListProducts33(url)
//		//strings.Split(next, "/product-search/en/capacitors")
//		fmt.Printf("next: '%v'\n", next)
//		//strings.TrimPrefix(next, "/product-search/en/capacitors")
//		//	fmt.Printf("next222 '%v' \n", next )
//		if err != nil {
//			log.Printf("ERORR: '%v'\n", err)
//		}
//		log.Println(caps)
//		url = next
//		fmt.Printf("URL: '%v'\n", url)
//		fmt.Printf("Capacitors slice: '%v'\n", caps)
//	}
//}

//type Characteristic struct {
//	Id          string
//	Description string
//	Value       string
//}

// "http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H224U-HL/604-1020-2-ND/970181"
func GetCapacitorsINFO(pageUrl string) (*gcrawl.Object, error) {
	doc, err := goquery.NewDocument(pageUrl)
	if err != nil {
		return nil, err
	}
	var Slice []gcrawl.Property
	doc.Find("table.attributes-table-main>tbody>tr").Each(func(i int, s *goquery.Selection) {
		Left := strings.TrimSpace(s.Find("th").Text())
		Right := strings.TrimSpace(s.Find("td").Text())
		RightWithLink := s.Find("td a")
		link, _ := RightWithLink.Attr("href")
		if strings.HasPrefix(link, "javascript:") {
			link = ""
		}
		//log.Printf("link after 0: '%v'", link)
		Key := Left
		Value := Right
		var gcrawl_value gcrawl.Value
		//	log.Printf("LEFT IS: '%v'\n", Left)
		//log.Printf("RIGHT IS: '%v'\n", Right)
		if link != "" {
			if strings.HasSuffix(link, ".pdf") || strings.HasPrefix(link, CapacitorsURL) {
				fmt.Printf("link with suffix: '%v' \n", link)
				return
			}
			Value = CapacitorsURL + link
			fmt.Printf("Value with prefix CapacitorsURL: '%v'\n", Value)
		}
		_, ok := s.Attr("id")
		if ok {
			return
		}
		/////////////////////////////////////////////////////////////////////////////////////////////
		if Value == "-" {
			gcrawl_value = gcrawl.Bool(false)
		}else if unitValue, err := funcParseUnit.ParseUnits(Value); err == nil {
			gcrawl_value = *unitValue
			log.Printf("gcrawl_Value is '%v'\n", gcrawl_value)
		//}else if _, err := funcParseUnit.ParseUnits(Value); err != nil {
		//	log.Printf("FUNC ParseUnits FAILED, error is '%v'\n", err)
		}else if strings.HasPrefix(Value, CapacitorsURL){
			gcrawl_value = gcrawl.URL(Value)
		}else {
			gcrawl_value = gcrawl.String(Value)
		}
		//log.Printf("gcrawl_VALUE: '%v'\n", gcrawl_value)
		Slice = append(Slice, gcrawl.Property{
			ID:    "",
			Name:  Key,
			Value: gcrawl_value})
	})
	//var data gcrawl.Object
	//data.Properties = append(data.Properties, Slice...)
	//var data gcrawl.Object
	//data.Properties = Slice
	return &gcrawl.Object{
		URL:        pageUrl,
		Properties: Slice,
	}, nil
}

const CapacitorsURL2 = "http://www.digikey.com/product-search/en/capacitors"

func GetCapacitorsFamily() ([]string, error) {
	var Slice []string
	doc, err := goquery.NewDocument(CapacitorsURL2)
	if err != nil {
		return nil, err
	}
	doc.Find("ul.catfiltersub li>a").Each(func(i int, s *goquery.Selection) {
		Link, ok := s.Attr("href")
		if !ok {
			return
		}
		CapacitorsLINKS := strings.TrimSpace(Link)
		//log.Printf("CAPACITORS LINKS: '%v'", CapacitorsLINKS)
		CAPASITORS_ID := strings.TrimPrefix(CapacitorsLINKS, "/product-search/en/capacitors")
		//?????? ??? "/product-search/en/capacitors/niobium-oxide-capacitors/131747"
		// ?????? ? ??????? ?????? ?? "/niobium-oxide-capacitors/131747"
		CAPASITORS_ID = CapacitorsURL2 + CAPASITORS_ID
		Slice = append(Slice, CAPASITORS_ID)
	})
	//fmt.Printf("SLICE1: '%v' \n", Slice)
	return Slice, nil
}
