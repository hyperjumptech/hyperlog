package hyperlog

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	FormatYear  = "2006"
	FormatMonth = "2006_01"
	FormatDate  = "2006_01_02"

	DateMode TimeMode = iota
	MonthMode
	YearMode
)

type TimeMode int

type TimeBasedWriter struct {
	Directory  string
	Extension  string
	FilePrefix string
	FileSuffix string
	Mode       TimeMode
	Log        LogEngine

	currentFileName string
	currentFile     *os.File
}

func (lfw *TimeBasedWriter) getFile() *os.File {
	t := ""
	now := time.Now()
	switch lfw.Mode {
	case DateMode:
		t = now.Format(FormatDate)
	case MonthMode:
		t = now.Format(FormatMonth)
	case YearMode:
		t = now.Format(FormatYear)
	default:
		t = now.Format(FormatYear)
	}
	fileName := fmt.Sprintf("%s%s%s%s%s.%s", lfw.Directory, string(os.PathSeparator), lfw.FilePrefix, t, lfw.FileSuffix, lfw.Extension)
	if lfw.currentFileName != fileName {
		if lfw.currentFile != nil {
			err := lfw.currentFile.Close()
			if err != nil {
				panic("can not close log file")
			}
		}
	} else {
		if lfw.currentFile != nil {
			return lfw.currentFile
		}
	}

	if _, err := os.Stat(fileName); err != nil && errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(fileName)
		if err != nil {
			panic("can not create log file " + fileName)
		}
		lfw.currentFile = file
		lfw.currentFileName = fileName
		return file
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic("can not open log file " + fileName)
	}
	lfw.currentFile = file
	lfw.currentFileName = fileName
	return file
}

func (lfw *TimeBasedWriter) Write(p []byte) (n int, err error) {
	f := lfw.getFile()
	if f != nil {
		return f.Write(p)
	}
	return 0, fmt.Errorf("can not nil file to write log")
}
