package main

import (
	"fmt"
	"time"
)
//<-chan 受信用のチャネル 
func ticker(interval time.Duration) <-chan time.Time {
	t := make(chan time.Time)
	go func() {
		for {
			time.Sleep(interval)
			t <- time.Now()
		}
	}()
	return t
}

func main() {
	//チャネルをループで利用
	// チャネルに対する range は、データが送られてくるまで待機して、受信できたらループの処理を実行する。
	// さっき書いた<-doneと動き方は一緒っぽい
	for now := range ticker(5 * time.Second) {
		fmt.Printf("now: %v\n", now)
	}
}
