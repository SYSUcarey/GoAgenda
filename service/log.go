package service

import (
	"io"
	"log"
	"os"
)

var (
	// Info : Discard
	Info    *log.Logger
	// Error : Stderr
	Error   *log.Logger
)

var errlog, infolog *os.File

var infofile = os.Getenv("GOPATH") + "/src/github.com/chenf99/GoAgenda/data/info.log"
var errfile = os.Getenv("GOPATH") + "/src/github.com/chenf99/GoAgenda/data/error.log"

var infowriter, errwriter []io.Writer

var fileAndStdoutWriter1 io.Writer
var fileAndStdoutWriter2 io.Writer

func init() {
	infolog = getLogFile(infofile)
	errlog = getLogFile(errfile)
	errwriter = []io.Writer{
		errlog,
		os.Stdout,
	}
	infowriter = []io.Writer{
		os.Stdout,
		infolog,
	}
	fileAndStdoutWriter1 = io.MultiWriter(infowriter...)
	fileAndStdoutWriter2 = io.MultiWriter(errwriter...)
	Info =  log.New(fileAndStdoutWriter2, "ERROR: ", log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(fileAndStdoutWriter2, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)
}

func getLogFile(logPath string) *os.File  {
	//判断文件是否存在
	_, err := os.Stat(userfile)
	if os.IsNotExist(err) {
		os.Mkdir(os.Getenv("GOPATH") + "/src/github.com/chenf99/GoAgenda/data", 0777)
	}

	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("file open error : %v\n", err)
	}
	return file;
}