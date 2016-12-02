package service

type ServiceFunction func() *ServiceError


// Generate a function
func (f1 ServiceFunction) Chain(f2 ServiceFunction) ServiceFunction {
    return func() *ServiceError {
        e := f1()
        if e != nil {
            return e
        }
        return f2()
    }
}

// Execute f1, if no error occurs f2 is executed
func (f1 ServiceFunction) Run(f2 ServiceFunction) *ServiceError {
    return f1.Chain(f2)()
}