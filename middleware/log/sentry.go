package log

import (
	"github.com/getsentry/raven-go"
)

func CaptureErrorWithSentry(err error) error {
	raven.CaptureErrorAndWait(err, nil)
	return nil
}
