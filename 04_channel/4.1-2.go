package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	//inlineで無名関数を作成してその場でgoroutine
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}
