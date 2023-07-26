package logger

import (
	"log"
	"os"
)

var Write *conditionalLogger

func init() {
	Write = &conditionalLogger{
		consoleLogger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

type conditionalLogger struct {
	hasErrorOccurred bool
	fileLogger       *log.Logger
	consoleLogger    *log.Logger
	logFile          *os.File
}

func (c *conditionalLogger) Log(v ...interface{}) {
	if c.hasErrorOccurred {
		c.fileLogger.Println(v...)
	} else {
		c.consoleLogger.Println(v...)
	}
}

func (c *conditionalLogger) Logf(format string, v ...interface{}) {
	if c.hasErrorOccurred {
		c.fileLogger.Printf(format, v...)
	} else {
		c.consoleLogger.Printf(format, v...)
	}
}

func (c *conditionalLogger) Error(v ...interface{}) {
	c.initFileLoggerIfNecessary()
	c.fileLogger.Println(v...)
}

func (c *conditionalLogger) Errorf(format string, v ...interface{}) {
	c.initFileLoggerIfNecessary()
	c.fileLogger.Printf(format, v...)
}

// Private helper method to initialize the file logger if an error occurs.
func (c *conditionalLogger) initFileLoggerIfNecessary() {
	if !c.hasErrorOccurred {
		c.logFile, _ = os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		c.fileLogger = log.New(c.logFile, "", log.LstdFlags)
		c.hasErrorOccurred = true
	}
}
