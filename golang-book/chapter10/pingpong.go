package main

import (
	"fmt"
	"time"
	"strconv"
)

func pinger(c chan string){
	for i := 0; ; i++{
		c <- "ping " + strconv.Itoa(i)
	}
}

func printer(c chan string){
	for {
		fmt.Println(<- c)
		time.Sleep(time.Millisecond * 500)
	}
}

func ponger(c chan string){
	for i := 0; ; i++ {
		c <- "pong " + strconv.Itoa(i)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)


	var input string
	fmt.Scanln(&input)
}
