package main

import (
	"log"
	"os"
	"strings"

	"github.com/lindgrenj6/qmpv.go/mpv"
)

func main() {
	err := mpv.PlayMediaFile(strings.Join(os.Args[1:], " "))
	if err != nil {
		log.Fatal(err)
	}
}
