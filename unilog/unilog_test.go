package unilog

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	orig := *os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(fmt.Sprintf("failed to create pipe: %s", err.Error()))
	}

	*os.Stdout = *w
	f()
	os.Stdout = &orig
	w.Close()

	out, err := io.ReadAll(r)
	if err != nil {
		panic(fmt.Sprintf("failed to read from pipe: %s", err.Error()))
	}

	return string(out)
}

func TestDefaultOutput(t *testing.T) {
	rxOut := regexp.MustCompile(`^\d{4}-\d\d-\d\d \d\d:\d\d:\d\d \| INFO \| Hello, world!`)

	out := captureOutput(func() {
		Infof("Hello, %s!", "world")
	})

	assert.Regexp(t, rxOut, out)
}
