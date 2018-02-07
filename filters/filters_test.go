package filters

import (
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
}
