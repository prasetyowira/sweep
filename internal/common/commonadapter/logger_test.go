package commonadapter

import (
	"context"
	"testing"

	"logur.dev/logur"
	"logur.dev/logur/conformance"
	"logur.dev/logur/logtesting"
)

func TestLogger(t *testing.T) {
	t.Run("WithFields", testLoggerWithFields)
	t.Run("WithContext", testLoggerWithContext)

	suite := conformance.TestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, conformance.TestLogger) {
			testLogger := &logur.TestLoggerFacade{}

			return NewLogger(testLogger), testLogger
		},
	}
	t.Run("Conformance", suite.Run)
}

func TestContextAwareLogger(t *testing.T) {
	t.Run("WithContext", testContextAwareLoggerWithContext)

	suite := conformance.TestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, conformance.TestLogger) {
			testLogger := &logur.TestLoggerFacade{}

			return NewContextAwareLogger(
				testLogger,
				func(ctx context.Context) map[string]interface{} {
					return nil
				},
			), testLogger
		},
	}
	t.Run("Conformance", suite.Run)
}

func testLoggerWithFields(t *testing.T) {
	testLogger := &logur.TestLoggerFacade{}

	fields := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	logger := NewLogger(testLogger).WithFields(fields)

	logger.Debug("sweep", nil)

	event := logur.LogEvent{
		Level:  logur.Debug,
		Line:   "sweep",
		Fields: fields,
	}

	logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
}

func testLoggerWithContext(t *testing.T) {
	testLogger := &logur.TestLoggerFacade{}

	logger := NewLogger(testLogger).WithContext(context.Background())

	logger.Debug("sweep", nil)

	event := logur.LogEvent{
		Level: logur.Debug,
		Line:  "sweep",
	}

	logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
}

func testContextAwareLoggerWithContext(t *testing.T) {
	testLogger := &logur.TestLoggerFacade{}

	logger := NewContextAwareLogger(
		testLogger,
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			}
		},
	).WithContext(context.Background())

	logger.Debug("sweep", nil)

	event := logur.LogEvent{
		Level: logur.Debug,
		Line:  "sweep",
		Fields: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
}
