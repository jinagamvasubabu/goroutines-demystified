package main

import (
	"fmt"
	"time"
)

func main() {
	// This below code doesn't run concurrently and never executes the 1 sec and 2 sec functions at all
	/*runFor500Ms()
	runFor1Sec()
	runFor2Secs()*/

	//This below function run concurrently but there is no synchorization and you cannot guarantee the order of execution
	go runFor500Ms()
	go runFor1Sec()
	go runFor2Secs()

	time.Sleep(4 * time.Minute)
}

func runFor500Ms() {
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Ran for 500 MS")
	}
}

func runFor1Sec() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Ran for 1 Sec")
	}
}

func runFor2Secs() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Ran for 2 Sec")
	}
}
