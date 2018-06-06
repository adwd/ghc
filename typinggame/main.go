package main

import (
	"fmt"
	"time"
)

var (
	words = []string{"foo", "bar"}
	count = 0
)

func main() {
	// タイピングゲームを作ろう
	// - 標準出力に英単語を出す（出すものは自由）
	// - 標準入力から1行受け取る
	// - 制限時間内に何問解けたか表示する
	println("start")

	done := make(chan bool)
	timeout := time.After(10 * time.Second)

	solved := 0
	go runQuiz(done)

OK:
	for {
		select {
		case ok := <-done:
			if ok {
				fmt.Println("OK")
				solved = solved + 1
			} else {
				fmt.Println("NO")
			}
		case <-timeout:
			break OK
		}
	}

	fmt.Printf("score: %d ", solved)
}

func runQuiz(ch chan<- bool) {
	for {
		word := words[count%len(words)]
		fmt.Printf("Enter text: %s\n", word)
		var input string
		fmt.Scanln(&input)
		fmt.Printf("You entered: %s\n", input)
		ch <- input == word
		count = count + 1
	}
}
