// product_test.go

package models

import (
    "testing"
    "os"
)

func TestProduct(m *testing.M) {
    code := m.Run()
    os.Exit(code)
}

