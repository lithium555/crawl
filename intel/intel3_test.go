package intel

import (
	"testing"
	"fmt"
	"reflect"
)

func TestListProducts(t *testing.T){
	slice := []string{"93339", "88195", "88200", "88196", "93336", "88972", "93340", "88970", "88969", "93341", "88967", "91169", "91497", "88192", "91167", "91163", "88194", "90615", "90426", "88201",
	"93336", "88972", "93340", "88970", "88969", "93341", "88967", "91169", "91497", "88192", "91167", "91163", "88194",
	"93339", "88195", "88200", "88196",
	"88196", "88192", "90615", "90426", "88201"}
	rez, err := ListProducts("88392")
	if err!= nil{
			t.Fatal(err)
	}
	//Test already done for FAMILY: http://ark.intel.com/products/94196
	fmt.Println()
	fmt.Println()
	fmt.Println(reflect.DeepEqual(slice, rez))
	if !reflect.DeepEqual(slice, rez){
		t.Error("The slices are nt equal!!!")
	}
}