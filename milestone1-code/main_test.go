package main

import (
	"bytes"
	"fmt"
	"testing"
)

func reportFailure(got, expected interface{}, t *testing.T) {
	t.Errorf("Did not get expected result. Got: '%v', wanted: '%v'", got, expected)
}

var flagtests = []struct {
	in  []string
	out projectConfig
}{
	{[]string{"-n", "name", "-d", "disk", "-r", "repo"}, projectConfig{"disk", "name", "repo", false}},
	{[]string{"-n", "name", "-d", "disk", "-r", "repo", "-s", "true"}, projectConfig{"disk", "name", "repo", true}},
}

func TestSetupParseFlags(t *testing.T) {
	buffer := new(bytes.Buffer)

	for _, tt := range flagtests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got, err := setupParseFlags(buffer, tt.in)
			if err != nil {
				t.Errorf("Parsing error")
				return
			}
			expected := tt.out
			if got != expected {
				reportFailure(got, expected, t)
			}
		})
	}
}

var validateTests = []struct {
	in  projectConfig
	out int
}{
	{projectConfig{"disk", "name", "repo", false}, 0},
	{projectConfig{"", "", "", false}, 3},
	{projectConfig{"disk", "", "", false}, 2},
	{projectConfig{"disk", "name", "", false}, 1},
}

func TestValidateConf(t *testing.T) {

	for _, tt := range validateTests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			var got = validateConf(tt.in)
			expected := tt.out
			if len(got) != expected {
				reportFailure(got, expected, t)
			}
		})
	}
}

func TestGenerateScaffold(t *testing.T) {
	buffer := new(bytes.Buffer)

	generateScaffold(buffer, projectConfig{"disk", "name", "repo", false})
	got := buffer.String()

	expected := "Generating scaffold for project name in disk"
	if got != expected {
		reportFailure(got, expected, t)
	}
}
