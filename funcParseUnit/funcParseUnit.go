package funcParseUnit

import (
	"github.com/dennwc/gcrawl"
	"regexp"
	"strconv"
	"fmt"
	"log"
)

var reUnits = regexp.MustCompile(`^(\d+(?:\.\d+)?)\s([a-zA-Z]+)$`)

func ParseUnits(s string) (*gcrawl.Unit, error){
	//regular, _ := regexp.Compile(`^(\d+(?:\.\d+)?)\s([a-zA-Z]+)$`)
	//log.Printf("s: '%v'\n", s)
	submatchSlice := reUnits.FindStringSubmatch(s)
	if len(submatchSlice) == 0 {
		return nil, fmt.Errorf("The value is not found, submatchSlice: '%v'\n", submatchSlice)
	}
	log.Printf("SubmatchSlice: '%v'\n", submatchSlice)
	var need gcrawl.Unit
	stringTofloat, _ := strconv.ParseFloat(submatchSlice[1], 64)
	log.Printf("stringtoFloat '%v' \n", stringTofloat)
	need.Value  = stringTofloat
	need.Unit = submatchSlice[2]
	//log.Printf("Need.Value: '%v', Need.Unit: '%v'", need.Value, need.Unit)
	//log.Printf("SubmatchSlice: '%v'\n", submatchSlice)
	return &need, nil
}

