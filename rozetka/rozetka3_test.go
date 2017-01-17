package rozetka

import (
	"compress/gzip"
	"github.com/dennwc/gcrawl/nquads"
	"log"
	"os"
	"sync"
	"testing"
)

func TestAllCharacteristics3(t *testing.T) {
	file, err := os.Create("testROZETKA_2_0.nq")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()
	//data := []string{
	//	"http://rozetka.com.ua/computers-notebooks/c80253/",
	//	"http://rozetka.com.ua/telefony-tv-i-ehlektronika/c4627949/",
	//	"http://bt.rozetka.com.ua/",
	//	"http://rozetka.com.ua/tovary-dlya-doma/c2394287/",
	//	"http://rozetka.com.ua/instrumenty-i-avtotovary/c4627858/",
	//	"http://rozetka.com.ua/santekhnika-i-remont/c4628418/",
	//	"http://rozetka.com.ua/dacha-sad-ogorod/c2394297/",
	//	"http://rozetka.com.ua/sport-i-uvlecheniya/c4627893/",
	//	"http://rozetka.com.ua/shoes_clothes/c1162030/",
	//	"http://rozetka.com.ua/krasota-i-zdorovje/c4629305/",
	//	"http://rozetka.com.ua/kids/c88468/",
	//	"http://rozetka.com.ua/office-school-books/c4625734/",
	//	"http://rozetka.com.ua/alkoholnie-napitki-i-produkty/c4626923/",
	//	"http://rozetka.com.ua/tovary-dlya-biznesa/c4627851/",
	//	"http://rozetka.com.ua/payments-transfers-travel/",
	//	//"http://rozetka.com.ua/freeshipping/#page_top",
	//	"http://rozetka.com.ua/gifts2017/",
	//	"http://rozetka.com.ua/all-categories-goods/",
	//	//	"fvelhlbhehrbe",
	//}
	var wg sync.WaitGroup
	//var SL []string
	//for _, v := range data{
	//	string,  _:= Categories(v)
	//	for _, m := range string{
	//		SL = append(SL, m)
	//	}
	//}
	//fmt.Printf("SL = '%v'\n", SL)
	//-----------------------------------------------------------------------------------------------
	//	write := make(chan string)
	//	for _ , n := range data{
	//		write <- n
	//	}
	//
	//	go func(){
	//		defer wg.Done()
	//		value := <-write
	//		log.Println(value)
	//		 //for _, v := range value{
	//			// Categories(v)
	//		 //}
	//	}()
	wg.Wait()

	write2 := make(chan dt)
	write3 := make(chan dt)
	write4 := make(chan dt)
	write5 := make(chan dt)
	family, err := FamiliesId()
	if err != nil {
		log.Printf("ERROR in rozetka.FamiliesId(): '%v'\n", err)
	}

	var jz dt

	go func() {
		defer wg.Done()
		info := <-write2
		list, err := ListProduct(info.val)
		if err != nil {
			log.Printf("ERROR in rozetka.ListProduct(): '%v'\n", err)
		}
		//var f DT
		//f.val = list
		for _, v := range list {
			var m dt
			m.val = v
			write3 <- m
		}
	}()

	for _, v := range family {
		jz.val = v
		write2 <- jz
	}

	go func() {
		defer wg.Done()
		in := <-write3
		category, err := Categories(in.val)
		if err != nil {
			log.Printf("ERROR in rozetka.Categories(): '%v'\n", err)
		}
		for _, v := range category {
			var n dt
			n.val = v
			write4 <- n
		}
		//	}
	}()
	go func() {
		defer wg.Done()
		m := <-write4
		for next, i := m.val, 0; next != "" && i < 3; i++ {
			getLink, nextPage, err := GetLinkOnProduct(m.val)
			if err != nil {
				log.Printf("ERROR in rozetka.GetLinkOnProduct(): '%v'\n", err)
			}
			next = nextPage
			for _, v := range getLink {
				var n dt
				n.val = v
				write5 <- n
			}
		}
	}()
	go func() {

		rez := <-write5
		allcharacteristics, err := AllCharacteristics(rez.val)
		if err != nil {
			log.Printf("ERROR in rozetka.AllCharacteristics(): '%v'\n", err)
		}
		err3 := nquads.WriteObject(gzipWriter, allcharacteristics)
		log.Println(err3)
	}()
}

type DT struct {
	val []string
}

type dt struct {
	val string
}
