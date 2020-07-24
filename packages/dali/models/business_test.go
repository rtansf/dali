// business_test.go

package models

import (
    "testing"
    "os"
)

func TestBusiness(m *testing.M) {
    code := m.Run()
    os.Exit(code)
}

