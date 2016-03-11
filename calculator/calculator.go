package calculator
import "github.com/mateuszdyminski/how-not-to-write-go/logger"

func Add(a, b int) int {
	return a + b



}

func Multiple(a, b int) 	int {
	l, _ := logger.NewLogger("/tmp/test1")

	l.UnusedExt = "test7"

	return a * b
}

func Divide(a, b int) int {
	return a / b
}