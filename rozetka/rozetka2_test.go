package rozetka

import (
	"log"
	"testing"
)

func TestFamiliesId2(t *testing.T) {
	slice, err := ListProduct("http://rozetka.com.ua/all-categories-goods/")
	if err != nil {
		t.Error(err)
	}
	log.Printf("Slice = '%v'\n", slice)
}
