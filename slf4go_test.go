package slf4go

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dynamicgo/go-config/source/memory"

	config "github.com/dynamicgo/go-config"
)

func TestLog(t *testing.T) {
	logger := Get("test")
	SetLevel(Error)
	logger.Debug("test")
	logger.DebugF("test %s", "hello")
	logger.Error("error")
	logger.ErrorF("error %p", t)
}

func TestConfig(t *testing.T) {
	config := config.NewConfig()

	config.Load(memory.NewSource(memory.WithData([]byte(
		`
		{
			"default":{
				"backend":"console"
			},
			"logger":{
				"test":{
					"backend":"console",
					"level":"warn|error"
				}
			}
		}
		`,
	))))

	require.NoError(t, Load(config))

	logger := Get("test")

	logger.Warn("hello world")
	logger.Error("hello world")

	logger = Get("test3")

	logger.Trace("trace test")
}
