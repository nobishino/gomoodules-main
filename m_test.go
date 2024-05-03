package m_test

import (
	"testing"

	m "github.com/nobishino/gomoodules-main"
)

func ExampleMainModulesFunction() {
	m.MainModulesFunction()
	// Output: This is gomodules-m. I am developing this module.
}

func TestAdd(t *testing.T) {
	got := m.Add(1, 2)
	if got != 3 {
		t.Errorf("want Add(1, 2) = 3; got %d", got)
	}
}
