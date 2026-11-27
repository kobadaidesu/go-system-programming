package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		// 非同期で仕事をする (1)
		fmt.Println("仕事１")
	})
	wg.Go(func() {
		// 非同期で仕事をする (2)
		fmt.Println("仕事２")
	})
	// すべての処理が終わるのを待つ
	wg.Wait()
	fmt.Println("終了")
}
