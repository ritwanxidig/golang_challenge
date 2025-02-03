package main

func countReports(numSentCh chan int) int {
	var sentReports int = 0
	for {
		channs, ok := <-numSentCh
		if !ok {
			break
		}
		sentReports += channs
	}
	return sentReports
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
