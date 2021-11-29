package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

var SOCK = os.Getenv("HOME") + "/.umpv_socket"

func main() {
	err := PlayMediaFile(strings.Join(os.Args[1:], " "))
	if err != nil {
		log.Fatal(err)
	}
}

func PlayMediaFile(file string) error {
	sock, err := net.Dial("unix", SOCK)

	// if mpv is not running - err will be present, so we need to run mpv
	if err != nil {
		err = RunMpv(file)
		if err != nil {
			return err
		}
	} else {
		err = AddFileToQueue(sock, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func RunMpv(file string) error {
	cmd := exec.Command("mpv", []string{
		"--no-terminal",
		"--on-all-workspaces",
		"--ytdl",
		"--ontop",
		"--no-border",
		"--force-window",
		"--autofit=960x540",
		"--geometry=+20+50",
		fmt.Sprintf("--input-ipc-server=%s", SOCK),
		"--",
		file,
	}...)

	err := cmd.Start()
	if err != nil {
		return err
	}

	// Sleep a bit to wait for MPV to fire up before exiting the program.
	// Otherwise it never starts.
	time.Sleep(1 * time.Second)
	return nil
}

func AddFileToQueue(sock net.Conn, file string) error {
	// send the loadfile command to mpv via the socket we specify, neat!
	_, err := sock.Write([]byte(fmt.Sprintf("raw loadfile %v append\n", file)))
	if err != nil {
		return err
	}

	return nil
}
