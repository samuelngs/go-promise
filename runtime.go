package promise

import (
	"reflect"
)

var kind = reflect.TypeOf(new(Promise)).Elem()

type promise struct {
	fn    Executor
	val   interface{}
	state State
	pre   chan struct{}
	done  chan struct{}
	ops   *operations
}

func (v *promise) Result() interface{} {
	if v.state == Pending && v.pre != nil {
		if v.fn != nil {
			go v.fn(v.accept, v.decline)
		}
	}
	go v.process()
	if v.state == Pending {
		<-v.done
	}
	return v.val
}

func (v *promise) State() State {
	return v.state
}

func (v *promise) Then(f interface{}) Promise {
	v.state = Pending
	v.ops.add(then, f)
	return v
}

func (v *promise) Catch(f interface{}) Promise {
	v.state = Pending
	v.ops.add(catch, f)
	return v
}

func (v *promise) accept(o ...interface{}) {
	if len(o) > 0 {
		v.val = o[0]
	}
	v.state = Resolved
	v.pre <- struct{}{}
}

func (v *promise) decline(o ...interface{}) {
	if len(o) > 0 {
		v.val = o[0]
	}
	v.state = Rejected
	v.pre <- struct{}{}
}

func (v *promise) resolve(x reflect.Value) (reflect.Value, bool) {
	k := x.Type()
	if k == kind {
		if p, ok := x.Interface().(Promise); ok {
			return reflect.ValueOf(p.Result()), p.State() == Resolved
		}
		return x, true
	}
	return x, true
}

func (v *promise) process() {
	if v.pre != nil {
		<-v.pre
	}
	if len(v.ops.o) > 0 {
		var pass = v.state == Resolved || v.state == Pending
		var args = make([]reflect.Value, v.ops.o[0].numOfArgs)
		if v.ops.o[0].numOfArgs > 0 {
			args[0] = reflect.ValueOf(v.val)
		}
		for _, o := range v.ops.o {
			switch {
			case pass && o.class == catch:
			case !pass && o.class == then:
			default:
				if o.numOfArgs < len(args) {
					args = args[:o.numOfArgs]
				}
				fn := v.ops.m[o.name]
				args = fn.Call(args)
				if len(args) > 0 {
					for i, arg := range args {
						o, c := v.resolve(arg)
						args[i] = o
						pass = c
					}
					v.val = args[0].Interface()
				}
			}
		}
		switch pass {
		case true:
			v.state = Resolved
		case false:
			v.state = Rejected
		}
	}
	v.done <- struct{}{}
}
