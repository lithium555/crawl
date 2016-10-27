package capacitors

import (
	"fmt"
	"github.com/dennwc/gcrawl/nquads"
	"log"
	"os"
	"testing"
)

func TestMedia(t *testing.T) {
	file, err := os.Create("testMedia.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	testSlice := []string{
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-55/31762-55-ND/4466211",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-59/31762-59-ND/4466212",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-61/31762-61-ND/4466213",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-62/31762-62-ND/4466214",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-72/31762-72-ND/4466215",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-77/31762-77-ND/4466216",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-79/31762-79-ND/4466217",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-82/31762-82-ND/4466218",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-83/31762-83-ND/4466219",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-87/31762-87-ND/4466220",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-88/31762-88-ND/4466221",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-91/31762-91-ND/4466222",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-94/31762-94-ND/4466223",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-97/31762-97-ND/4466224",
		"http://www.digikey.com/product-detail/en/cornell-dubilier-electronics-cde/31762-99/31762-99-ND/4466225",
		"//media.digikey.com/pdf/PCNs/Cornell%20Dubilier/30434%2031762%20Br%20Series%204-11-2013.pdf",
		"http://media.digikey.com/pdf/PCNs/Cornell%20Dubilier/30434%2031762%20Br%20Series%204-11-2013.pdf",
		"http://www.cde.com/resources/catalogs/hardware.pdf",
		"http://www.cde.com/resources/catalogs/hardware.pdf",
		"//media.digikey.com/pdf/PCNs/Cornell%20Dubilier/30434%2031762%20Br%20Series%204-11-2013.pdf",
		"http://www.cde.com/resources/catalogs/hardware.pdf",
		"http://www.digikey.com/product-search/en/capacitors",
	}

	for _, a := range testSlice {
		object, err := GetCapacitorsINFO(a)
		if err != nil {
			log.Printf("Error in the func capacitors.GetCapacitorsINFO, err = '%v'\n", err)
		}
		err3 := nquads.WriteObject(file, object)
		fmt.Println(err3)
	}
}
