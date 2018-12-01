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
	Output int64
}

func TestSumSequence(t *testing.T) {
	testCases := []testCase{
		{Input: []string{"+1", "-2", "+3", "+1"}, Output: 3},
		{Input: []string{"+1", "-1", "+2", "-2"}, Output: 0},
		{Input: []string{"0", "0", "0", "0"}, Output: 0},
		{Input: []string{"+1", "+10", "+100", "+1234"}, Output: 1345},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r, _ := SumSequence(f.Name())
		if r != v.Output {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output, i, r))
		}

		os.Remove(f.Name()) // clean up
	}
	if len(errors) > 0 {
		t.Error("\n" + strings.Join(errors, "\n"))
	}
}

func TestSumSequenceError(t *testing.T) {
	_, err := SumSequence("wrongpath")
	if err == nil {
		t.Error("No error returned")
	}
}

func ExampleSumSequence() {
	f, _ := createFileWithContent([]string{"+1", "-2", "+3", "+1"})
	fmt.Println(SumSequence(f.Name()))
	// Output:
	// 3 <nil>
}

func TestFindLoop(t *testing.T) {
	testCases := []testCase{
		{Input: []string{"+1", "-2", "+3", "+1"}, Output: 2},
		{Input: []string{"+1", "-1"}, Output: 0},
		{Input: []string{"-6", "+3", "+8", "+5", "-6"}, Output: 5},
		{Input: []string{"+7", "+7", "-2", "-7", "-4"}, Output: 14},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r, _ := FindLoop(f.Name())
		if r != v.Output {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output, i, r))
		}

		os.Remove(f.Name()) // clean up
	}
	if len(errors) > 0 {
		t.Error("\n" + strings.Join(errors, "\n"))
	}

}

func ExampleFindLoop() {
	f, _ := createFileWithContent([]string{"+1", "-2", "+3", "+1"})
	fmt.Println(FindLoop(f.Name()))
	// Output:
	// 2 <nil>
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
