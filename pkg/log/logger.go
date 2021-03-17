package log

import (
	"fmt"
	"log"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type coloredPrefix struct {
	prefix string
	color  string
}

var (
	logger *log.Logger

	errorPrefix = coloredPrefix{
		prefix: "ERR ",
		color:  colorRed,
	}

	warningPrefix = coloredPrefix{
		prefix: "WRN ",
		color:  colorYellow,
	}

	infoPrefix = coloredPrefix{
		prefix: "INF ",
		color:  colorCyan,
	}

	debugPrefix = coloredPrefix{
		prefix: "DBG ",
		color:  colorGreen,
	}
)

func init() {
	logger = log.Default()
}

func Error() *log.Logger {
	setColoredPrefix(logger, errorPrefix)
	return logger
}

func Warning() *log.Logger {
	setColoredPrefix(logger, warningPrefix)
	return logger
}

func Info() *log.Logger {
	setColoredPrefix(logger, infoPrefix)
	return logger
}

func Debug() *log.Logger {
	setColoredPrefix(logger, debugPrefix)
	return logger
}

func setColoredPrefix(logger *log.Logger, coloredPref coloredPrefix) {
	pref := fmt.Sprint(string(coloredPref.color), coloredPref.prefix, string(colorReset))
	logger.SetPrefix(pref)
}
