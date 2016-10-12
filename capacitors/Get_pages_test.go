package capacitors

import (
	"strings"
	"testing"
	"fmt"
	"reflect"
)

func TestList(t *testing.T) {
	// TEST for func LIST
	//==============================================================
	url2 := "http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-supercaps/131084"
	capSlice, next, err := List(url2)
	if err != nil {
		t.Fatalf("Err is: '%v''", err)
	}
	if next != "http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-edlc-supercapacitors/131084/page/2" {
		t.Fatalf("next is: '%v'", next)
	}
	if len(capSlice) <= 1 {
		t.Fatalf("The capacity of slice capSlice is: '%v'. The have in it: '%v'", len(capSlice), capSlice)
	}
	for _, value := range capSlice {
		if !strings.HasPrefix(value, "http://www.digikey.com/product-detail/en/") {
			t.Fatalf("Link does not have a Prefix 'http://www.digikey.com' it consists of: '%v'", value)
		}
	}

	fmt.Printf("capSlice: '%v' \n", capSlice)
	fmt.Printf("next '%v' \n", next)

	testSlice := []string{"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H204T614-H2L/604-1147-2-ND/2171198",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H204T614-H2L/604-1147-1-ND/2171202",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H204T614-H2L/604-1147-6-ND/2171204",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H703T414-HRL/604-1148-2-ND/2171199",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H703T414-HRL/604-1148-1-ND/2171203",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H703T414-HRL/604-1148-6-ND/2171205",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H703T414-HLL/604-1165-2-ND/2171245",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H703T414-HLL/604-1165-1-ND/3283567",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H703T414-HLL/604-1165-6-ND/3283568",
	"http://www.digikey.com/product-detail/en/nichicon/JUWT1105MCD/493-4330-ND/2538684",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H334T-HL/604-1160-2-ND/3283816",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H334T-HL/604-1160-1-ND/3280653",
	"http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H334T-HL/604-1160-6-ND/3283815",
	"http://www.digikey.com/product-detail/en/eaton/HV0810-2R7105-R/283-4198-ND/3878078",
	"http://www.digikey.com/product-detail/en/kemet/FC0V474ZFTBR24/399-10942-2-ND/4291008",
	"http://www.digikey.com/product-detail/en/kemet/FC0V474ZFTBR24/399-10942-1-ND/4506394",
	"http://www.digikey.com/product-detail/en/kemet/FC0V474ZFTBR24/399-10942-6-ND/4506484",
	"http://www.digikey.com/product-detail/en/elna-america/DCK-3R3E224U-E/604-1007-ND/970168",
	"http://www.digikey.com/product-detail/en/nichicon/JUWT1155MPD/493-4331-ND/2538685",
	"http://www.digikey.com/product-detail/en/elna-america/DCK-3R3E204T614-E/604-1078-ND/1658299",
	"http://www.digikey.com/product-detail/en/panasonic-electronic-components/EEC-S0HD224H/P10788-ND/285584",
	"http://www.digikey.com/product-detail/en/panasonic-electronic-components/EEC-S0HD224V/P10792-ND/285587",
	"http://www.digikey.com/product-detail/en/eaton/KR-5R5H104-R/283-2818-ND/1556246",
	"http://www.digikey.com/product-detail/en/panasonic-electronic-components/EEC-S0HD334H/P11064-ND/300480",
	"http://www.digikey.com/product-detail/en/panasonic-electronic-components/EEC-S0HD334V/P11065-ND/300481",
	}
	if !reflect.DeepEqual(testSlice, capSlice){
		t.Errorf("The slices are not equally: testSlice: '%v' \n capSlice: '%v'\n", testSlice, capSlice)
	}

}
