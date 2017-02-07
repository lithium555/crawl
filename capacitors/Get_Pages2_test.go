package capacitors

import (
	"log"
	"reflect"
	"testing"
)

func TestGetCapacitorsFamily(t *testing.T) {
	testSlice := []string{"http://www.digikey.com/product-search/en/capacitors/accessories/131156",
		"http://www.digikey.com/product-search/en/capacitors/aluminum-polymer-capacitors/132402",
		"http://www.digikey.com/product-search/en/capacitors/aluminum-capacitors/131081",
		"http://www.digikey.com/product-search/en/capacitors/capacitor-networks-arrays/131080",
		"http://www.digikey.com/product-search/en/capacitors/ceramic-capacitors/131083",
		"http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-edlc-supercapacitors/131084",
		"http://www.digikey.com/product-search/en/capacitors/film-capacitors/131088",
		"http://www.digikey.com/product-search/en/capacitors/mica-and-ptfe-capacitors/131309",
		"http://www.digikey.com/product-search/en/capacitors/niobium-oxide-capacitors/131747",
		"http://www.digikey.com/product-search/en/capacitors/silicon-capacitors/132347",
		"http://www.digikey.com/product-search/en/capacitors/tantalum-polymer-capacitors/132403",
		"http://www.digikey.com/product-search/en/capacitors/tantalum-capacitors/131082",
		"http://www.digikey.com/product-search/en/capacitors/thin-film-capacitors/131736",
		"http://www.digikey.com/product-search/en/capacitors/trimmers-variable-capacitors/131670",
	}

	rezSlice, err := GetCapacitorsFamily()
	if err != nil {
		t.Error(err)
	}
	log.Println(rezSlice)
	if !reflect.DeepEqual(testSlice, rezSlice) {
		t.Errorf("The slices are not equally: testSlice: '%v' \n rezSlice: '%v'\n", testSlice, rezSlice)
	}
}
