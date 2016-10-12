package funcParseUnit
//package main
//
//import (
//	"regexp"
//	"fmt"
//	//	"github.com/dennwc/gcrawl"
//	"github.com/dennwc/gcrawl"
//	"log"
//	"strconv"
//)
//
//func main(){
//	value, _ := parseUnits("10 nm")
//	fmt.Printf("The RESULT is: '%v'\n", value)
//}
//
//func parseUnits(s string) (*gcrawl.Unit, error){
//	Regular, REGerror := regexp.Compile(`^(\d+(?:\.\d+)?)\s([a-zA-Z]+)$`)
//	fmt.Printf("REGerror is '%v'\n", REGerror)
//	fmt.Printf("Regular is: '%v'\n", Regular)
//	SubmatchSlice := Regular.FindStringSubmatch(s)
//	if len(SubmatchSlice) == 0{
//		return nil, fmt.Errorf("The value is not found, length of slice: '%v' \n", len(SubmatchSlice))
//	}
//	fmt.Printf("SubmatchSlice: '%q'\n", SubmatchSlice)
//	var Need gcrawl.Unit
//	stringTofloat, _ := strconv.ParseFloat(SubmatchSlice[1], 64)
//	fmt.Printf("stringtoFloat '%v' \n", stringTofloat)
//	Need.Value  = stringTofloat
//	Need.Unit = SubmatchSlice[2]
//	log.Printf("Need.Value: '%v', Need.Unit: '%v'", Need.Value, Need.Unit)
//	log.Printf("SubmatchSlice: '%v'\n", SubmatchSlice)
//	return &Need, nil
//}