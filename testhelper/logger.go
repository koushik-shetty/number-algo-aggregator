package testhelper

import "fmt"

// StubLogger Records calls to logs in the code.
type StubLogger struct {
	called bool
	msg    string
	args   []interface{}
}

func (sl *StubLogger) Errorf(msg string, args ...interface{}) {
	sl.called = true
	sl.msg = msg
	sl.args = args
}

func (sl *StubLogger) Called(msg string, args ...interface{}) bool {
	return sl.called
}

func (sl *StubLogger) ErrorMessage() string {
	return fmt.Sprintf(sl.msg, sl.args)
}
