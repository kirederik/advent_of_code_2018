package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type testCase struct {
	Input  []string
	Output int
}

func TestPartI(t *testing.T) {
	testCases := []testCase{
		{
			Input:  []string{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"},
			Output: 138,
		},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r := PartI(f.Name())
		if r != v.Output {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output, i, r))
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
