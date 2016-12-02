package service

import (
    "testing"
    "github.com/dohrm/gotools/test"
)

func TestServiceFunction_ChainRun(t *testing.T) {
    var a int
    var b int

    var f1 ServiceFunction = func() *ServiceError {
        a++
        return nil
    }
    var f2 ServiceFunction = func() *ServiceError {
        b++
        return nil
    }
    var f3 ServiceFunction = func() *ServiceError {
        return NewServiceError(nil, "", 42)
    }

    if f1.Run(f2) != nil {
        t.Fatal("f1.Run(f2) not return nil")
    }
    test.AssertEquals(t, a, 1, "f1.Run(f2) error with a")
    test.AssertEquals(t, b, 1, "f1.Run(f2) error with b")

    if f1.Chain(f2).Run(f3) == nil {
        t.Fatal("f1.Chain(f2).Run(f3) return nil")
    }
    test.AssertEquals(t, a, 2, "f1.Chain(f2).Run(f3) error with a")
    test.AssertEquals(t, b, 2, "f1.Chain(f2).Run(f3) error with b")

    if f1.Chain(f3).Run(f2) == nil {
        t.Fatal("f1.Chain(f3).Run(f2) return nil")
    }
    test.AssertEquals(t, a, 3, "f1.Chain(f3).Run(f2) error with a")
    test.AssertEquals(t, b, 2, "f1.Chain(f3).Run(f2) error with b")
}
