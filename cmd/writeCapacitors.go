package main

import (
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"github.com/lithium555/crawl/capacitors"
	"log"
	"os"
)

func main() {
	file, err := os.Create("testCap.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	families, err := capacitors.GetCapacitorsFamily()
	if err != nil {
		log.Printf("ERROR in capacitors.GetCapacitorsFamily(): '%v'\n", err)
	}
	for _, v := range families {
		for next := v; next != ""; {
			var list []string
			list, next, err = capacitors.List(next)
			if err != nil {
				log.Printf("ERROR in capacitors.List(): '%v'\n", err)
			}
			for _, z := range list {
				capInfo, err := capacitors.GetCapacitorsINFO(z)
				if err != nil {
					log.Printf("ERROR in capacitors.GetCapacitorsINFO(): '%v'\n", err)
				}
				err3 := nquads.WriteObject(file, capInfo)
				fmt.Println(err3)
			}
		}
	}
}
