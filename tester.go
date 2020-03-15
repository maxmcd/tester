package tester

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type Tester struct {
	*testing.T
}

// New creates a new tester instance
func New(t *testing.T) Tester {
	return Tester{t}
}

// Assert returns a testify/assert instance
func (t *Tester) Assert() *assert.Assertions {
	return assert.New(t)
}

func (t *Tester) ErrorContains(err error, contains string, msgAndArgs ...interface{}) {
	if err == nil {
		t.Assert().Error(err)
		return
	}
	t.Assert().Contains(err.Error(), contains, msgAndArgs...)
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func (t *Tester) PanicOnErr(err error, a ...interface{}) {
	if err != nil {
		if er, ok := err.(stackTracer); ok {
			t.Print(er)
		}
		if len(a) > 0 {
			t.Print(a...)
		}
		panic(err)
	}
}

func Print(a ...interface{}) {
	spew.Dump(a...)
}

func (t *Tester) Print(a ...interface{}) {
	Print(a...)
}
