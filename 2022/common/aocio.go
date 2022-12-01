package aocio

import (
	"log"
	"os"
)

func Readfile(path string) string {
	buf, _ := os.ReadFile(path)
	return string(buf)
}

func TestOutput(name string, expected, real int) {
	if real != expected {
		log.Fatalf("%s should be %d but is %d", name, expected, real)
	} else {
		log.Printf("%s is ok: %d", name, real)
	}
}

func PrintOutput(name string, real int) {
	log.Printf("%s output is: %d", name, real)
}
