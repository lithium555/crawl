package capacitors

import (
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	// TEST for func LIST
	//==============================================================
	url2 := "http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-supercaps/131084"
	capSlice, next, err := List(url2)
	if err != nil {
		t.Fatalf("Err is: '%v''", err)
	}
	if next != "http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-supercaps/131084/page/2" {
		t.Fatalf("next is: '%v'", next)
	}
	if len(capSlice) <= 1 {
		t.Fatalf("The capacity of slice capSlice is: '%v'. The have in it: '%v'", len(capSlice), capSlice)
	}
	for _, value := range capSlice {
		if !strings.HasPrefix(value, "http://www.digikey.com/product-search/en/") {
			t.Fatalf("Link does not have a Prefix 'http://www.digikey.com' it cpnsists of: '%v'", value)
		}
	}
	//==============================================================
	//TEST for func GetCapacitorsINFO
	//=============================================================
	Slice, errINFO := GetCapacitorsINFO("http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H224U-HL/604-1020-2-ND/970181")
	if errINFO != nil {
		t.Fatalf("Error is: '%v'", errINFO)
	}
	if len(Slice) <= 1 {
		t.Fatalf("The capacity of Slice is: '%v', it have in it '%v' elements", len(Slice), cap(Slice))
	}
}
