package main

import (
	"time"
	"strconv"
	"log"
)

func getTime() (int, int, int, float64) {

	Time := time.Now().Format(time.StampNano)

	hr := Time[7:9]
	min := Time[10:12]
	sec := Time[13:15]
	nano := Time[16:]

	hrTime,err := strconv.Atoi(hr)
	
	if err != nil {
		log.Fatal(err)
	}

	minTime,err := strconv.Atoi(min)

	if err != nil {
		log.Fatal(err)
	}

	secTime,err := strconv.Atoi(sec)

	if err != nil {
		log.Fatal(err)
	}

	nanotime,err := strconv.Atoi(nano)

	nanoTime := float64(nanotime)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return hrTime, minTime, secTime, nanoTime
}