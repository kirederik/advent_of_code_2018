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

func TestFindOverlaps(t *testing.T) {
	testCases := []testCase{
		{
			Input: []string{
				"#1 @ 1,3: 4x4",
				"#2 @ 3,1: 4x4",
				"#3 @ 5,5: 2x2",
			},
			Output: 4,
		},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r := FindOverlaps(f.Name())
		if r != v.Output {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output, i, r))
		}

		os.Remove(f.Name()) // clean up
	}
	if len(errors) > 0 {
		t.Error("\n" + strings.Join(errors, "\n"))
	}
}

func TestFindNonOverlapping(t *testing.T) {
	testCases := []testCase{
		{
			Input: []string{
				"#1 @ 1,3: 4x4",
				"#2 @ 3,1: 4x4",
				"#3 @ 5,5: 2x2",
			},
			Output: 3,
		},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r := FindNonOverlapping(f.Name())
		if r != fmt.Sprintf("#%d", v.Output) {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %s", v.Output, i, r))
		}

		os.Remove(f.Name()) // clean up
	}
	if len(errors) > 0 {
		t.Error("\n" + strings.Join(errors, "\n"))
	}
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
