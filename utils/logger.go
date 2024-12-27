package utils

import (
    "log"
    "os"
)
// Instead of writing log.println() multiple times throughout the codebase, you can use the function in logger.go to manage how logs are handled

var logger *log.Logger  // This is the pointer to a log.logger object that will handle logging messages

func InitLogger() {
    logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogMessage(message string) {
    if logger == nil {
        InitLogger()
    }
    logger.Println(message)
}
