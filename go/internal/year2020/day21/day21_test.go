package day21

import (
	"testing"

	"github.com/Saser/adventofcode/internal/testcase"
)

const (
	inputFile   = "../testdata/21"
	exampleFile = "testdata/example"
)

var (
	tcPart1 = testcase.NewFile("input", inputFile, "1815")
	tcPart2 = testcase.NewFile("input", inputFile, "kllgt,jrnqx,ljvx,zxstb,gnbxs,mhtc,hfdxb,hbfnkq")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", exampleFile, "5"),
		tcPart1,
	} {
		tc.Test(t, Part1)
	}
}

func BenchmarkPart1(b *testing.B) {
	tcPart1.Benchmark(b, Part1)
}

func TestPart2(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", exampleFile, "mxmxvkd,sqjhc,fvjkl"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
