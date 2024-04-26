package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var Logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	TimeFormat:      "[2006-01-02 15:04:05.000]",
	Level:           log.DebugLevel,
	Prefix:          "yukabot",
})
