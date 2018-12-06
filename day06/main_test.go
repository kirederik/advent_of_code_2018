package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type testCase struct {
	Input   []string
	Output1 int
	Output2 int
}

func TestPartI(t *testing.T) {
	testCases := []testCase{
		{
			Input: []string{
				"1, 1",
				"1, 6",
				"8, 9	",
				"8, 3",
				"3, 4",
				"5, 5",
			},
			Output1: 17,
			Output2: 72,
		},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r, r2 := PartIandII(f.Name())
		if r != v.Output1 || r2 != v.Output2 {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output1, i, r))
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output2, i, r2))
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
