package main

import (
	//"fmt"
	"os/exec"
	"log"
	"runtime"
)

func user() string {

	if runtime.GOOS == "windows" {
		log.Fatal("Can't execute this on a windows machine")
	}

	out, err := exec.Command("whoami").Output()

	if err != nil {
		log.Fatal(err)
	}

	user := string(out[:len(out)-1])

	return user
}
