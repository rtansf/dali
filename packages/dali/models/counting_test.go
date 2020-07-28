// counting_test.go

package models

import (
    "io/ioutil"
    "testing"
    "os"
    "fmt"
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

    data2, err2 := ioutil.ReadFile("sql/ddl.sql")
    if err2 != nil {
        t.Errorf("ddl.sql not found")
    }
    fmt.Println(string(data2))

}

