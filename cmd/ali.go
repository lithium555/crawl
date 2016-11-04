package main

import (
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"github.com/lithium555/crawl/aliExpress"
	"log"
	"os"
)

const ElectronComponents = "https://ru.aliexpress.com/af/category/202000051.html"
const noteBooks = "https://ru.aliexpress.com/af/category/202000104.html"
const desktopComputers = "https://ru.aliexpress.com/af/category/202000103.html"

func main() {
	file, err := os.Create("ali.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sl := []string{ElectronComponents, noteBooks, desktopComputers}

	for _, v := range sl {
		for next := v; next != ""; {
			data, np, err := aliExpress.GetList(next)
			if err != nil {
				log.Printf("GetList() doesn`t work, err = '%v'\n", err)
			}
			for _, k := range data {
				obj, err := aliExpress.GetAliSpecification(k)
				if err != nil {
					log.Printf("GetAliSpecification() doesn`t work, err = '%v'", err)
				}
				err2 := nquads.WriteObject(file, obj)
				fmt.Println(err2)
			}
			next = np
		}
	}
}
