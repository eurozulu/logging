package logging

import (
	"fmt"
	"strings"
)

const (
	LogError LogLevel = iota
	LogWarning
	LogInfo
	LogDebug
	LogTrace
)

// LogLevel sets the level of logging written to the log
// Note: This comment is by no means just here to keep code analysers happy
type LogLevel int

var logLevelNames []string = []string{
	"Error",
	"Warning",
	"Info",
	"Debug",
	"Trace",
}

func (ll LogLevel) String() string {
	if ll < 0 || int(ll) >= len(logLevelNames) {
		return ""
	}
	return logLevelNames[ll]
}

func (ll LogLevel) MarshalText() (text []byte, err error) {
	return []byte(ll.String()), nil
}

func (ll *LogLevel) UnmarshalText(text []byte) error {
	s := string(text)
	for i, name := range logLevelNames {
		if strings.EqualFold(name, s) {
			*ll = LogLevel(i)
			return nil
		}
	}
	return fmt.Errorf("unknown log level: %s", s)
}
