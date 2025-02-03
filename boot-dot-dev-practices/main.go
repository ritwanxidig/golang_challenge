package main

// ================================================================= Select Default Case =================================================================

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

// ================================================================ Select =================================================================

// empty

// ================================================================= Ranges =================================================================

// func concurrentFib(n int) []int {
// 	ch := make(chan int)
// 	go fibonacci(n, ch)
// 	var result []int
// 	for v := range ch {
// 		result = append(result, v)
// 	}
// 	return result
// }

// // don't touch below this line

// func fibonacci(n int, ch chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		ch <- x
// 		x, y = y, x+y
// 	}
// 	close(ch)
// }
