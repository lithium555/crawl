package aliExpress

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetComputerId(t *testing.T) {
	testSlice := []string{
		"https://ru.aliexpress.com/category/202003357/audio.html",
		"https://ru.aliexpress.com/category/202003040/alarm-systems-security.html",
		"https://ru.aliexpress.com/category/200003200/digital-cables.html",
		"https://ru.aliexpress.com/category/63710/mp3-player.html",
		"https://ru.aliexpress.com/category/100005062/tablet-pcs.html",
		"https://ru.aliexpress.com/category/100005062/tablet-pcs.html",
		"https://ru.aliexpress.com/category/3008/sensors-alarms.html",
		"https://ru.aliexpress.com/category/2118/printers.html",
		"https://ru.aliexpress.com/category/3011/surveillance-products.html",
		"https://ru.aliexpress.com/category/202000014/security-protection.html",
	}
	rez, err := GetComputerId()
	if err != nil {
		t.Error(err)
	}
	if len(rez) == 0 {
		t.Error("Func GetComputerId doesn`t work, the lel(rez)==0")
	}
	//if !reflect.DeepEqual(testSlice, rez) {
	//	t.Errorf("The slices are not equally: testSlice: '%v' \n rezSlice: '%v'\n", testSlice, rez)
	//}
	//if !containsAll(testSlice, rez) {
	//	t.Error("Values doesn`t exist in rez!!!")
	//}
	errorArr := exist(testSlice, rez)
	for _, u := range errorArr {
		t.Errorf("errorArr = '%v'\n", u)
	}

	//sl, err := Get()
	//if err != nil {
	//	t.Errorf("Func Get() does not work, err = '%v'\n", err)
	//}
	//for _, e := range sl {
	//	log.Printf("e = '%v'\n", e)
	//}
	//val, err := GetList(noteBooks)
	//if err != nil {
	//	t.Errorf("GetList() doesn`t work, err = '%v'\n", err)
	//}
	//for _, b := range val {
	//	fmt.Printf("b = '%v'\n", b)
	//}
	qrawl, err := GetAliSpecification(OneNote)
	if err != nil {
		t.Errorf("func GetAliSpecification() doen`t work, er = '%v'\n", err)
	}
	fmt.Printf("qrawl = '%q'\n", qrawl)
}

const noteBooks = "https://ru.aliexpress.com/af/category/202000104.html"
const OneNote = "https://ru.aliexpress.com/item/13-3inch-newest-laptop-computer-aluminium-ultrabook-I3-5th-Gen-processor-4GB-128GB-SSD-backlit-keyboard/32672611082.html"

//func containsAll(expect, arr []string) bool {
//	m := make(map[string]struct{}, len(expect)) // создал мапу динной слайса expect, который = testSlice
//	for _, v := range expect {                  // проходим по всем знаечниям слайса expect
//		m[v] = struct{}{} //we write all elements as keys into a map (ложим все элементы со слайса экспект как ключи значений в мапе m)
//	} // alues does not matter, we just want to be able to find any element in it (as key)
//	for _, v := range arr { // Далее мы перебираем все значения в основном массиве arr = rez
//		if len(m) == 0 { //если длинна мапы равна нулю, выйти
//			return true /* the loop will have an end anyway, because `arr` is limited
//			but, why to continue loop if we already found everything? this condition
//			terminates the loop as soon as it finds all element in `expect`*/
//		}
//		delete(m, v) // удалить с мапы m элемент v
//	}
//	return len(m) == 0
//}
func exist(testSlice, rez []string) []string {
	var arr3 []string
	for i := 0; i < len(testSlice); i++ {
		isUnique := true
		for j := 0; j < len(rez); j++ {
			if strings.HasPrefix(rez[j], testSlice[i]) { // testSlice[i] == rez[j]
				//	if strings.HasPrefix(data[i], "//") {
				isUnique = false
				break /* брейк позволяет не сравнивать дальше если и так уже понятно что это не уникальный елемент*/
			}
		}
		if isUnique == true {
			arr3 = append(arr3, testSlice[i])
		}
	}
	return arr3
}

/*
 that is why I can`t understand, because as I want to see.. that whole elements
 from the slice `expect` should be in the map `m` and if I found this element in
 the arr, I can delete it from map, and if the map is empty - I wil return true

 we write all elements as keys into a map
 values does not matter, we just want to be able to find any element in it (as key)
 next we iterate over all values in main array
 in fact we may want to find the key in map first, so we can tell if it's there or not, and then delete it (mark as found)
 by why to search for the value if we can just remove it
 if it's there, we will mark it as found
 if it's not in the map we do nothing
 and at the end we check if map still contains elements

 this will mean that loop haven't removed it, meaning it's not in the main slice
 `1+2` is an operation which takes two ints and returns int (result)
 `1 == 2` takes two ints and returns a bool as a result
 it's like

 var b bool
 b = len(m) == 0
 return b

 or

 if len(m) == 0 {
    b = true
 }

 *about return true*
  at least for the case when all expected element are somewhere at the beginning... if at least one is at the end it will iterate over whole loop, which is fine
 regardless "algorithm complexity", you case in O(N^2), my is O(2N) or O(N)
*/
