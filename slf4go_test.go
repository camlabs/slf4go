package slf4go

import "testing"

func TestLog(t *testing.T) {
	logger := Get("test")
	SetLevel(Error)
	logger.Debug("test")
	logger.DebugF("test %s", "hello")
	logger.Error("error")
	logger.ErrorF("error %p", t)
}
