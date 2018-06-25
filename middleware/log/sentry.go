package log

import (
	"github.com/getsentry/raven-go"
)

func CaptureErrorWithSentry(err error) error {
	raven.CaptureError(err, nil)
	return nil
}
