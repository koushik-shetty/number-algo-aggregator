package testhelper

import "testing"
import "reflect"

type Assert struct {
	*testing.T
}

func (t *Assert) NoError(err error, msgOnError ...interface{}) {
	if err != nil {
		t.Error(msgOnError...)
	}
}

func (t *Assert) Equal(first, second interface{}, msgOnError ...interface{}) {
	if !reflect.DeepEqual(first, second) {
		t.Error(msgOnError...)
	}
}
