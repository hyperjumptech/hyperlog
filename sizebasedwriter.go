package hyperlog

import (
	"fmt"
	"os"
	"time"
)

type SizeBasedWriter struct {
	Directory        string
	Extension        string
	FilePrefix       string
	FileSuffix       string
	MaxFileSizeBytes int64

	currentFile *os.File
}

func (writer *SizeBasedWriter) getFile() *os.File {
	if writer.currentFile != nil {
		if info, err := os.Stat(writer.currentFile.Name()); err == nil {
			if info.Size() < writer.MaxFileSizeBytes {
				return writer.currentFile
			}
		}
		_ = writer.currentFile.Close()
		_ = os.Remove(writer.currentFile.Name())
		writer.currentFile = nil
	}

	fileName := fmt.Sprintf("%s%s%s%d%s.%s", writer.Directory, string(os.PathSeparator), writer.FilePrefix, int(time.Since(time.UnixMilli(0))/time.Second), writer.FileSuffix, writer.Extension)
	file, err := os.Create(fileName)
	if err != nil {
		return nil
	}
	writer.currentFile = file
	return file
}

func (writer *SizeBasedWriter) Write(p []byte) (n int, err error) {
	f := writer.getFile()
	if f != nil {
		return f.Write(p)
	}
	return 0, fmt.Errorf("can not nil file to write log")
}
