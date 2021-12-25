package helpers

import "log"

// PanicIf raises panic if error's occured
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// LogFatal writtes error in log
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
