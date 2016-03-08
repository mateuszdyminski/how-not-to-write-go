package logger

import (
	"errors"
	"fmt"
	"os"
)

const DefaultLevel string = "info"

type Logger struct {
	Path  string
	Level string
}

func NewLogger(path string) (*Logger, error) {
	if path == "" {
		return nil, errors.New("Path is required")
	} else {
		return &Logger{path, DefaultLevel}, nil
	}
}

func (l *Logger) Log(data string) {
	file, _ := os.Create(l.Path)
	file.Write([]byte(data))
	file.Close()
}

func (l *Logger) Void() {
	a := 0
	if a != 1 || a != 2 {
		a++
	}

	fmt.Printf("a = %s\n", a)
}
