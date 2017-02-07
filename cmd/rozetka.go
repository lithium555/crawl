package main

import (
	"compress/gzip"
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	//"github.com/lithium555/crawl/rozetka"
	"github.com/lithium555/crawl/rozetka"
	"log"
	"os"
	"sync"
)

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go http.ListenAndServe(":6060", nil)
}

func main() {
	filename := flag.String("out", "new_file.nq.gz", "the name for the new File")
	flag.Parse()

	file, err := os.Create(*filename) //"testROZETKA2.nq.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	var wg sync.WaitGroup
	write := make(chan string) // t

	family, err := rozetka.FamiliesId()

	if err != nil {
		log.Printf("ERROR in rozetka.FamiliesId(): '%v'\n", err)
		return
	}

	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer log.Println("++++++++++++++++++routine exit")

			for it := range write {
				category, err := rozetka.ListProduct(it)
				if len(category) > 3 {
					category = category[:3]
				}
				if err != nil {
					log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
					continue
				}
				for _, v := range category {
					category2, err := rozetka.Categories(v)
					if len(category2) > 3 {
						category2 = category2[:3]
					}
					if err != nil {
						log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
						continue
					}
					for _, m := range category2 {
						for next, i := m, 0; next != "" && i < 3; i++ {
							getLink, nextPage, err := rozetka.GetLinkOnProduct(m)
							if err != nil {
								log.Printf("ERROR in rozetka.GetLinkOnProduct(): '%v'\n", err)
								continue
							}
							next = nextPage
							for _, r := range getLink {
								allcharacteristics, err := rozetka.AllCharacteristics(r)
								if err != nil {
									log.Printf("ERROR in rozetka.AllCharacteristics(): '%v'\n", err)
									log.Printf("r = '%v'\n", r)
									continue
								}
								mu.Lock()
								err3 := nquads.WriteObject(gzipWriter, allcharacteristics)
								mu.Unlock()
								if err3 != nil {
									log.Println(err3)
								}
							}
						}
					}
				}
			}
			log.Println("ROOTINE ENDED!!!!!!!!")
		}()
	}

	for _, v := range family {
		write <- v
	}
	close(write)

	wg.Wait()
	fmt.Println("We are done!")
}
