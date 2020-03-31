package main

import (
    "flag"
    "fmt"
    "os"
)
 
func main() {



	sayptr := flag.String("say","","Hi for start of command \nBye for terminating the command")
	flag.Parse()

	if len(os.Args) < 2{
		printusage()
	}

	if (*sayptr == "") {
		printusage()
	}

	var command string = os.Args[2]

	if command == "Hi" {
		start()
	} else{
		if command == "Bye" {
			end()
		} else {
			printusage();
		}
	}
}

func printusage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Printf("options:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func start() {
	add()
}

func end() {
	diff()
}