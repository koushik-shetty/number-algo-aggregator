package testhelper

import "fmt"

// StubLogger Records calls to logs in the code.
type StubLogger struct {
	called bool
	msg    string
	args   []interface{}
}

//Errorf records the call to the logger.
func (sl *StubLogger) Errorf(msg string, args ...interface{}) {
	sl.called = true
	sl.msg = msg
	sl.args = args
}

//Called returns if the logger was called
func (sl *StubLogger) Called(msg string, args ...interface{}) bool {
	return sl.called
}

//ErrorMessage returns the error message that the logger would write to the destination.
func (sl *StubLogger) ErrorMessage() string {
	return fmt.Sprintf(sl.msg, sl.args)
}
