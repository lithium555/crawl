package intel

import "testing"

func TestIntelStuff(t *testing.T) {
	GetSpecification("94196") // /Intel-Core-i7-6900K-Processor-20M-Cache-up-to-3_70-GHz")
	ListProducts("88392")
	GetFamiliesId()
}


