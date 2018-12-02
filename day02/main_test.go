package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type testCase struct {
	Input     []string
	Output    int
	StrOutput string
}

func TestCheckSum(t *testing.T) {
	testCases := []testCase{
		{
			Input: []string{
				"abcdef",
				"bababc",
				"abbcde",
				"abcccd",
				"aabcdd",
				"abcdee",
				"ababab",
			},
			Output: 12,
		},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r, _ := CheckSum(f.Name())
		if r != v.Output {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output, i, r))
		}

		os.Remove(f.Name()) // clean up
	}
	if len(errors) > 0 {
		t.Error("\n" + strings.Join(errors, "\n"))
	}
}

func TestCheckSumError(t *testing.T) {
	_, err := CheckSum("wrongpath")
	if err == nil {
		t.Error("No error returned")
	}
}

func TestFindLoop(t *testing.T) {
	testCases := []testCase{
		{
			Input: []string{
				"abcde",
				"fghij",
				"klmno",
				"pqrst",
				"fguij",
				"axcye",
				"wvxyz",
			},
			StrOutput: "fgij",
		},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r, _ := FindCommonLetters(f.Name())
		if r != v.StrOutput {
			errors = append(errors, fmt.Sprintf("Expecting %s for test case %d. Got %s", v.StrOutput, i, r))
		}

		os.Remove(f.Name()) // clean up
	}
	if len(errors) > 0 {
		t.Error("\n" + strings.Join(errors, "\n"))
	}

}

var result int

func BenchmarkCheckSum(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r, _ = CheckSum("./input")
	}
	result = r
}

var str string

func BenchmarkFindCommonLetters(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s, _ = FindCommonLetters("./input")
	}
	str = s
}

func createFileWithContent(s []string) (*os.File, error) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		return nil, err
	}
	if _, err := tmpfile.WriteString(strings.Join(s, "\n")); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}
	return tmpfile, nil
}
