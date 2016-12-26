package rozetka

import (
	"fmt"
	"github.com/dennwc/gcrawl"
	"github.com/stretchr/testify/require"
	"log"
	"reflect"
	"sort"
	"testing"
)

const link = "http://rozetka.com.ua/computers-notebooks/c80253/"
const link2 = "http://rozetka.com.ua/telefony-tv-i-ehlektronika/c4627949/"
const value = "http://rozetka.com.ua/notebooks/c80004/"
const testLink = "http://rozetka.com.ua/notebooks/c80004/filter/preset=workteaching/"
const notebook = "http://rozetka.com.ua/hewlett_packard_250_y8c06es/p12009754/"
const dataTable = "http://rozetka.com.ua/hewlett_packard_250_y8c06es/p12009754/characteristics/"

func TestRozetka(t *testing.T) {
	FamiliesId()
	ListProduct(link)
	//	ListProduct(link2)
	Categories(value)
	GetLinkOnProduct(testLink)
	AllCharacteristics(notebook)
	//getSpecification(dataTable)
}

func TestFamiliesId(t *testing.T) {

	slice := []string{
		"http://rozetka.com.ua/computers-notebooks/c80253/",
		"http://rozetka.com.ua/telefony-tv-i-ehlektronika/c4627949/",
		"http://bt.rozetka.com.ua/",
		"http://rozetka.com.ua/tovary-dlya-doma/c2394287/",
		"http://rozetka.com.ua/instrumenty-i-avtotovary/c4627858/",
		"http://rozetka.com.ua/santekhnika-i-remont/c4628418/",
		"http://rozetka.com.ua/dacha-sad-ogorod/c2394297/",
		"http://rozetka.com.ua/sport-i-uvlecheniya/c4627893/",
		"http://rozetka.com.ua/shoes_clothes/c1162030/",
		"http://rozetka.com.ua/krasota-i-zdorovje/c4629305/",
		"http://rozetka.com.ua/kids/c88468/",
		"http://rozetka.com.ua/office-school-books/c4625734/",
		"http://rozetka.com.ua/alkoholnie-napitki-i-produkty/c4626923/",
		"http://rozetka.com.ua/tovary-dlya-biznesa/c4627851/",
		"http://rozetka.com.ua/payments-transfers-travel/",
		//"http://rozetka.com.ua/freeshipping/#page_top",
		"http://rozetka.com.ua/gifts2017/",
		"http://rozetka.com.ua/all-categories-goods/",
		//	"fvelhlbhehrbe",
	}
	rez, err := FamiliesId()
	if err != nil {
		t.Fatal(err)
	}
	if len(rez) < len(slice) {
		t.Error("We don`t have all Families ID!!!")
	}

	//for _, v := range slice{
	//	for _, z := range rez{
	//		if z != v{
	//			t.Errorf("Missing value in the rezult slice = '%v'\n", z)
	//		}
	//	}
	//}
	if !reflect.DeepEqual(slice, rez) {
		t.Error("The slices are nt equal!!!")
		log.Printf("slice = '%v'\n", slice)
		log.Println()
		log.Printf("rez = '%v'\n", rez)
	}
}

const testLink1 = "http://rozetka.com.ua/computers-notebooks/c80253/"
const testLink2 = "http://rozetka.com.ua/telefony-tv-i-ehlektronika/c4627949/"

func TestListProduct(t *testing.T) {
	expectSlice2 := []string{
		"http://rozetka.com.ua/telefony/c4627900/",
		"http://rozetka.com.ua/tv/c80015/",
		"http://rozetka.com.ua/280596/c280596/",
		"http://rozetka.com.ua/portativnaya-ehlektronika/c4627865/",
		"http://rozetka.com.ua/2628092/c2628092/",
		"http://rozetka.com.ua/foto-i-video/c4628124/",
		"http://rozetka.com.ua/memory-cards/c80044/",
		"http://rozetka.com.ua/monopody-dlia-selfi-i-aksessuary/c4625067/",
		"http://rozetka.com.ua/headsets/c80032/",
		"http://rozetka.com.ua/mobile-cases/c146229/",
		"http://rozetka.com.ua/universalnye-mobilnye-batarei/c387969/",
		"http://rozetka.com.ua/uslugi/c153670/",
	}

	expectSlice1 := []string{
		"http://rozetka.com.ua/notebooks/c80004/",
		"http://rozetka.com.ua/tablets/c130309/",
		"http://rozetka.com.ua/computers-notebooks-accessories/c80256/",
		"http://rozetka.com.ua/aksessuary-dlya-planshetov/c108714/",
		"http://hard.rozetka.com.ua/",
		"http://hard.rozetka.com.ua/computers/c80095/",
		"https://soft.rozetka.com.ua/",
		"http://rozetka.com.ua/e-books/c80023/",
		"http://rozetka.com.ua/office-equipment/c80254/",
		"http://rozetka.com.ua/game-zone/c80261/",
		"http://rozetka.com.ua/network-equipment/c80111/",
	}

	rez1, err := ListProduct(testLink1)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("rez 1 = '%v'\n", rez1)

	rez2, err := ListProduct(testLink2)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("rez 2 = '%v'\n", rez2)

	//if !reflect.DeepEqual(expectSlice1, rez1){
	//	t.Error("The slices are not EQUAL !!!")
	//}
	//if !reflect.DeepEqual(expectSlice2, rez2){
	//	t.Error("The slices are not EQUAL !!!")
	//}
	require.Equal(t, expectSlice2, rez2, "THE SLICES SHOULD BE THE SAME")
	require.Equal(t, expectSlice1, rez1, "THE SLICES SHOULD BE THE SAME")
}

//expected: []string{"http://rozetka.com.ua/telefony/c4627900/",
//"http://rozetka.com.ua/tv/c80015/",
//"http://rozetka.com.ua/280596/c280596/",
//"http://rozetka.com.ua/portativnaya-ehlektronika/c4627865/",
//"http://rozetka.com.ua/2628092/c2628092/",
//"http://rozetka.com.ua/foto-i-video/c4628124/",
//"http://rozetka.com.ua/memory-cards/c80044/",
//"http://rozetka.com.ua/monopody-dlia-selfi-i-aksessuary/c4625067/",
//"http://rozetka.com.ua/headsets/c80032/",
//"http://rozetka.com.ua/mobile-cases/c146229/",
//"http://rozetka.com.ua/universalnye-mobilnye-batarei/c387969/",
//"http://rozetka.com.ua/uslugi/c153670/]"}
//
//
//received: []string{"http://rozetka.com.ua/notebooks/c80004/",
//"http://rozetka.com.ua/tablets/c130309/",
//"http://rozetka.com.ua/computers-notebooks-accessories/c80256/",
//"http://rozetka.com.ua/aksessuary-dlya-planshetov/c108714/",
//"http://hard.rozetka.com.ua/",
//"http://hard.rozetka.com.ua/computers/c80095/",
//"https://soft.rozetka.com.ua/",
//"http://rozetka.com.ua/e-books/c80023/",
//"http://rozetka.com.ua/office-equipment/c80254/",
//"http://rozetka.com.ua/game-zone/c80261/",
//"http://rozetka.com.ua/network-equipment/c80111/"}

const testValue = "http://rozetka.com.ua/notebooks/c80004/"

func TestCategories(t *testing.T) {
	slice, err := Categories(testValue)
	if err != nil {
		t.Error(err)
	}
	//for _, v := range slice{
	//	log.Printf(v)
	//}

	testSlice := []string{
		"http://rozetka.com.ua/notebooks/c80004/filter/preset=game/",
		"http://rozetka.com.ua/notebooks/c80004/filter/preset=entertainment/",
		"http://rozetka.com.ua/notebooks/c80004/filter/preset=tonkie-i-legkie/",
		"http://rozetka.com.ua/notebooks/c80004/filter/preset=transformery/",
		"http://rozetka.com.ua/notebooks/c80004/filter/preset=dlya-biznesa/",
		"http://rozetka.com.ua/notebooks/c80004/filter/preset=workteaching/",
		"http://rozetka.com.ua/notebooks/c80004/filter/preset=budget/",
		"http://rozetka.com.ua/notebooks/c80004/filter/preset=netbooks/",
	}
	//if !reflect.DeepEqual(testSlice, slice){
	//	t.Error("The slices are nt equal!!!")
	//}
	sort.Strings(testSlice)
	sort.Strings(slice)

	require.Equal(t, testSlice, slice, "THE SLICES SHOULD BE THE SAME")
}

const tesTLink = "http://rozetka.com.ua/notebooks/c80004/filter/preset=workteaching/"

func TestGetLinkOnProduct(t *testing.T) {
	getLink, _, err := GetLinkOnProduct(tesTLink)
	if err != nil {
		log.Printf("ERROR in rozetka.GetLinkOnProduct(): '%v'\n", err)
	}
	log.Printf("getLink = '%v'\n", getLink)

	expect := []string{
		"http://rozetka.com.ua/hp_probook_430_x0p48es/p11092932/",
		"http://rozetka.com.ua/acer_packard_bell_enlg81ba_p1b4_nx_c44eu_014/p11827428/",
		"http://rozetka.com.ua/asus_x441sa_wx021d/p12482472/",
		"http://rozetka.com.ua/acer_nx_mzseu_035/p11785260/",
		"http://rozetka.com.ua/acer_nx_mz8eu_062/p12200665/",
		"http://rozetka.com.ua/acer_nx_mz8eu_074/p12200658/",
		"http://rozetka.com.ua/hp_x0n95es/p11073269/",
		"http://rozetka.com.ua/acer_nx_g7beu_007/p7088369/",
		"http://rozetka.com.ua/lenovo_80t700dfra/p12164510/",
		"http://rozetka.com.ua/lenovo_80t60076ra/p12227916/",
		"http://rozetka.com.ua/dell_i35p45dil_d1/p10524231/", //
		"http://rozetka.com.ua/asus_x751sa_ty095d/p12408636/",
		"http://rozetka.com.ua/hp_x0p75es/p12305910/",
		"http://rozetka.com.ua/acer_aspire_es1_731_p84r_nx_mzseu_033/p11784959/", //
		"http://rozetka.com.ua/hewlett_packard_250_z2x73es/p12009775/",
		"http://rozetka.com.ua/hewlett_packard_255_y8c03es/p12010027/",
		"http://rozetka.com.ua/hewlett_packard_250_y8c06es/p12009754/",
		"http://rozetka.com.ua/lenovo_ideapad_110_15ibr_80t70034ra/p12128516/",
		"http://rozetka.com.ua/hp_255_w4m53ea/p11408618/",
		"http://rozetka.com.ua/lenovo_ideapad_110_14ibr_80t6006dra/p12129650/",
		"http://rozetka.com.ua/lenovo_80tj00fbra/p12228378/",
		"http://rozetka.com.ua/lenovo_80m300lxua/p10219248/",
		"http://rozetka.com.ua/lenovo_80tj005yra/p11256613/",
		"http://rozetka.com.ua/asus_x541sa_xo058d/p11681709/",
		"http://rozetka.com.ua/dell_i35p45dilelk/p10506654/", //
		"http://rozetka.com.ua/hewlett_packard_255_y8c04es/p12009922/",
		"http://rozetka.com.ua/asus_vivobook_max_x541sa_xo137d/p12259850/",
		"http://rozetka.com.ua/lenovo_110_17acl_80um002fra/p11781151/",
		"http://rozetka.com.ua/acer_nx_c3yeu_022/p12200602/",
		"http://rozetka.com.ua/hewlett_packard_250_w4n53ea/p12009852/",
		"http://rozetka.com.ua/dell_i35p25dil_f46/p8370524/",
		//"http://rozetka.com.ua/hewlett_packard_250_w4n53ea/p12009852/",
		"http://rozetka.com.ua/dell_inspiron_i57p45dil_46s/p6937193/", //
	}
	sort.Strings(expect)
	sort.Strings(getLink)
	require.Equal(t, expect, getLink, "THE SLICES SHOULD BE THE SAME")
}

const testLinK = "http://rozetka.com.ua/notebooks/c80004/filter/preset=workteaching/"
const lInk2 = "http://rozetka.com.ua/notebooks/c80004/filter/page=2;preset=workteaching/"
const LInk3 = "http://rozetka.com.ua/notebooks/c80004/filter/page=3;preset=workteaching/"

func TestGetSpecification(t *testing.T) {
	//for _, v := range getLink{
	//	allcharacteristics, err := AllCharacteristics(v)
	//	if err != nil{
	//		log.Printf("ERROR in rozetka.AllCharacteristics(): '%v'\n", err)
	//	}
	//	object, err := GetSpecification(allcharacteristics)
	//	if err != nil{
	//		log.Printf("ERROR in ozetka.GetSpecification(): '%v'\n", err)
	//	}
	//	log.Printf("OBJECT = '%v'\n", object)
	//}
	_, nextPage, err := GetLinkOnProduct(testLinK)
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, nextPage, lInk2, "THE LINKS SHOULD BE THE SAME")
	//log.Printf("nextPage = '%v'\n", nextPage)
	_, nextPage2, err2 := GetLinkOnProduct(lInk2)
	if err2 != nil {
		t.Error(err2)
	}
	require.Equal(t, nextPage2, LInk3, "THE LINKS SHOULD BE THE SAME")
	//log.Printf("nextPage2 = '%v'\n", nextPage2)

	_, nextPage3, err3 := GetLinkOnProduct(LInk3)
	if nextPage3 != "" {
		log.Printf("ERROR in nextPage3 is not empty, nextPage3 = '%v'\n", nextPage3)
	}
	if err3 != nil {
		t.Error(err3)
	}
	//log.Printf("nextPage3 = '%v'\n", nextPage3)
}

const notebooK = "http://rozetka.com.ua/hewlett_packard_250_y8c06es/p12009754/"
const noTeBOOK = "http://rozetka.com.ua/hewlett_packard_250_y8c06es/p12009754/characteristics/"

func TestAllCharacteristics(t *testing.T) {
	object, err := AllCharacteristics(notebooK)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Object.Name = '%v'\n", object.Name)
	//fmt.Printf("Object.Properties = '%v'\n", object.Properties)
	//fmt.Printf("Object.URL = '%v'\n", object.URL)
	require.Equal(t, expect, object, "Objects are not equal")
}

var expect = &gcrawl.Object{
	URL: noTeBOOK,
	Properties: []gcrawl.Property{
		{ID: "", Name: "Процессор", Value: gcrawl.String("Четырехъядерный Intel Pentium N3710 (1.6 - 2.56 ГГц)")},
		{ID: "", Name: "Экран", Value: gcrawl.String("15.6\" (1920x1080) Full HD")},
		{ID: "", Name: "Объем оперативной памяти", Value: gcrawl.String("8 ГБ")},
		{ID: "", Name: "Операционная система", Value: gcrawl.String("DOS")},
		{ID: "", Name: "Краткие характеристики", Value: gcrawl.String("Экран 15.6” (1920x1080) Full HD, глянцевый с антибликовым покрытием / Intel Pentium N3710 (1.6 - 2.56 ГГц) / RAM 8 ГБ / HDD 1 ТБ / Intel HD Graphics 405 / DVD Super Multi / LAN / Wi-Fi / Bluetooth / веб-камера / DOS / 1.96 кг / черный")},
		{ID: "", Name: "Объём накопителя", Value: gcrawl.String("1 ТБ")},
		{ID: "", Name: "Оптический привод", Value: gcrawl.String("DVD Super Multi")},
		{ID: "", Name: "Цвет", Value: gcrawl.String("Черный")},
		{ID: "", Name: "Вес", Value: gcrawl.String("1.96 кг")},
		{ID: "", Name: "Клавиатура", Value: gcrawl.String("Без подсветки")},
		{ID: "", Name: "Тип оперативной памяти", Value: gcrawl.String("DDR3L-1600 МГц")},
		{ID: "", Name: "Батарея", Value: gcrawl.String("Cъемная")},
		{ID: "", Name: "Украинская раскладка клавиатуры", Value: gcrawl.String("Без")},
		{ID: "", Name: "Количество слотов для оперативной памяти", Value: gcrawl.String("1")},
		{ID: "", Name: "Дополнительные возможности", Value: gcrawl.String("Веб-камера HP VGA\nАудиосистема HD с технологией DTS Studio Sound\nВстроенные стереодинамики\nПоддержка технологии HP подавления шума\nПолноразмерная клавиатура островного типа с цифровой клавишной панелью и черной текстурой")},
		{ID: "", Name: "Чипсет", Value: gcrawl.String("Intel SoC")},
		{ID: "", Name: "Графический адаптер", Value: gcrawl.String("Интегрированный, Intel HD Graphics 405")},
		{ID: "", Name: "Сетевые адаптеры", Value: gcrawl.String("Wi-Fi 802.11a/b/g/n/ac\nBluetooth 4.2\nWi-Di\nGigabit Ethernet")},
		{ID: "", Name: "Разъемы и порты ввода-вывода", Value: gcrawl.String("1 x USB 3.0 / 2 x USB 2.0 / VGA / HDMI / LAN (RJ-45) / комбинированный аудиоразъем для наушников/микрофона / кардридер")},
		{ID: "", Name: "Характеристики батареи", Value: gcrawl.String("3-элементный литий-ионный аккумулятор, 31 Вт*ч")},
		{ID: "", Name: "Габариты (Ш х Г х В)", Value: gcrawl.String("384.3 x 254.6 x 24.3  мм")},
		{ID: "", Name: "Комплект поставки", Value: gcrawl.String("Ноутбук\nБлок питания\nРуководство пользователя\nГарантийный талон")},
		{ID: "", Name: "Гарантия", Value: gcrawl.String("12 месяцев")},
		//{ID:"", Name:"", Value:gcrawl.String("") },
		//{ID:"", Name:"", Value:gcrawl.String("") },
	},
	Name: "Ноутбук HP 250 G5 (Y8C06ES) Black Суперцена!!!",
}

const notebooK2 = "http://rozetka.com.ua/hewlett_packard_250_y8c06es/p12009754/"

func TestAllCharacteristics2(t *testing.T) {
	object, err := AllCharacteristics(notebooK2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Object.Name = '%v'\n", object.Name)
	//fmt.Printf("Object.Properties = '%v'\n", object.Properties)
	//fmt.Printf("Object.URL = '%v'\n", object.URL)
}

//http://rozetka.com.ua/hewlett_packard_250_y8c06es/p12009754/
