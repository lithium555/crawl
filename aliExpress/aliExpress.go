package aliExpress

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dennwc/gcrawl"
	"log"
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

// BASIC SITE:
//   https://ru.aliexpress.com/all-wholesale-products.html

func GetList(v string) ([]string, string, error) {
	var data []string
	doc, err := goquery.NewDocument(v)
	if err != nil {
		return nil, "", err
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
	var np string
	doc.Find("a.page-next").Each(func(i int, s *goquery.Selection) {
		/*
					Also to doc.Find(".ui-pagination-navi.util-left>a:last-child") this selector works great too
						:last-child - псевдокласс
			 			типа последний элемент которому подходит селектор
			 			в родительськом элементе
		*/
		link, ok := s.Attr("href")
		if !ok {
			return
		}
		nextPage := strings.TrimSpace(link)
		if strings.HasPrefix(nextPage, "//") {
			nextPage = "https:" + nextPage
		}
		np = nextPage
	})
	return data, np, nil
}

func GetAliSpecification(url string) (*gcrawl.Object, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	var slice []gcrawl.Property
	doc.Find(".ui-box-body .product-property-list.util-clearfix li").Each(func(i int, s *goquery.Selection) {
		left := strings.TrimSpace(s.Find(".propery-title").Text())
		right := strings.TrimSpace(s.Find(".propery-des").Text())
		left = strings.Trim(left, ":")
		//	log.Printf("LEFT = '%v'\n", left)
		//	log.Printf("RIGHT = '%v'\n", right)
		var data gcrawl.Value
		data = gcrawl.String(right)
		slice = append(slice, gcrawl.Property{
			ID:    "",
			Name:  left,
			Value: data,
		})
	})
	doc.Find(".detail-wrap").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".product-name").Text()
		log.Printf("title = '%v'\n", title)
		//		price := s.Find(".product-price").Text()
		var data3 gcrawl.Value
		data3 = gcrawl.String(title)
		slice = append(slice, gcrawl.Property{
			ID:    "",
			Name:  "title",
			Value: data3,
		})
	})
	doc.Find(".product-price").Each(func(i int, s *goquery.Selection) {
		/*>input#sku-price-store*/
		//s.Find(“.product-price [itemprop='priceCurrency']”).getAttribute(“content”)
		price := s.Find("#j-sku-price").Text()
		price = strings.Replace(price, ",", ".", -1)
		price = strings.Replace(price, " ", "", -1)
		/*
		 links for testing
		 https://ru.aliexpress.com/item/Free-shipping-original-motherboard-for-Gigabyte-GA-MA770T-UD3-DDR3-AM3-motherboards-all-solid/32709149208.html?spm=2114.30010608.3.1.6HLi8f&ws_ab_test=searchweb0_0,searchweb201602_2_10093_10091_10090_10088_10089,searchweb201603_1&btsid=3089ecbd-32e6-4879-8969-f65158b35ac9
		 https://ru.aliexpress.com/item/CHUWI-Vi8-Plus-8-Tablets-Windows10Intel-Cherry-Trail-Z8300-IPS-2GB-RAM-32GB-ROM-Quad-Core/32719330780.html?spm=2114.03010108.3.37.2XuhWJ&ws_ab_test=searchweb0_0,searchweb201602_2_10065_10068_10084_10083_10080_10082_10081_10060_10061_10062_10056_10055_10054_10059_10099_10078_10079_10093_10073_10103_10102_10096_10052_10050_10051,searchweb201603_2&btsid=99eac70b-7841-4441-8b90-ca3dea15d292
		*/
		//	log.Printf("price = '%v'\n", price)
		//	log.Printf("price2 = '%x' '%x'", price, " ")
		var data2 gcrawl.Value
		data2 = gcrawl.String(price)
		slice = append(slice, gcrawl.Property{
			ID:    "",
			Name:  "price",
			Value: data2,
		})
	})
	//doc.Find(".product-price").Each(func(i int, s *goquery.Selection){
	doc.Find(".p-symbol").Each(func(i int, s *goquery.Selection) {
		a, _ := s.Attr("content")
		fmt.Printf("a = '%v'\n", a)
		var currency gcrawl.Value
		currency = gcrawl.String(a)
		slice = append(slice, gcrawl.Property{
			ID:    "",
			Name:  "currency",
			Value: currency,
		})
	})
	return &gcrawl.Object{
		URL:        url,
		Properties: slice,
	}, nil
}
