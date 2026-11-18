package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	//終了をうけとるための終了関数付きコンテキスト
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		//終了を通知
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
