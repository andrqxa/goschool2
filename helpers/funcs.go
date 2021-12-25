package helpers

import "log"

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
