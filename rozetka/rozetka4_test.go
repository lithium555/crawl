package rozetka

import (
	"compress/gzip"
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"log"
	"os"
	"sync"
	"testing"
)

import (
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go http.ListenAndServe(":6060", nil)
}

func TestAllCharacteristics5(t *testing.T) {
	file, err := os.Create("testROZETKA2.nq")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	var wg sync.WaitGroup

	write := make(chan string) // t
	wg.Add(10)

	family, err := FamiliesId()
	//for _, v := range family{
	//	fmt.Printf("FAMILYid = '%q'\n", v)
	//}
	if err != nil {
		log.Printf("ERROR in rozetka.FamiliesId(): '%v'\n", err)
	}

	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			defer log.Println("++++++++++++++++++routine exit")
			//  This CODE I can replace for
			//for {
			//	it , ok := <- write
			//	if !ok{
			//		break
			//	}
			// THAT CODE^
			for it := range write {
				fmt.Printf("it = '%v'\n", it)
				category, err := ListProduct(it)
				if len(category) > 3 {
					category = category[:3]
				}
				if err != nil {
					log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
				}
				for _, v := range category {
					//	fmt.Printf("v to write2 = '%v'\n", v)
					//write2 <- v
					//	log.Printf("DATA = '%v'\n", v)
					category2, err := Categories(v)
					if len(category2) > 3 {
						category2 = category2[:3]
					}
					if err != nil {
						log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
					}
					for _, m := range category2 {
						fmt.Printf("category2 m= '%v'\n", m)
						for next, i := m, 0; next != "" && i < 3; i++ {
							getLink, nextPage, err := GetLinkOnProduct(m)
							if err != nil {
								log.Printf("ERROR in rozetka.GetLinkOnProduct(): '%v'\n", err)
							}
							next = nextPage
							for _, r := range getLink {
								if r == "" {
									t.Errorf("r = '%v'\n", r)
								}
								allcharacteristics, err := AllCharacteristics(r)
								if err != nil {
									log.Printf("ERROR in rozetka.AllCharacteristics(): '%v'\n", err)
									log.Printf("r = '%v'\n", r)
								}
								//object, err := rozetka.GetSpecification(allcharacteristics)
								//if err != nil {
								//	log.Printf("ERROR in ozetka.GetSpecification(): '%v'\n", err)
								//}

								//if allcharacteristics == nil{
								//	//t.Errorf("allcharacteristics has nil pointer = '%v'", allcharacteristics)
								//	return
								//}
								if allcharacteristics == nil {
									break
								}
								mu.Lock()
								err3 := nquads.WriteObject(gzipWriter, allcharacteristics)
								mu.Unlock()
								log.Println(err3)
							}
						}
					}
				}
			}
			//time.Sleep(time.Second * 4)
			log.Println("ROOTINE ENDED!!!!!!!!")
		}()
	}

	for _, v := range family {
		//log.Printf("write (variable.val) = '%v'\n", v)
		write <- v
	}
	close(write)
	//for _, v := range family {
	//	list, err := rozetka.ListProduct(v)
	//	if err != nil {
	//		log.Printf("ERROR in rozetka.ListProduct(): '%v'\n", err)
	//	}
	//var data string
	//for i:= 0; i < 100; i++{
	//	data = <- write2
	//	log.Printf("DATA = '%v'\n", data)
	//}
	//list := data.val
	//for _, t := range list {
	//	time.Sleep(time.Second * 2)

	wg.Wait()
	fmt.Println("We are done!")
}
