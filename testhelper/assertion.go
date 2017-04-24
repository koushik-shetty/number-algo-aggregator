package testhelper

import "testing"
import "reflect"

//Assert Custom asertion type
type Assert struct {
	*testing.T
}

//NoError asserts if the erro object is nil
func (t *Assert) NoError(err error, msgOnError ...interface{}) {
	if err != nil {
		t.Error(msgOnError...)
	}
}

//Equal compares the two values for deep equality
func (t *Assert) Equal(first, second interface{}, msgOnError ...interface{}) {
	if !reflect.DeepEqual(first, second) {
		t.Error(msgOnError...)
	}
}
