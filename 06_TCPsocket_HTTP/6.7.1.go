package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn
	var reader *bufio.Reader

	// リトライ用にループで全体を囲う
	for {
		var err error

		// まだコネクションを張ってない / エラーでリトライ
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			reader = bufio.NewReader(conn)
			fmt.Printf("Access: %d\n", current)
		}

		// POSTで文字列を送るリクエストを作成
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}

		// gzip圧縮を受け入れ可能であることを通知
		request.Header.Set("Accept-Encoding", "gzip")

		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		// サーバーからレスポンスを読み込む
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			fmt.Println("Retry")
			conn.Close()
			conn = nil
			continue
		}

		// ヘッダーのみ表示
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		// gzip圧縮されている場合は展開する
		if response.Header.Get("Content-Encoding") == "gzip" {
			gzipReader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(os.Stdout, gzipReader); err != nil {
				panic(err)
			}
			gzipReader.Close()
		} else {
			if _, err := io.Copy(os.Stdout, response.Body); err != nil {
				panic(err)
			}
		}
		response.Body.Close()

		// 全部送信完了していれば終了
		current++
		if current == len(sendMessages) {
			break
		}
	}

	conn.Close()
}
