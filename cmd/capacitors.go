package main

import (
	"amazing/Project_1/crawl/capacitors"
	"fmt"
)

func main() {
	//"http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-supercaps/131084"
	nextLink := "/electric-double-layer-capacitors-supercaps/131084"
	nextLink = "http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-supercaps/131084"
	//capacitors, nextPage, err := CapacitorsListProducts33(nextLink)
	//fmt.Printf("capacitors: '%v'\n", capacitors)
	//fmt.Printf("nextPage: '%v'\n", nextPage)
	//fmt.Printf("err: '%v'\n", err)
	url := nextLink
	//for url != "" {
	//	caps, next, err := capacitors.List(url)
	//	//strings.Split(next, "/product-search/en/capacitors")
	//	fmt.Printf("next: '%v'\n", next)
	//	//strings.TrimPrefix(next, "/product-search/en/capacitors")
	//	//	fmt.Printf("next222 '%v' \n", next )
	//	if err != nil {
	//		log.Printf("ERORR: '%v'\n", err)
	//	}
	//	log.Println(caps)
	//	url = next
	//	fmt.Printf("URL: '%v'\n", url)
	//	fmt.Printf("Capacitors slice: '%v'\n", caps)
	//}
	caps, next, _ := capacitors.List(url)
	fmt.Printf("next is: '%v' \n", next)
	fmt.Printf("caps is: '%v'\n", caps)

	fmt.Println()
	fmt.Println()
	Data, _ := capacitors.GetCapacitorsINFO("http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H224U-HL/604-1020-2-ND/970181")
	fmt.Printf("DATA is: '%q' \n", Data)

	SliceRez, _ := capacitors.GetCapacitorsFamily()
	fmt.Printf("GetCapacitorsFamily cosists of slice, named SliceRez: '%v'\n", SliceRez)
}
