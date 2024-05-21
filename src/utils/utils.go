package utils

import (
	"fmt"
	"time"
)

type LogLevel byte

const (
	DEBUG LogLevel = iota
	VERBOSE
	INFO
	ERROR
)

func (s LogLevel) String() string {
	switch s {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case VERBOSE:
		return "VERBOSE"
	case ERROR:
		return "ERROR"
	}
	return "unknown"
}

func Log(logLevel LogLevel, data any) {
	fmt.Printf(
		"[DesignSphere-%s] [%+v] Request JSON: %+v\n",
		logLevel,
		time.Now().Format("Jan 02, 2006 3:04 PM"),
		data)
}
