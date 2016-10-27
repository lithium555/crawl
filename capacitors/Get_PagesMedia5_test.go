package capacitors

import (
	"fmt"
	"testing"
)

func TestMedia6(t *testing.T) {
	url := "http://www.digikey.com/en/help/contact-us"

	cap, np, err := List(url)
	fmt.Printf("NextPage = '%v'\n", np)
	if err == nil {
		t.Errorf("func List doesn`t work, err = '%v'\n", err)
	}
	if np != "" {
		t.Errorf("Np is = '%v' \n", np)
	}
	if len(cap) != 0 {
		t.Errorf("The length of cap = '%v'", len(cap))
	}
}
