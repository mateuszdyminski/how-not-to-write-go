package main

import "github.com/mateuszdyminski/how-not-to-write-go/logger"

func main() {
	l, _ := logger.NewLogger("/tmp/test")
	l.Log("Test log")
}
