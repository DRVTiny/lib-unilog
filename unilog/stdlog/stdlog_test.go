package stdlog

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	rxOut := regexp.MustCompile(`^\d{4}-\d\d-\d\d \d\d:\d\d:\d\d \| INFO \| Hello, world!`)

	u4l, err := NewUniStdLog(os.Stdout, "info")
	assert.NoError(t, err)
	out := u4l.CaptureOutput(func() {
		u4l.Infof("Hello, %s!", "world")
	})

	assert.Regexp(t, rxOut, out)

	out = u4l.CaptureOutput(func() {
		u4l.Info("Hello,", "world!")
	})

	assert.Regexp(t, rxOut, out)
}

func TestNoOutput(t *testing.T) {
	u4l, err := NewUniStdLog(os.Stdout, "error")
	assert.NoError(t, err)
	out := u4l.CaptureOutput(func() {
		u4l.Infof("Hello, %s!", "world")
	})

	assert.Empty(t, out)

	out = u4l.CaptureOutput(func() {
		u4l.Info("Hello,", "world!")
	})

	assert.Empty(t, out)
}
