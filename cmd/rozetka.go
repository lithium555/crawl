package main

import (
	"compress/gzip"
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"github.com/lithium555/crawl/rozetka"
	"log"
	"os"
	"sync"
)

func main() {
	file, err := os.Create("testROZETKA.nq")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	var wg sync.WaitGroup
	wg.Wait()
	write := make(chan dt)  // t
	write2 := make(chan dt) // t
	//channel2 := make(chan string)  // m
	//channel3 := make(chan string) // r

	family, err := rozetka.FamiliesId()
	if err != nil {
		log.Printf("ERROR in rozetka.FamiliesId(): '%v'\n", err)
	}
	go func() {
		defer wg.Done()
		//  This CODE I can replace for
		//for {
		//	it , ok := <- write
		//	if !ok{
		//		break
		//	}
		// THAT CODE^
		for it := range write {
			category, err := rozetka.ListProduct(it.val)
			if err != nil {
				log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
			}
			for _, v := range category {
				var n dt
				n.val = v
				write2 <- n
			}
		}
	}()
	var variable dt
	for _, v := range family {
		variable.val = v
		log.Printf("write (variable.val) = '%v'\n", variable.val)
		write <- variable
	}
	//for _, v := range family {
	//	list, err := rozetka.ListProduct(v)
	//	if err != nil {
	//		log.Printf("ERROR in rozetka.ListProduct(): '%v'\n", err)
	//	}
	var data dt
	for i := 0; i < 100; i++ {
		data = <-write2
		log.Printf("DATA = '%v'\n", data)
	}
	//list := data.val
	//for _, t := range list {
	category, err := rozetka.Categories(data.val)
	if err != nil {
		log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
	}
	for _, m := range category {
		for next, i := m, 0; next != "" && i < 3; i++ {
			getLink, nextPage, err := rozetka.GetLinkOnProduct(m)
			if err != nil {
				log.Printf("ERROR in rozetka.GetLinkOnProduct(): '%v'\n", err)
			}
			next = nextPage
			for _, r := range getLink {
				allcharacteristics, err := rozetka.AllCharacteristics(r)
				if err != nil {
					log.Printf("ERROR in rozetka.AllCharacteristics(): '%v'\n", err)
				}
				//object, err := rozetka.GetSpecification(allcharacteristics)
				//if err != nil {
				//	log.Printf("ERROR in ozetka.GetSpecification(): '%v'\n", err)
				//}

				err3 := nquads.WriteObject(gzipWriter, allcharacteristics)

				fmt.Println(err3)
			}
		}
	}

}

type dt struct {
	val string
}
