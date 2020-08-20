package filters

import (
	"fmt"
	"testing"
)

func Test_parseCoverage(t *testing.T) {
	exp1 := "\nCoverage: 0.0% ╠          ╣"
	exp2 := "\nCoverage: 89.4% ╠▓▓▓▓▓▓▓▓  ╣"

	res1 := parseCoverage("coverage: 0.0% of statements")
	res2 := parseCoverage("coverage: 89.4% of statements")

	if res1 != exp1 {
		t.Errorf("wanted: %s, but got: %s", exp1, res1)
	}
	if res2 != exp2 {
		t.Errorf("wanted: %s, but got: %s", exp2, res2)
	}
	t.Run("sub test output", func(t *testing.T) {
		fmt.Println("something")
	})
}

func Test_fail(t *testing.T) {
	if Fail("FAIL    example/pkg/failing [build failed]") == "" {
		t.Errorf("Expecting a failure\n")
	}

	if Fail("--- FAIL: Test_fail (0.00s)") == "" {
		t.Errorf("Expecting a failure\n")
	}

	if Fail("ok      github.com/MarcusNoble/gorgeous/filters 0.026s") != "" {
		t.Errorf("Wasn't expecting a failure\n")
	}
}
