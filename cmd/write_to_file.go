package main

import (
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"github.com/lithium555/crawl/intel"
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//gcrawlObject, _ := intel.GetSpecification("94196")
	//err2 :=  nquads.WriteObject(newFile, gcrawlObject)
	//fmt.Println(err2)

	//familiesSlice, err := intel.GetFamiliesId()
	//if err != nil{
	//	log.Printf("ERROR in intel.GetFamiliesId: '%v'\n", err)
	//}
	//for i := 0; i < 5; i++ {
	//	listproductsSlice, err := intel.ListProducts(familiesSlice[i])
	//	if err != nil{
	//		log.Printf("ERROR in intel.ListProducts: '%v'\n", err)
	//		break
	//	}
	//	for j:=0; j< 5; j++{
	//		gcrawlObj, err := intel.GetSpecification(listproductsSlice[j])
	//		if err != nil{
	//			log.Printf("ERROR in intel.GetSpecification: '%v'\n", err)
	//			break
	//		}
	//		err2 :=  nquads.WriteObject(newFile, gcrawlObj)
	//		fmt.Println(err2)
	//	}
	//}
	families, err := intel.GetFamiliesId()
	if err != nil {
		log.Printf("ERROR in intel.GetFamiliesId: '%v'\n", err)
	}
	for _, v := range families {
		listproducts, err := intel.ListProducts(v)
		if err != nil {
			log.Printf("ERROR in intel.ListProducts: '%v'\n", err)
		}
		for _, t := range listproducts {
			gcrawlObj, err := intel.GetSpecification(t)
			if err != nil {
				log.Printf("ERROR in intel.GetSpecification: '%v'\n", err)
			}
			err3 := nquads.WriteObject(newFile, gcrawlObj)
			fmt.Println(err3)
		}
	}

}
