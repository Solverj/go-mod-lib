package logger

import (
	"encoding/json"
	"fmt"
)

type Severity string

const (
	Info Severity = "INFO"
)

type Entry struct {
	Message  string   `json:"message"`
	Severity Severity `json:"severity"`
}

func (e Entry) string() (string, error) {
	if e.Severity == "" {
		e.Severity = Info
	}
	if e.Message == "" {
		e.Message = "logger: No message provided"
	}
	out, jsonMarshalErr := json.Marshal(e)
	if jsonMarshalErr != nil {
		return "", jsonMarshalErr
	}
	return string(out), nil
}

func INFO(message string) {
	println(message, Info)
}

func println(message string, severity Severity) {
	entry := Entry{Message: message, Severity: severity}
	entryAsString, entryErr := entry.string()
	if entryErr != nil {
		fmt.Printf("Error: %v\n", entryErr)
	}
	fmt.Println(entryAsString)
}
