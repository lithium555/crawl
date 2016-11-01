package aliExpress

import (
	"reflect"
	"testing"
)

func TestGetComputerId(t *testing.T) {
	testSlice := []string{
		"https://ru.aliexpress.com/af/category/202002176.html?g=y&d=n&catName=tablet-pcs&origin=n&isViewCP=y&spm=2114.11020108.105.2.tEMEQT&pvId=19012-200658763",
		"https://ru.aliexpress.com/af/category/202004359.html?origin=n&g=y&isViewCP=y&spm=2114.11020108.105.6.tEMEQT&d=n&catName=tablet-accessories",
		"https://ru.aliexpress.com/af/category/202002183.html?origin=n&isViewCP=y&spm=2114.11020108.105.23.tEMEQT&d=n&catName=internal-solid-state-disks-ssd",
		"https://ru.aliexpress.com/af/category/202000058.html?origin=n&isViewCP=y&spm=2114.11020108.110.3.tEMEQT&d=n&catName=speakers",
		"https://ru.aliexpress.com/af/category/202000400.html?origin=n&isViewCP=y&spm=2114.11020108.110.4.tEMEQT&d=n&catName=memory-card",
		"https://ru.aliexpress.com/af/category/202001121.html?origin=n&isViewCP=y&spm=2114.11020108.110.5.tEMEQT&d=n&catName=digital-battery",
		"https://ru.aliexpress.com/af/category/202040673.html?origin=n&isViewCP=y&spm=2114.11020108.110.23.tEMEQT&d=n&catName=smart-watches",
		"https://ru.aliexpress.com/af/category/202040732.html?origin=n&isViewCP=y&spm=2114.11020108.110.22.tEMEQT&d=n&catName=smart-home",
		"https://ru.aliexpress.com/af/category/202005798.html?origin=n&isViewCP=y&spm=2114.11020108.110.25.tEMEQT&d=n&catName=psp",
		"https://ru.aliexpress.com/af/category/202059382.html?origin=n&isViewCP=y&spm=2114.11020108.105.16.tEMEQT&d=n&catName=sensors-alarms",
		"https://ru.aliexpress.com/af/category/202000321.html?origin=n&isViewCP=y&spm=2114.11020108.105.15.tEMEQT&d=n&catName=surveillance-products",
	}
	rez, err := GetComputerId()
	if err != nil {
		t.Error(err)
	}
	if len(rez) == 0 {
		t.Error("Func GetComputerId doesn`t work, the lel(rez)==0")
	}
	if !reflect.DeepEqual(testSlice, rez) {
		t.Errorf("The slices are not equally: testSlice: '%v' \n rezSlice: '%v'\n", testSlice, rez)
	}
	if !containsAll(testSlice, rez) {
		t.Error("Values doesn`t exist in rez!!!")
	}
}
func containsAll(expect, arr []string) bool {
	m := make(map[string]struct{}, len(expect)) // создал мапу динной слайса expect, который = testSlice
	for _, v := range expect {                  // проходим по всем знаечниям слайса expect
		m[v] = struct{}{} //we write all elements as keys into a map (ложим все элементы со слайса экспект как ключи значений в мапе m)
	} // alues does not matter, we just want to be able to find any element in it (as key)
	for _, v := range arr { // Далее мы перебираем все значения в основном массиве arr = rez
		if len(m) == 0 { //если длинна мапы равна нулю, выйти
			return true /* the loop will have an end anyway, because `arr` is limited
			but, why to continue loop if we already found everything? this condition
			terminates the loop as soon as it finds all element in `expect`*/
		}
		delete(m, v) // удалить с мапы m элемент v
	}
	return len(m) == 0
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
