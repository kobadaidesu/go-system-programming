package main

import (
	"fmt"
	"iter"
	"time"
)

func ticker(interval time.Duration) iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		for {
			time.Sleep(interval)
			if !yield(time.Now()) {
				break
			}
		}
	}
}

func main() {
	//チャネルをループで利用
	// チャネルに対する range は、データが送られてくるまで待機して、受信できたらループの処理を実行する。
	// さっき書いた<-doneと動き方は一緒っぽい
	for now := range ticker(5 * time.Second) {
		fmt.Printf("now: %v\n", now)
	}
}

