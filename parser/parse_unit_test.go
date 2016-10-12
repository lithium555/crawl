package parser

import (
	"fmt"
	"github.com/dennwc/gcrawl"
	"reflect"
	"regexp"
	"testing"
)

type caseStruct struct {
	re     *regexp.Regexp
	check  string
	expect *gcrawl.Unit
}

var cases = []caseStruct{
	//{re:reUnits, check:"10aaa", expect:nil},
	{re: reUnits, check: "14 nm", expect: &gcrawl.Unit{14, "nm"}},
	{re: reUnits, check: "20 MB", expect: &gcrawl.Unit{20, "MB"}},
	{re: reUnits, check: "3.20 GHz", expect: &gcrawl.Unit{3.20, "GHz"}},
	{re: reUnits, check: "1 nm", expect: &gcrawl.Unit{1, "nm"}},
	{re: reUnits, check: ".1 nm", expect: nil},
	{re: reUnits, check: "10 1", expect: nil},
	{re: reUnits, check: "10 _", expect: nil},
	{re: reUnits, check: "1 __", expect: nil},
	{re: reUnits, check: "1 __-", expect: nil},
	{re: reUnits, check: "10 __0", expect: nil},
	{re: reUnits, check: "10 0__-", expect: nil},
	{re: reUnits, check: "10 20", expect: nil},
	{re: reUnits, check: "10 _0", expect: nil},
	{re: reUnits, check: "bla 10 nm bla", expect: nil},
}

func TestParseUnits(t *testing.T) {
	for _, c := range cases {
		got, err := ParseUnits(c.check)
		/*
			dennwc [10:09 PM]
			смотри, тест должен быть зеленым (ни разу не вызвать Errf) если:
			а) есть экспект, нет ошибки и экспект совпадает (DeepEq)
			б) нет експекта, есть ошибка и экспект совпадает
			[10:10]
			и он должен быть красным, если:
			а) есть експект и есть ошибка либо експект не совпадает
			б) нет экспекта и нет ошибеи либо экспект не совпадает
					[10:11]
			хотя я обычно в таких случаях делаю иф/елс по первому условию и внутри еще по ифу на второе условие
		*/
		if c.expect != nil && err != nil && reflect.DeepEqual(got, c.expect) || c.expect == nil && err != nil && reflect.DeepEqual(got, c.expect) || c.expect != nil && err == nil {
			fmt.Printf("All if GREAT, the value is:  '%v' \n", got)
		}
		//if c.expect == nil && err != nil && reflect.DeepEqual(got, c.expect){
		//	fmt.Printf("All if GREAT, the value is:  '%v' \n", got)
		//}
		//if c.expect != nil && err == nil {
		//	fmt.Printf("All is fine, the value is: '%v' \n", got)
		//}
		if (c.expect != nil && err != nil || !reflect.DeepEqual(got, c.expect)) || (c.expect == nil && err == nil || !reflect.DeepEqual(got, c.expect)) {
			t.Errorf("unexpected match: %v vs %v \n", got, c.expect)
			t.Errorf("The string, whitch broken: '%q' \n", c.check)
		}
		//if c.expect == nil && err == nil || !reflect.DeepEqual(got, c.expect){
		//	t.Errorf("unexpected match: %v vs %v \n", got, c.expect)
		//	t.Errorf("The string, whitch broken: '%q' \n", c.check)
		//}
	}
}
