package intel

import (
	"testing"
	"github.com/dennwc/gcrawl"
	"reflect"
)

func TestParse(t *testing.T) {
	object, err := GetSpecification("94196")
	if err != nil{
		t.Errorf("Error is: '%v' \n", err)
	}
	//http://ark.intel.com/products/94196
	//fmt.Printf("RESULt is: A'%v' \n  B'%v' \n", object, err)
		got := object
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("unexpected match:\n%#v\n%#v", got, expect)
			for j:=0; j < len(got.Properties); j++{
				if !reflect.DeepEqual(expect.Properties[j], got.Properties[j]){
					t.Errorf("unexpected match:\n%#v\n%#v", expect.Properties[j], got.Properties[j])
				}
			}
		}
}

//var reUnits = regexp.MustCompile(`^(\d+(?:\.\d+)?)\s([a-zA-Z]+)$`)
//
//type caseStruct struct {
//	re     *regexp.Regexp
//	check  string
//	expect []string
//}

var expect = &gcrawl.Object{
	URL: "http://ark.intel.com/products/94196",
	Properties: []gcrawl.Property{
		{ID:"ProcessorNumber", Name:"Processor Number", Value: gcrawl.String("i7-6900K")},
		{ID:"StatusCodeText", Name:"Status", Value: gcrawl.String("Launched")},
		{ID:"BornOnDate", Name:"Launch Date", Value: gcrawl.String("Q2'16") },
		{ID:"Lithography", Name:"Lithography", Value: gcrawl.Unit{14, "nm"} },
		{ID:"Price1KUnits", Name:"Recommended Customer Price", Value: gcrawl.String("$1089.00 - $1109.00") },
		{ID:"CoreCount", Name:"# of Cores", Value: gcrawl.Int(8) },
		{ID:"ThreadCount", Name:"# of Threads", Value: gcrawl.Int(16) },
		{ID:"ClockSpeed", Name:"Processor Base Frequency", Value: gcrawl.Unit{3.20 ,"GHz"} },
		{ID:"ClockSpeedMax", Name:"Max Turbo Frequency", Value: gcrawl.Unit{3.70, "GHz"} },
		{ID:"Cache", Name:"Cache", Value: gcrawl.Unit{20, "MB"} },
		{ID:"TurboBoostMaxTechMaxFreq", Name:"Intel® Turbo Boost Max Technology 3.0 Frequency", Value: gcrawl.Unit{4.00, "GHz"} },
		{ID:"MaxTDP", Name:"TDP", Value: gcrawl.Unit{140, "W"} },
		{ID:"Embedded", Name:"Embedded Options Available", Value: gcrawl.Bool(false) },
		{ID:"DrcConflictFree", Name:"Conflict Free", Value: gcrawl.Bool(true) },
		{ID:"DatasheetUrl", Name:"Datasheet", Value: gcrawl.URL("http://www.intel.com/content/www/us/en/processors/core/core-technical-resources.html") },
		{ID:"ProductBriefUrl", Name:"Product Brief", Value: gcrawl.URL("http://www.intel.com/content/www/us/en/processors/core/core-i7-for-x-series-platform-product-brief.html") },
		{ID:"MaxMem", Name:"Max Memory Size (dependent on memory type)", Value: gcrawl.Unit{128, "GB"} },
		{ID:"MemoryTypes", Name:"Memory Types", Value: gcrawl.String("DDR4 2400/2133") },
		{ID:"NumMemoryChannels", Name:"Max # of Memory Channels", Value: gcrawl.Int(4) },
		{ID:"ECCMemory", Name:"ECC Memory Supported", Value: gcrawl.Bool(false) },
		{ID:"GraphicsModel", Name:"Processor Graphics", Value: gcrawl.String("None") },
		{ID:"ScalableSockets", Name:"Scalability", Value: gcrawl.String("1S Only") },
		{ID:"PCIExpressRevision", Name:"PCI Express Revision", Value: gcrawl.Float(3.0) },
		{ID:"NumPCIExpressPorts", Name:"Max # of PCI Express Lanes", Value: gcrawl.Int(40) },
		{ID:"SocketsSupported", Name:"Sockets Supported", Value: gcrawl.String("FCLGA2011-3") },
		{ID:"MaxCPUs", Name:"Max CPU Configuration", Value: gcrawl.Int(1) },
		{ID:"IsHalogenFree", Name:"Low Halogen Options Available", Value: gcrawl.String("See MDDS") },
		{ID:"TurboBoostMaxTechVersion", Name:"Intel® Turbo Boost Max Technology 3.0", Value: gcrawl.Bool(true) },
		{ID:"TBTVersion", Name:"Intel® Turbo Boost Technology", Value: gcrawl.Float(2.0) },
		{ID:"VProTechnology", Name:"Intel® vPro Technology", Value: gcrawl.Bool(false) },
		{ID:"HyperThreading", Name:"Intel® Hyper-Threading Technology", Value: gcrawl.Bool(true) },
		{ID:"VTX", Name:"Intel® Virtualization Technology (VT-x)", Value: gcrawl.Bool(true) },
		{ID:"VTD", Name:"Intel® Virtualization Technology for Directed I/O (VT-d)", Value: gcrawl.Bool(true) },
		{ID:"ExtendedPageTables", Name:"Intel® VT-x with Extended Page Tables (EPT)", Value: gcrawl.Bool(true) },
		{ID:"EM64", Name:"Intel® 64", Value: gcrawl.Bool(true) },
		{ID:"InstructionSet", Name:"Instruction Set", Value: gcrawl.String("64-bit") },
		{ID:"HaltState", Name:"Idle States", Value: gcrawl.Bool(true) },
		{ID:"SpeedstepTechnology", Name:"Enhanced Intel SpeedStep® Technology", Value: gcrawl.Bool(true) },
		{ID:"SmartResponseTechVersion", Name:"Intel® Smart Response Technology", Value: gcrawl.Bool(true) },
		{ID:"AESTech", Name:"Intel® AES New Instructions", Value: gcrawl.Bool(true) },
		{ID:"TXT", Name:"Trusted Execution Technology", Value: gcrawl.Bool(false) },
		{ID:"ExecuteDisable", Name:"Execute Disable Bit", Value: gcrawl.Bool(true) },
	},
}
