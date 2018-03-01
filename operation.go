package promise

import (
	"crypto/rand"
	"fmt"
	"reflect"
)

// class type
type class int

// List of available promise states
const (
	then class = iota
	catch
)

func id() string {
	b := new([16]byte)
	rand.Read(b[:])
	b[8] = (b[8] | 0x40) & 0x7F
	b[6] = (b[6] & 0xF) | (4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

type operation struct {
	class     class
	name      string
	numOfArgs int
}

type operations struct {
	o []*operation
	m map[string]reflect.Value
}

func (v *operations) add(c class, f interface{}) {
	id := id()
	if v.m == nil {
		v.m = make(map[string]reflect.Value)
	}
	d := reflect.ValueOf(f)
	if d.Kind() != reflect.Func {
		return
	}
	v.m[id] = d
	v.o = append(v.o, &operation{
		class:     c,
		name:      id,
		numOfArgs: d.Type().NumIn(),
	})
}
