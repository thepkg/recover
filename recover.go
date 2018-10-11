// Package recover contains basic functions
// which helps to work with `recover` in bit pleasant way.
package recover

// All performs call cb function with recovered value.
func All(cb func(v interface{})) {
	r := recover()

	if r == nil {
		return
	}

	cb(r)
}

// One performs call cb function with recovered value
// in case when recovered value equals to e.
func One(e error, cb func(v interface{})) {
	r := recover()

	if r == nil {
		return
	}

	if r == e {
		cb(r)
		return
	}

	panic(r)
}

// Any performs call cb function with recovered value
// in case when recovered value exists in slice errors.
func Any(errors []error, cb func(v interface{})) {
	r := recover()

	if r == nil {
		return
	}

	if len(errors) == 0 || inErrors(r, errors) {
		cb(r)
		return
	}

	panic(r)
}

func inErrors(e interface{}, errors []error) bool {
	for _, err := range errors {
		if e == err {
			return true
		}
	}

	return false
}
