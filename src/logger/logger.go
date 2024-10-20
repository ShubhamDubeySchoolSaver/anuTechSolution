package logger

import (
	"flag"
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	_ "github.com/natefinch/lumberjack"
)

var (
	Log *log.Logger
)

func init() {
	// Set the absolute path of the log file
	var logpath = "/app/log/logfile.log"

	flag.Parse()
	var file, err1 = os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.SetOutput(&lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    25, // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     3,  // days
	})
}
