package asserts

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func AssertStringEqual(expected string, actual string, t *testing.T) {
	if actual != expected {
		fmt.Printf("expected %s\nreceived %s\n", expected, actual)
		t.Fail()
	}
}

func AssertFileStringEqual(expected string, filename string, t *testing.T) {
	content, err := ioutil.ReadFile(filename)
	if err == nil {
		AssertStringEqual(expected, string(content), t)
	} else {
		fmt.Printf("error reading file %s\n", filename)
		t.Fail()
	}
}
