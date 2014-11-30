package godeal

import "log"

func Debug(v ...interface{}) {
	log.Println("[DEBUG]", v)
}

func Info(v ...interface{}) {
	log.Println("[INFO]", v)
}

func Warn(v ...interface{}) {
	log.Println("[WARN]", v)
}

func Error(v ...interface{}) {
	log.Println("[ERROR]", v)
}


