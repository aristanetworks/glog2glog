// Copyright (c) 2022 Arista Networks, Inc.  All rights reserved.
// Arista Networks, Inc. Confidential and Proprietary.
// Subject to Arista Networks, Inc.'s EULA.
// FOR INTERNAL USE ONLY. NOT FOR DISTRIBUTION.

package glog

import (
	"fmt"

	aglog "github.com/aristanetworks/glog"
)

var MaxSize = aglog.MaxSize

var Stats = aglog.Stats

type Level = aglog.Level
type OutputStats = aglog.OutputStats
type Verbose = aglog.Verbose

func V(l Level) Verbose {
	return aglog.V(l)
}

func CopyStandardLogTo(name string) {
	aglog.CopyStandardLogTo(name)
}

func Error(args ...interface{}) {
	aglog.ErrorDepth(1, args...)
}

func ErrorDepth(depth int, args ...interface{}) {
	aglog.ErrorDepth(depth+1, args...)
}

func Errorf(format string, args ...interface{}) {
	aglog.ErrorDepth(1, fmt.Sprintf(format, args...))
}

func Errorln(args ...interface{}) {
	aglog.ErrorDepth(1, args...)

}

func Exit(args ...interface{}) {
	aglog.ExitDepth(1, args...)
}

func ExitDepth(depth int, args ...interface{}) {
	aglog.ExitDepth(depth+1, args...)
}

func Exitf(format string, args ...interface{}) {
	aglog.ExitDepth(1, fmt.Sprintf(format, args...))
}

func Exitln(args ...interface{}) {
	aglog.ExitDepth(1, args...)
}

func Fatal(args ...interface{}) {
	aglog.FatalDepth(1, args...)
}

func FatalDepth(depth int, args ...interface{}) {
	aglog.FatalDepth(depth+1, args...)
}

func Fatalf(format string, args ...interface{}) {
	aglog.FatalDepth(1, fmt.Sprintf(format, args...))
}

func Fatalln(args ...interface{}) {
	aglog.FatalDepth(1, args...)
}

func Flush() {
	aglog.Flush()
}

func Info(args ...interface{}) {
	aglog.InfoDepth(1, args...)
}

func InfoDepth(depth int, args ...interface{}) {
	aglog.InfoDepth(depth+1, args...)
}

func Infof(format string, args ...interface{}) {
	aglog.InfoDepth(1, fmt.Sprintf(format, args...))
}

func Infoln(args ...interface{}) {
	aglog.InfoDepth(1, args...)
}

func Warning(args ...interface{}) {
	aglog.WarningDepth(1, args...)
}

func WarningDepth(depth int, args ...interface{}) {
	aglog.WarningDepth(depth+1, args...)
}

func Warningf(format string, args ...interface{}) {
	aglog.WarningDepth(1, fmt.Sprintf(format, args...))
}

func Warningln(args ...interface{}) {
	aglog.WarningDepth(1, args...)
}
