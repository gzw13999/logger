package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func StartLogger() {

	// 打开日志文件
	//	file, err := os.OpenFile("./"+config.Cfg.Log.Runfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	//	if err != nil {
	//		log.Println("cannot open logfile %v\n", err)
	//	}

	//	// 创建MUX
	//	var r Repeater

	//	switch config.Cfg.Log.Output {
	//	case "both":
	//		r.out1 = os.Stdout
	//		r.out2 = file
	//	case "file":
	//		r.out2 = file
	//	}
	//	log.SetOutput(&r)
}

var (
	runlog       *log.Logger
	debuglog     *log.Logger
	infolog      *log.Logger
	errlog       *log.Logger
	runlogfile   *os.File
	debuglogfile *os.File
	infologfile  *os.File
	errlogfile   *os.File
	newline      string
	logpath      string
	date         time.Time
)

//func init() {

//	if runtime.GOOS == "windows" {
//		newline = "\r\n"
//	} else {
//		newline = ""
//	}
//}

func LogInit(path string) {

	date = time.Now()

	if runtime.GOOS == "windows" {
		newline = "\r\n"
	} else {
		newline = ""
	}

	if path == "" {
		logpath = "./"
	} else {
		logpath = path

		if f, err := os.Stat(logpath); err != nil {
			if err := os.MkdirAll(logpath, 0777); err != nil {
				panic("logger: os.MkdirAl err")
			}
		} else {
			if f.IsDir() != true {
				if err := os.MkdirAll(logpath, 0777); err != nil {
					panic("logger: os.MkdirAl err")
				}
			}
		}

		if _, err := os.Stat(logpath + "/run"); err != nil {
			if err := os.Mkdir(logpath+"/run", 0777); err != nil {
				panic(err)
			}
		}

		if _, err := os.Stat(logpath + "/debug"); err != nil {
			if err := os.Mkdir(logpath+"/debug", 0777); err != nil {
				panic("logger: os.MkdirAl err")
			}
		}

		if _, err := os.Stat(logpath + "/info"); err != nil {
			if err := os.Mkdir(logpath+"/info", 0777); err != nil {
				panic("logger: os.MkdirAl err")
			}
		}

		if _, err := os.Stat(logpath + "/error"); err != nil {
			if err := os.Mkdir(logpath+"/error", 0777); err != nil {
				panic("logger: os.MkdirAl err")
			}
		}

	}
	Run("log", logpath, "init...")
}

//func timeSub(t1, t2 time.Time) int {
//t1 = t1.UTC().Truncate(24 * time.Hour)
//t2 = t2.UTC().Truncate(24 * time.Hour)
//return int(t1.Sub(t2).Hours() / 24)
//}

func Run(args ...interface{}) {

	if runlog == nil {
		var err error
		if runlogfile, err = os.OpenFile(fmt.Sprintf("%v/run/%v.log", logpath, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0); err != nil {
			fmt.Printf("%s\r\n", err.Error())
			return
		} else {
			runlog = log.New(runlogfile, newline, log.LstdFlags)
		}
	}

	if date.Format("2006-01-02") != time.Now().Format("2006-01-02") {
		if err := runlogfile.Close(); err != nil {
		}
		var err error
		if runlogfile, err = os.OpenFile(fmt.Sprintf("%v/run/%v.log", logpath, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0); err != nil {
			fmt.Printf("%s\r\n", err.Error())
			return
		} else {
			runlog = log.New(runlogfile, newline, log.LstdFlags)
		}
	}
	runlog.Output(3, fmt.Sprintln(args...))
}

func Debug(args ...interface{}) {
	if debuglog == nil {
		var err error
		if debuglogfile, err = os.OpenFile(fmt.Sprintf("%v/debug/%v.log", logpath, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0); err != nil {
			fmt.Printf("%s\r\n", err.Error())
			return
		} else {
			debuglog = log.New(debuglogfile, newline, log.LstdFlags)
		}
	}

	if date.Format("2006-01-02") != time.Now().Format("2006-01-02") {
		if err := debuglogfile.Close(); err != nil {
		}
		var err error
		if debuglogfile, err = os.OpenFile(fmt.Sprintf("%v/debug/%v.log", logpath, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0); err != nil {
			fmt.Printf("%s\r\n", err.Error())
			return
		} else {
			debuglog = log.New(debuglogfile, newline, log.LstdFlags)
		}
	}

	debuglog.Output(3, fmt.Sprintln(args...))
}

func Info(args ...interface{}) {
	if infolog == nil {
		var err error
		if infologfile, err = os.OpenFile(fmt.Sprintf("%v/info/%v.log", logpath, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0); err != nil {
			fmt.Printf("%s\r\n", err.Error())
			return
		} else {
			infolog = log.New(infologfile, newline, log.LstdFlags)
		}
	}

	if date.Format("2006-01-02") != time.Now().Format("2006-01-02") {
		if err := infologfile.Close(); err != nil {
		}
		var err error
		if infologfile, err = os.OpenFile(fmt.Sprintf("%v/info/%v.log", logpath, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0); err != nil {
			fmt.Printf("%s\r\n", err.Error())
			return
		} else {
			infolog = log.New(infologfile, newline, log.LstdFlags)
		}
	}
	infolog.Output(3, fmt.Sprintln(args...))
}

func Error(args ...interface{}) {
	if errlog == nil {
		var err error
		if errlogfile, err = os.OpenFile(fmt.Sprintf("%v/error/%v.log", logpath, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0); err != nil {
			fmt.Printf("%s\r\n", err.Error())
			return
		} else {
			errlog = log.New(errlogfile, newline, log.LstdFlags)
		}
	}

	if date.Format("2006-01-02") != time.Now().Format("2006-01-02") {
		if err := errlogfile.Close(); err != nil {
		}
		var err error
		if errlogfile, err = os.OpenFile(fmt.Sprintf("%v/error/%v.log", logpath, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0); err != nil {
			fmt.Printf("%s\r\n", err.Error())
			return
		} else {
			errlog = log.New(errlogfile, newline, log.LstdFlags)
		}
	}

	errlog.Output(3, fmt.Sprintln(args...))
}
