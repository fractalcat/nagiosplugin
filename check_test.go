package nagiosplugin

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	c := NewCheck("CHECK TEST")

	expected := "CHECK TEST CRITICAL: 200000 terrifying space monkeys in the engineroom | space_monkeys=200000c;10000;100000;0;4294967296"

	nSpaceMonkeys := float64(200000)
	maxSpaceMonkeys := float64(1 << 32)

	_ = c.AddPerfDatum("space_monkeys", "c", nSpaceMonkeys, 0, maxSpaceMonkeys, 10000, 100000)
	c.AddResult(CRITICAL, fmt.Sprintf("%v terrifying space monkeys in the engineroom", nSpaceMonkeys))

	// Check a WARNING can't override a CRITICAL
	c.AddResult(WARNING, fmt.Sprintf("%v slightly annoying space monkeys in the engineroom", nSpaceMonkeys))

	result := c.String()
	if expected != result {
		t.Errorf("Expected check output %v, got check output %v", expected, result)
	}
}
