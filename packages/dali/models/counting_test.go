// counting_test.go

package models

import (
    "testing"
    "os"
)

func TestCounting(m *testing.M) {
    code := m.Run()
    os.Exit(code)
}

func TestOne(t *testing.T) {
    actual := Counting(4,2)
    expected := 6
    if actual != expected {
        t.Errorf("Expected: %d. Got: %d", expected, actual)
    }
    
}

