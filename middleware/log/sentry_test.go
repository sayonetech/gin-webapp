package log

import (
	"fmt"
	"testing"
)

func TestCaptureErrorWithSentry(ts *testing.T) {
	testErr := fmt.Errorf("test error")
	err := CaptureErrorWithSentry(testErr)
	if err != nil {
		ts.Error("sentry error", err)
	}
}
