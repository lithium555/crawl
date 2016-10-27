package capacitors

import (
	"log"
	"os"
	"testing"
)

func TestMedia2(t *testing.T) {
	file, err := os.Create("testMedia2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	GetcapFamily, err := GetCapacitorsFamily()
	if err != nil {
		t.Errorf("ERROR in capacitors.GetCapacitorsFamily(): '%v'\n", err)
	}
	for _, v := range GetcapFamily {
		for next := v; next != ""; {
			ListSlice, nextPage, err := List(next)
			if err != nil {
				t.Errorf("ERROR in capacitors.List(): '%v'\n", err)
			}
			for _, z := range ListSlice {
				_, err := GetCapacitorsINFO(z)
				if err != nil {
					t.Errorf("Error in the func capacitors.GetCapacitorsINFO, err = '%v'\n", err)
				}

			}
			next = nextPage
		}
	}
}
