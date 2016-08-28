package main

import (
	"bytes"
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"github.com/lithium555/crawl/intel"
	"log"
)

func main() {
	//	Here, _ := GetId("http://ark.intel.com/products/family/88392/6th-Generation-Intel-Core-i7-Processors#@Desktop")
	val, _ := intel.GetSpecification("94196") // /Intel-Core-i7-6900K-Processor-20M-Cache-up-to-3_70-GHz")
	fmt.Printf("VAL.Properties IS: '%q'\n", val.Properties)
	fmt.Printf("VAL.NAME '%q'\n", val.Name)
	fmt.Printf("VAL.url: '%q'\n", val.URL)

	//fmt.Println()
	//fmt.Println()
	//st, _ := intel.ListProducts("88392")
	//fmt.Printf("st: '%q'\n", st)
	//fmt.Println()
	//fmt.Println()
	//id, _ :=intel.GetFamiliesId()
	//fmt.Printf("id: '%q'\n", id)

	var b bytes.Buffer
	err := nquads.WriteObject(&b, val)
	log.Printf("b is == '%v'", b.String())
	fmt.Println(err)
}
