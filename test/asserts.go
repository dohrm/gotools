package test

import (
    "testing"
    "fmt"
)


// Assert equals
func AssertEquals(t *testing.T, a interface{}, b interface{}, message string) {
    if a != b {
        if len(message) == 0 {
            message = fmt.Sprintf("%v != %v", a, b)
        }
        t.Fatal(message)
    }
}
