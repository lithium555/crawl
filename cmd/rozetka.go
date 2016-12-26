package main

import (
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"github.com/lithium555/crawl/rozetka"
	"log"
	"os"
)

func main() {
	file, err := os.Create("testROZETKA.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	family, err := rozetka.FamiliesId()
	if err != nil {
		log.Printf("ERROR in rozetka.FamiliesId(): '%v'\n", err)
	}
	for _, v := range family {
		list, err := rozetka.ListProduct(v)
		if err != nil {
			log.Printf("ERROR in rozetka.ListProduct(): '%v'\n", err)
		}
		for _, t := range list {
			category, err := rozetka.Categories(list)
			if err != nil {
				log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
			}
			for _, m := range category {
				getLink, _, err := rozetka.GetLinkOnProduct(m)
				if err != nil {
					log.Printf("ERROR in rozetka.GetLinkOnProduct(): '%v'\n", err)
				}
				allcharacteristics, err := rozetka.AllCharacteristics(getLink)
				if err != nil {
					log.Printf("ERROR in rozetka.AllCharacteristics(): '%v'\n", err)
				}
				object, err := rozetka.GetSpecification(allcharacteristics)
				if err != nil {
					log.Printf("ERROR in ozetka.GetSpecification(): '%v'\n", err)
				}
				err3 := nquads.WriteObject(file, object)
				fmt.Println(err3)
			}
		}
	}
}
