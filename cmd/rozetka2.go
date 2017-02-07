package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"github.com/lithium555/crawl/rozetka"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	filename := flag.String("out", "new_file.nq.gz", "the name for the new File")
	workers := flag.Int("w", 10, "number of workers")
	limit := flag.Int("limit", 0, "limit of values in slice")
	flag.Parse()

	file, err := os.Create(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	var wg sync.WaitGroup
	write := make(chan string) // t

	var mu sync.Mutex
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range write {
				allcharacteristics, err := rozetka.AllCharacteristics(job)
				if err != nil {
					log.Printf("ERROR in rozetka.AllCharacteristics(): '%v'\n", err)
					log.Printf("r = '%v'\n", job)
					continue
				}
				mu.Lock()
				err = nquads.WriteObject(gzipWriter, allcharacteristics)
				mu.Unlock()
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}
	start := time.Now()
	defer func() {
		fmt.Println("done in", time.Since(start))
	}()

	families, err := rozetka.FamiliesId()
	if err != nil {
		log.Printf("ERROR in rozetka.FamiliesId(): '%v'\n", err)
		return
	}

	for _, it := range families {
		categories, err := rozetka.ListProduct(it)
		if n := *limit; n > 0 && len(categories) > n {
			categories = categories[:n]
		}
		if err != nil {
			log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
			continue
		}
		for _, v := range categories {
			categories2, err := rozetka.Categories(v)
			if n := *limit; n > 0 && len(categories2) > n {
				categories2 = categories2[:n]
			}
			if err != nil {
				log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
				continue
			}
			for _, m := range categories2 {
				for next := m; next != ""; {
					var links []string
					links, next, err = rozetka.GetLinkOnProduct(next)
					if err != nil {
						log.Printf("ERROR in rozetka.GetLinkOnProduct(): '%v'\n", err)
						log.Println(err)
						continue
					}
					for _, url := range links {
						write <- url
					}
				}
			}
		}
	}
	close(write)
	wg.Wait()
	fmt.Println("We are done!")
}
