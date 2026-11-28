package main

import (
	"fmt"
	"sync"
)

func main() {
	// Poolを作成。Newで新規作成時のコードを実装
	var count int

	pool := sync.Pool{
		New: func() any {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}
	// 追加した要素から受け取れる
	// プールが空だと新規作成
	pool.Put("manually added: 1")
	pool.Put("manually added: 2")
	fmt.Print(pool.Get())
	fmt.Print(pool.Get())
	fmt.Print(pool.Get()) // これは新規作成

}
