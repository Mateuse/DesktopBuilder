package utils

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	// Initialize loggers with different prefixes
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Log is the main logging function that determines the appropriate logging level and format
// based on the provided parameters:
//
// Usage patterns:
// 1. Log(message, nil) -> Info log without formatting
// 2. Log(message, nil, arg1, arg2, ...) -> Info log with formatting
// 3. Log(message, error) -> Error log without formatting
// 4. Log(message, error, arg1, arg2, ...) -> Error log with formatting
func Log(message string, err error, args ...interface{}) {
	if err != nil {
		// Error logging
		if len(args) > 0 {
			// Format the message first, then add error
			formattedMessage := fmt.Sprintf(message, args...)
			errorLogger.Printf("%s: %v", formattedMessage, err)
		} else {
			// Simple error message
			errorLogger.Printf("%s: %v", message, err)
		}
	} else {
		// Info logging
		if len(args) > 0 {
			// Formatted info message
			infoLogger.Printf(message, args...)
		} else {
			// Simple info message
			infoLogger.Print(message)
		}
	}
}

// Legacy functions kept for backward compatibility (marked as deprecated)
// These should not be used directly - use Log() instead

// LogInfo logs an informational message
// Deprecated: Use Log(message, nil, args...) instead
func LogInfo(message string, args ...interface{}) {
	Log(message, nil, args...)
}

// LogError logs an error message with optional error details
// Deprecated: Use Log(message, err, args...) instead
func LogError(message string, err error, args ...interface{}) {
	Log(message, err, args...)
}

// LogInfof is a convenience function for formatted info logging
// Deprecated: Use Log(message, nil, args...) instead
func LogInfof(format string, args ...interface{}) {
	Log(format, nil, args...)
}

// LogErrorf is a convenience function for formatted error logging
// Deprecated: Use Log(message, err) instead for simple errors, or Log(message, err, args...) for formatted
func LogErrorf(format string, args ...interface{}) {
	// This function is ambiguous without an error parameter, so we'll create a generic error
	genericErr := fmt.Errorf("logged error")
	Log(format, genericErr, args...)
}
