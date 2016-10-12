package capacitors

import (
	"testing"
	"reflect"
	"github.com/dennwc/gcrawl"
)

func TestGetCapacitorsINFO(t *testing.T) {
	//==============================================================
	//TEST for func GetCapacitorsINFO
	//=============================================================
	Slice, errINFO := GetCapacitorsINFO("http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H224U-HL/604-1020-2-ND/970181")
	if errINFO != nil {
		t.Fatalf("Error is: '%v'", errINFO)
	}
	if len(Slice.Properties) <= 1 {
		t.Fatalf("The capacity of Slice is: '%v', it have in it '%v' elements", len(Slice.Properties), cap(Slice.Properties))
	}

	//fmt.Printf("GetCapacitorsINFO function gives: '%v'\n", Slice)

	got := Slice
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("unexpected match:\n%#v\n%#v", got, expect)
		for j:=0; j < len(got.Properties); j++{
			if !reflect.DeepEqual(expect.Properties[j], got.Properties[j]){
				t.Errorf("unexpected match:\n%#v\n%#v", expect.Properties[j], got.Properties[j])
			}
		}
	}
}

var expect = &gcrawl.Object{
	URL: "http://www.digikey.com/product-detail/en/elna-america/DSK-3R3H224U-HL/604-1020-2-ND/970181",
	Properties: []gcrawl.Property{
		{ID:"", Name:"Online Catalog", Value: gcrawl.URL("http://www.digikey.com/catalog/en/partgroup/dsk-series/4122?mpart=DSK-3R3H224U-HL&vendor=604")},
		{ID:"", Name:"Category", Value: gcrawl.URL("http://www.digikey.com/product-search/en/capacitors")},
		{ID:"", Name: "Family", Value: gcrawl.URL("http://www.digikey.com/product-search/en/capacitors/electric-double-layer-capacitors-edlc-supercapacitors/131084")},
		{ID:"", Name:"Manufacturer", Value: gcrawl.String("Elna America") },
		{ID:"", Name:"Series", Value: gcrawl.URL("http://www.digikey.com/product-search/en?FV=ffec07d8") },
		{ID:"", Name:"Packaging", Value: gcrawl.String("Tape & Reel (TR)") },
		{ID:"", Name:"Part Status", Value: gcrawl.String("Active") },
		{ID:"", Name:"Capacitance", Value: gcrawl.Unit{220, "mF"} },
		{ID:"", Name:"Tolerance", Value: gcrawl.String("-20%, +80%") },
		{ID:"", Name:"Voltage - Rated", Value: gcrawl.Unit{3.3, "V"} },
		{ID:"", Name:"ESR (Equivalent Series Resistance)", Value: gcrawl.Unit{200, "Ohm"} },
		{ID:"", Name:"Lifetime @ Temp.", Value: gcrawl.String("1000 Hrs @ 60°C") },
		{ID:"", Name:"Termination", Value: gcrawl.String("Surface Mount") },
		{ID:"", Name:"Mounting Type", Value: gcrawl.String("Surface Mount") },
		{ID:"", Name:"Package / Case", Value: gcrawl.String("Coin, Wide Terminals - Opposite Sides") },
		{ID:"", Name:"Lead Spacing", Value: gcrawl.Bool(false) },
		{ID:"", Name:"Size / Dimension", Value: gcrawl.String("0.268\" Dia (6.80mm)") },
		{ID:"", Name:"Height - Seated (Max)", Value: gcrawl.String("0.118\" (3.00mm)") },
		{ID:"", Name:"Operating Temperature", Value: gcrawl.String("-10°C ~ 60°C") },
		{ID:"", Name:"Standard Package", Value: gcrawl.String("1,500")},
		{ID:"", Name:"Other Names", Value: gcrawl.String("604-1020-2 \nDSK3R3H224UHL")},
	},
}
