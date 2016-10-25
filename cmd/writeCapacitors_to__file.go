package main

import (
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"github.com/lithium555/crawl/capacitors"
	"log"
	"os"
)

func main() {
	newCapacitorsFile, err := os.Create("testCapacitors.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newCapacitorsFile.Close()

	rezSlice, err := capacitors.GetCapacitorsFamily()
	if err != nil {
		log.Printf("ERROR in capacitors.GetCapacitorsFamily(): '%v'\n", err)
	}
	//log.Printf("NEW rezSlice: '%v'\n", rezSlice)
	var capacitorsSlice []string
	log.Printf("capacitorsSlice BEFORE CYCLE = '%v'\n", capacitorsSlice)
	for _, v := range rezSlice {
		log.Printf("capacitorsSlice IN CYCLE 1 = '%v'\n", capacitorsSlice)
		for next := v; next != ""; {
			//log.Printf("Next link for file: '%v'\n", next)
			log.Printf("capacitorsSlice IN CYCLE 2  = '%v'\n", capacitorsSlice)
			capacitorsSlice, next, err = capacitors.List(next)
			log.Printf("capacitorsSlice IN CYCLE 2 after func List = '%v'\n", capacitorsSlice)
			if err != nil {
				log.Printf("ERROR in capacitors.List(): '%v'\n", err)
			}
			//log.Printf("Next capacitorsSlice is: '%v'\n", capacitorsSlice)
		}

		//SliceNextPage = append(SliceNextPage, nextPage)
		for _, t := range capacitorsSlice {
			log.Printf("capacitorsSlice IN CYCLE 3  = '%v'\n", capacitorsSlice)
			gcrawlCapacitors, err := capacitors.GetCapacitorsINFO(t)
			if err != nil {
				log.Printf("ERROR in capacitors.GetCapacitorsINFO(): '%v'\n", err)
			}
			err3 := nquads.WriteObject(newCapacitorsFile, gcrawlCapacitors)
			fmt.Println(err3)
		}
		break
	}

}
