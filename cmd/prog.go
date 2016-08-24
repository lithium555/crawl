package main

import "github.com/lithium555/crawl/intel"

func main() {
	//	Here, _ := GetId("http://ark.intel.com/products/family/88392/6th-Generation-Intel-Core-i7-Processors#@Desktop")
	intel.GetSpecification("94196") // /Intel-Core-i7-6900K-Processor-20M-Cache-up-to-3_70-GHz")
	intel.ListProducts("88392")
	intel.GetFamiliesId()
}
