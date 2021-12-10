package logger

import "log"

func Error(e error) {
	if e != nil {
		log.Println(e)
	}
}

func Info(v ...interface{}) {
	if v != nil {
		log.Println(v)
	}
}
