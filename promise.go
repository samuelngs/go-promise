package promise

// Promise type
type Promise interface {
	State() State
	Result() interface{}
	Then(interface{}) Promise
	Catch(interface{}) Promise
}

// Accept type
type Accept func(...interface{})

// Decline type
type Decline func(...interface{})

// Executor type
type Executor func(Accept, Decline)

// New creates a promise
func New(fn Executor) Promise {
	p := &promise{
		fn:    fn,
		state: Pending,
		pre:   make(chan struct{}, 1),
		done:  make(chan struct{}, 1),
		ops:   new(operations),
	}
	return p
}

// Resolve resolves value
func Resolve(o ...interface{}) Promise {
	p := &promise{
		state: Resolved,
		pre:   nil,
		done:  make(chan struct{}, 1),
		ops:   new(operations),
	}
	if len(o) > 0 {
		p.val = o[0]
	}
	return p
}

// Reject resolves value
func Reject(o ...interface{}) Promise {
	p := &promise{
		state: Rejected,
		pre:   nil,
		done:  make(chan struct{}, 1),
		ops:   new(operations),
	}
	if len(o) > 0 {
		p.val = o[0]
	}
	return p
}
