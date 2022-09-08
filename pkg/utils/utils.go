package utils

import "log"

// CheckErr handles any errors that come up while project is being executed.
func CheckErr(errMsg string, err error) {
	if err != nil {
		log.Fatal(errMsg, err)
	}
}
