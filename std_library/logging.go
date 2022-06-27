package main

import (
	"log"
	"os"
	"runtime/trace"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func writeLog(messagetype messageType, message string) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//Important, we need to set the output of the log package to a file
	log.SetOutput(file)

	log.Println("Log message")

	switch messagetype {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)

	}
}

func logging() {
	//If we use os.Open instead of OpenFile, the file needs to exist beforehand
	//We are using multiple constants for the second parameter, CREATE will run if the file doesnt exist, APPEND will run if the file already exists, and WRONLY sets
	//These are also known as arithmetic operators or bitwise operators
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//Important, we need to set the output of the log package to a file
	log.SetOutput(file)

	log.Println("Log message")
}

//Good for metrics when developing, not good for prod
func tracing() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("Couldnt create trace file: %v\n", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close trace file: %v\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("Failed to start trace: %v\n", err)
	}

	defer trace.Stop()

	//CALL FUNCTIONS HERE
}

/*
Error Levels:
Information: If you want to confirm something or log a transaction, good for debugging,
Warning: If you need to get users attention
Error (Something very unexpected happened, should probably exit)
Fatal (Close program immediately) Only use when program cannot run anymore
*/
