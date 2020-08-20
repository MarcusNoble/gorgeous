package filters

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"fmt"

	"github.com/fatih/color"
)

const (
	// runningTestPrefix    = "=== RUN   "                   // running test
	baseIndent            = "    "           // base indentation for subtests
	passingTestPrefix     = "--- PASS: "     // passing test
	failingTestPrefix     = "--- FAIL: "     // failing test
	coveragePrefix        = "coverage: "     // normal coverage prefix, useful for generating bars
	packageCoveragePrefix = "ok  	"          // package level coverage prefix
	buildFailedSuffix     = "[build failed]" // package level build failure

	// test output
	indentedTestPrefix = "\t"
)

// All filters in a slice
var All = []func(string) string{
	Pass,
	Fail,
	SubTest,
	Indent,
	PkgCoverage,
	RegCoverage,
}

// colors
var green = color.New(color.FgHiGreen).Sprintf
var red = color.New(color.FgHiRed).Sprintf

// Pass defines a passing test filter
func Pass(txt string) string {
	txt = filterPrefix(txt, passingTestPrefix)
	if txt != "" {
		return green("✔   ") + txt
	}
	return txt
}

// Fail defines a failing test filter
func Fail(txt string) string {
	result := filterPrefix(txt, failingTestPrefix)
	if result != "" {
		return red("✘   ") + result
	}
	result = filterSuffix(txt, buildFailedSuffix)
	if result != "" {
		return red("✘   ") + result
	}
	return result
}

// SubTest defines an indented sub test filter
func SubTest(txt string) string {
	if has(strings.TrimPrefix(txt, baseIndent), passingTestPrefix) != "" {
		txt = strings.TrimPrefix(txt, baseIndent)
		txt = Pass(txt)
		return green("├── ") + txt
	}

	if has(strings.TrimPrefix(txt, baseIndent), failingTestPrefix) != "" {
		txt = strings.TrimPrefix(txt, baseIndent)
		txt = Fail(txt)
		return red("├── ") + txt
	}

	if has(strings.TrimPrefix(txt, baseIndent), indentedTestPrefix) != "" {
		txt = strings.TrimPrefix(txt, baseIndent)
		return Indent(txt)
	}

	return ""
}

// Indent defines an indented line filter
func Indent(txt string) string {
	if txt = has(txt, indentedTestPrefix); txt != "" {
		return color.HiBlueString(txt)
	}
	return ""
}

// PkgCoverage filters the package coverage stats
func PkgCoverage(txt string) string {
	if txt = filterPrefix(txt, packageCoveragePrefix); txt != "" {
		parts := strings.Split(txt, "\t")
		parts[0] = filepath.Base(parts[0])
		txt = strings.Join(parts[0:2], "\t")

		return color.HiMagentaString("┗ " + txt + "\n")
	}
	return ""
}

// RegCoverage filters regular coverage flags
func RegCoverage(txt string) string {
	if txt = has(txt, coveragePrefix); txt != "" {
		return color.HiMagentaString(parseCoverage(txt))
	}
	return ""
}

func filterPrefix(txt, prefix string) string {
	if strings.HasPrefix(txt, prefix) {
		return strings.TrimPrefix(txt, prefix)
	}
	return ""
}

func filterSuffix(txt, suffix string) string {
	if strings.HasSuffix(txt, suffix) {
		return strings.TrimSuffix(txt, suffix)
	}
	return ""
}

func has(txt, prefix string) string {
	if strings.HasPrefix(txt, prefix) {
		return txt
	}
	return ""
}

func parseCoverage(txt string) string {
	const pre = "\nCoverage: %.1f%% ╠"
	const suf = "╣"
	const fil = "▓"
	const max = 10

	txt = strings.TrimPrefix(txt, coveragePrefix)
	parts := strings.Split(txt, "%")
	num, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		log.Fatalf("cannot parse coverage: %v", err)
	}
	if int(num) == 0 {
		return fmt.Sprintf(pre, num) + strings.Repeat(" ", 10) + suf
	}
	sum := int(num) / 10
	numSpaces := 10 - sum

	return fmt.Sprintf(pre, num) + strings.Repeat(fil, sum) + strings.Repeat(" ", numSpaces) + suf
}
