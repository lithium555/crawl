package parser

import (
	"regexp"
)

var reUnits = regexp.MustCompile(`^(\d+(?:\.\d+)?)\s([a-zA-Z]+)$`)//(`([0-9]+(?:\.[0-9]+)?)\s+([A-Za-z]+)`)//(`(\d+(?:,\d+))\s([a-zA-Z]+)`)(`\d+(,\d+)?\s[a-zA-Z]+`)
//`[0-9]*\.?[0-9]+\s+[A-Za-z]+`
//`\d*\.?\d+\s+[A-Za-z]*\D+[^_]+[^-]+`
//\d*\.?\d+\s+[0-9A-Za-z]+[^_]+
type caseStruct struct{
	re *regexp.Regexp
	check string
	expect []string
}
var cases = []caseStruct{
	{re:reUnits, check:"10aaa", expect:nil},
	{re:reUnits, check:"14 nm", expect:[]string{"14 nm","14","nm"}},
	{re:reUnits, check:"20 MB", expect:[]string{"20 MB","20","MB"}},
	{re:reUnits, check:"3.20 GHz", expect:[]string{"3.20 GHz","3.20","GHz"}},
	{re:reUnits, check:"1 nm", expect:[]string{"1 nm","1","nm"}},
	{re:reUnits, check:".1 nm", expect:nil},
	{re:reUnits, check:"10 1", expect:nil},
	{re:reUnits, check:"10 _", expect:nil},
	{re:reUnits, check:"1 __", expect:nil},
	{re:reUnits, check:"1 __-", expect:nil},
	{re:reUnits, check:"10 __0", expect:nil},
	{re:reUnits, check:"10 0__-", expect:nil},
	{re:reUnits, check:"10 20", expect:nil},
	{re:reUnits, check: "10 _0", expect:nil},
	{re:reUnits, check:"bla 10 nm bla", expect:nil},
	{re:reUnits, check:"3.3V", expect: []string{"3,3 V"}},
	{re:reUnits, check:"220mF", expect:[]string{"220 mF"}},
//	{re:reUnits, check:"1 nm", expect:[]string{"1 nm","11","nm"}},
}


