package unilog4t

import "testing"

type UniLog4T testing.T

func (at *UniLog4T) Fatal(args ...any) {
	tl := (*testing.T)(at)
	tl.Fatal(args...)
}

func (at *UniLog4T) Fatalf(fstr string, args ...any) {
	tl := (*testing.T)(at)
	tl.Fatalf(fstr, args...)
}

func (at *UniLog4T) Error(args ...any) {
	tl := (*testing.T)(at)
	tl.Error(args...)
}

func (at *UniLog4T) Errorf(fstr string, args ...any) {
	tl := (*testing.T)(at)
	tl.Errorf(fstr, args...)
}

func (at *UniLog4T) log(args []any) {
	tl := (*testing.T)(at)
	tl.Log(args...)
}

func (at *UniLog4T) logf(fstr string, args []any) {
	tl := (*testing.T)(at)
	tl.Logf(fstr, args...)
}

func (at *UniLog4T) Warn(args ...any) {
	at.log(args)
}

func (at *UniLog4T) Warnf(fstr string, args ...any) {
	at.logf(fstr, args)
}

func (at *UniLog4T) Info(args ...any) {
	at.log(args)
}

func (at *UniLog4T) Infof(fstr string, args ...any) {
	at.logf(fstr, args)
}

func (at *UniLog4T) Print(args ...any) {
	at.log(args)
}

func (at *UniLog4T) Printf(fstr string, args ...any) {
	at.logf(fstr, args)
}

func (at *UniLog4T) Println(args ...any) {
	at.log(args)
}
