package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {

	// by defualt this will go to the standard output, it includes a datestamp
	// which makes it suitable for showing progress of binary rather
	// than just using fmt.Println() all the time.
	log.Print("Hello Logs!")
	// we can customise this to send logs elsewhere.

	// Here we're finding out our binary's filename.
	// we need this to label our log entries with it as best practice
	fmt.Println(os.Args)
	programName := filepath.Base(os.Args[0])
	fmt.Println(programName)

	// Here we are choosing what log-level and log-facility we want to use
	// https://golang.org/pkg/log/syslog/#Priority has the available options
	infoLog, _ := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	// here we are activating our log info settings, in preparation to
	// writing to system logs
	log.SetOutput(infoLog)

	// this sends the messate to /var/log/messages along with other places
	// as defined in the rsyslog.conf file.
	log.Println("This info log entry was generated by my golang binary")

}
