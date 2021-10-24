package logger

import "log"

func Error(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Info(v ...interface{}) {
	if v != nil {
		log.Println(v)
	}
}
