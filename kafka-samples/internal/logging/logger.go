package utils

import (
	"fmt"
	"log"
)

var levelNames = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

// Level log level
type Level uint8

const (
	//TRACE level
	TRACE Level = iota
	// DEBUG  level
	DEBUG
	// INFO info level
	INFO
	// WARN the pattern has been found but some condition have not passed
	WARN
	// ERROR the pattern has been found but some condition have not passed
	ERROR
	// FATAL the pattern has been found but some condition have not passed
	FATAL
)

var activeLevel Level = INFO

func setLogLevel(level Level) {
	activeLevel = level
}

// SetLogLevel set the log level
func SetLogLevel(level string) {

	for index, levelName := range levelNames {
		if levelName == level {
			setLogLevel(Level(index))
			return
		}
	}
}

func _log(level Level, message string, args ...interface{}) {
	//params := append([]interface{}{level + ": ", message + " "}, args...)

	if level < activeLevel {
		return
	}
	var text = fmt.Sprintf(message, args...)

	if level == FATAL {
		log.Fatalf("%s : %s", levelNames[level], text)
	} else {
		log.Printf("%s : %s", levelNames[level], text)
	}
}

// Debug info message
func Trace(message string, args ...interface{}) {
	_log(TRACE, message, args...)
}

// Debug info message
func Debug(message string, args ...interface{}) {
	_log(DEBUG, message, args...)
}

// Info info message
func Info(message string, args ...interface{}) {
	_log(INFO, message, args...)
}

// Warn info message
func Warn(message string, args ...interface{}) {
	_log(INFO, message, args...)
}

// Error info message
func Error(message string, args ...interface{}) {
	_log(ERROR, message, args...)
}

// Fatal info message
func Fatal(message string, args ...interface{}) {
	_log(FATAL, message, args...)
}
