package util

import (
	"fmt"
	"os"
	"time"
)

const (
	DATETIME_LAYOUT string = "2006-01-02 15:04:05"
	DATETIME_LENGTH int    = len(DATETIME_LAYOUT)
)

func PrintInfo(format string, val ...interface{}) {
	fmt.Fprintf(os.Stdout, "%s [I] %s\n", time.Now().Format(DATETIME_LAYOUT), fmt.Sprintf(format, val...))
}

func PrintError(format string, val ...interface{}) {
	fmt.Fprintf(os.Stdout, "%s [E] %s\n", time.Now().Format(DATETIME_LAYOUT), fmt.Sprintf(format, val...))
}
