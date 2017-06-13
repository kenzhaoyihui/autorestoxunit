package adapters

import (
	"fmt"
	"testing"
)

func TestGetSummary(t *testing.T) {
	z := NewZoidberg("../final_results.json")
	r := z.GenTestSuites()
	fmt.Println(r)
}
