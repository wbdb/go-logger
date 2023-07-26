# Logger Go Package

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

The `logger` package is a simple Go logging library that provides conditional logging to the console and a log file. It allows you to log messages to the console when no error has occurred and switch to logging to a file once an error is encountered. This can be useful when you want to keep track of errors separately from regular log messages.

## Installation

To use this package in your Go project, you can use `go get`:

```bash
go get github.com/wbdb/logger
```

## Usage

Import the `logger` package into your Go code:

```go
import "github.com/wbdb/logger"
```

### Initialization

To initialize the logger, simply import the package. The logger will be available as `logger.Write`.

```go
logger.Write = &logger.conditionalLogger{
    consoleLogger: log.New(os.Stdout, "", log.LstdFlags),
}
```

### Logging

The logger provides three main logging methods:

1. `Log`: This method logs messages either to the console or the log file, based on whether an error has occurred or not.

```go
logger.Write.Log("This is a log message.")
```

2. `Logf`: This method logs formatted messages either to the console or the log file.

```go
logger.Write.Logf("Processed %d items.", 10)
```

3. `Error` and `Errorf`: These methods are used to log error messages. When an error is logged, the subsequent logs will be redirected to the log file.

```go
logger.Write.Error("An error occurred.")
logger.Write.Errorf("Error %s: %v", "code", err)
```

## Example

```go
package main

import (
    "github.com/wbdb/logger"
)

func main() {

    // Some regular log messages
    logger.Write.Log("Starting application...")
    logger.Write.Log("Application initialized.")

    // Simulate an error
    logger.Write.Error("Something went wrong!")

    //Errorf
    logger.Write.Errorf("Error read files %s: %v\n", txtPath, err)

    // Error messages and subsequent logs will be written to the log file
    logger.Write.Log("Continuing after the error.")
}
```
