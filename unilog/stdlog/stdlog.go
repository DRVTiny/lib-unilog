package stdlog

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/dsnet/try"
	"github.com/DRVTiny/lib-tools/tools"
)

const DATE_F = "2006-01-02 15:04:05"

type LogLevel uint8

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var ll2str = [5]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

var str2ll = map[string]LogLevel{
	"DEBUG": DEBUG,
	"INFO":  INFO,
	"WARN":  WARN,
	"ERROR": ERROR,
	"FATAL": FATAL,
}

func NewLogLevel(level string) (ll LogLevel, err error) {
	var ok bool

	if ll, ok = str2ll[strings.ToUpper(level)]; ok {
		return
	} else {
		return 0, fmt.Errorf("unknown log level specified: %s", level)
	}
}

func (ll LogLevel) String() string {
	return ll2str[int(ll)]
}

type FormatFunc func(at time.Time, level LogLevel, msg string) string

type Unilog4L struct {
	formatFn FormatFunc
	level    LogLevel
	logger   *log.Logger
}

func NewUniStdLog(where2write io.Writer, level string, maybeFmtFn ...any) (u4l *Unilog4L, err error) {
	defer try.Handle(&err)

	sl := log.New(where2write, "", 0)
	u4l = &Unilog4L{
		level:  try.E1(NewLogLevel(level)),
		logger: sl,
	}

	if len(maybeFmtFn) > 0 {
		if fmtFn, ok := maybeFmtFn[0].(FormatFunc); ok {
			u4l.formatFn = fmtFn
		} else {
			return nil, fmt.Errorf("you MAY pass format func as a second argument, but %T passed instead", maybeFmtFn[0])

		}
	} else {
		u4l.formatFn = DefaultFormat
	}

	return
}

func DefaultFormat(at time.Time, level LogLevel, msg string) string {
	return strings.Join([]string{at.Format(DATE_F), level.String(), msg}, " | ")
}

func (u4l *Unilog4L) SetOtuput(out io.Writer) {
	u4l.logger.SetOutput(out)
}

func (u4l *Unilog4L) write(level LogLevel, msgParts []any) {
	if level < u4l.level {
		return
	}

	now := time.Now()

	var logFn func(v ...any)
	if level == FATAL {
		logFn = u4l.logger.Fatal
	} else {
		logFn = u4l.logger.Println
	}

	logFn(u4l.formatFn(now, level, strings.Join(tools.Anys2Strings(msgParts), " ")))
}

func (u4l *Unilog4L) writef(level LogLevel, fstr string, placehs []any) {
	if level < u4l.level {
		return
	}

	now := time.Now()

	var logFn func(v ...any)
	if level == FATAL {
		logFn = u4l.logger.Fatal
	} else {
		logFn = u4l.logger.Println
	}

	logFn(u4l.formatFn(now, level, fmt.Sprintf(fstr, placehs...)))
}

func (u4l *Unilog4L) Fatalf(fstr string, args ...any) {
	u4l.writef(FATAL, fstr, args)
}

func (u4l *Unilog4L) Fatal(args ...any) {
	u4l.write(FATAL, args)
}

func (u4l *Unilog4L) Errorf(fstr string, args ...any) {
	u4l.writef(ERROR, fstr, args)
}

func (u4l *Unilog4L) Error(args ...any) {
	u4l.write(ERROR, args)
}

func (u4l *Unilog4L) Warnf(fstr string, args ...any) {
	u4l.writef(WARN, fstr, args)
}

func (u4l *Unilog4L) Warn(args ...any) {
	u4l.write(WARN, args)
}

func (u4l *Unilog4L) Infof(fstr string, args ...any) {
	u4l.writef(INFO, fstr, args)
}

func (u4l *Unilog4L) Info(args ...any) {
	u4l.write(INFO, args)
}

func (u4l *Unilog4L) Debugf(fstr string, args ...any) {
	u4l.writef(DEBUG, fstr, args)
}

func (u4l *Unilog4L) Debug(args ...any) {
	u4l.write(DEBUG, args)
}

func (u4l *Unilog4L) Printf(fstr string, args ...any) {
	u4l.writef(DEBUG, fstr, args)
}

func (u4l *Unilog4L) Print(args ...any) {
	u4l.write(DEBUG, args)
}

func (u4l *Unilog4L) Println(args ...any) {
	u4l.write(DEBUG, args)
}

func (u4l *Unilog4L) CaptureOutput(f func()) string {
	var buf bytes.Buffer
	saveOut := u4l.logger.Writer()
	u4l.logger.SetOutput(&buf)
	f()
	u4l.logger.SetOutput(saveOut)

	return buf.String()
}
