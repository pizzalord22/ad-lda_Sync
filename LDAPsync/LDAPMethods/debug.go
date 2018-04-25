package LDAPMethods

import (
	"log"
	"fmt"
)

func DebugMain(errorString string, err error, line int) {
	errormsg := fmt.Sprintf("%s", err)
	if err == nil {
		errormsg = fmt.Sprintf("%s", "none")
	}
	if DebugMainState == true {
		log.Printf("%s, error message: %s, main.go  trigger line %d\r\n", errorString, errormsg, line)
	}
}
func DebugCompare(errorString string, err error, line int) {
	errormsg := fmt.Sprintf("%s", err)
	if err == nil {
		errormsg = fmt.Sprintf("%s", "none")
	}
	if DebugMainState == true {
		log.Printf("%s, error message: %s, comapare.go trigger line %d\r\n", errorString, errormsg, line)
	}
}
func DebugModify(errorString string, err error, line int) {
	errormsg := fmt.Sprintf("%s", err)
	if err == nil {
		errormsg = fmt.Sprintf("%s", "none")
	}
	if DebugMainState == true {
		log.Printf("%s, error message: %s, modify.go trigger line %d\r\n", errorString, errormsg, line)
	}
}
func DebugDelete(errorString string, err error, line int) {
	errormsg := fmt.Sprintf("%s", err)
	if err == nil {
		errormsg = fmt.Sprintf("%s", "none")
	}
	if DebugMainState == true {
		log.Printf("%s, error message: %s, delete.go trigger line %d\r\n", errorString, errormsg, line)
	}
}
func DebugInitial(errorString string, err error, line int) {
	errormsg := fmt.Sprintf("%s", err)
	if err == nil {
		errormsg = fmt.Sprintf("%s", "none")
	}
	if DebugMainState == true {
		log.Printf("%s, error message: %s, initial.go trigger line %d\r\n", errorString, errormsg, line)
	}
}
func DebugAdd(errorString string, err error, line int) {
	errormsg := fmt.Sprintf("%s", err)
	if err == nil {
		errormsg = fmt.Sprintf("%s", "none")
	}
	if DebugMainState == true {
		log.Printf("%s, error message: %s, add.go trigger line %d\r\n", errorString, errormsg, line)
	}
}
