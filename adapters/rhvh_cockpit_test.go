package adapters

import (
	"testing"
	"fmt"
)

func TestCockpitGetSummary(t *testing.T) {
	z := NewCockpit("../he_result.json")
	r1 := z.GenTestSuites("RHEVM3")
	r2 := z.GenTestSuite()
	r3 := z.GenTestCases()

	fmt.Println(r1)
	fmt.Println(r2)

	for key, value := range r3 {
		fmt.Println(key, value)
	}
}
