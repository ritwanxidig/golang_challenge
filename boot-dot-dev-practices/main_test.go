package main

// ================================================================= Select Default Case =================================================================
import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func TestSaveBackups(t *testing.T) {
	expectedLogs := []string{
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"All backups saved!",
	}

	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	logChan := make(chan string)
	go saveBackups(snapshotTicker, saveAfter, logChan)
	actualLogs := []string{}
	for actualLog := range logChan {
		fmt.Println(actualLog)
		actualLogs = append(actualLogs, actualLog)
	}

	if !slices.Equal(expectedLogs, actualLogs) {
		t.Errorf(`
---------------------------------
Test Failed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
	} else {
		fmt.Printf(`
---------------------------------
Test Passed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
	}
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true

// ================================================================= Select =================================================================

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func logMessages(chEmails, chSms chan string) {
// 	for {
// 		select {
// 		case sms, ok := <-chSms:
// 			if !ok {
// 				return
// 			}
// 			logSms(sms)
// 		case email, ok := <-chEmails:
// 			if !ok {
// 				return
// 			}
// 			logEmail(email)
// 		}
// 	}
// }

// // don't touch below this line

// func logSms(sms string) {
// 	fmt.Println("SMS:", sms)
// }

// func logEmail(email string) {
// 	fmt.Println("Email:", email)
// }

// func test(sms []string, emails []string) {
// 	fmt.Println("Starting...")

// 	chSms, chEmails := sendToLogger(sms, emails)

// 	logMessages(chEmails, chSms)
// 	fmt.Println("===============================")
// }

// func main() {
// 	rand.Seed(0)
// 	test(
// 		[]string{
// 			"hi friend",
// 			"What's going on?",
// 			"Welcome to the business",
// 			"I'll pay you to be my friend",
// 		},
// 		[]string{
// 			"Will you make your appointment?",
// 			"Let's be friends",
// 			"What are you doing?",
// 			"I can't believe you've done this.",
// 		},
// 	)
// 	test(
// 		[]string{
// 			"this song slaps hard",
// 			"yooo hoooo",
// 			"i'm a big fan",
// 		},
// 		[]string{
// 			"What do you think of this song?",
// 			"I hate this band",
// 			"Can you believe this song?",
// 		},
// 	)
// }

// func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
// 	chSms = make(chan string)
// 	chEmails = make(chan string)
// 	go func() {
// 		for i := 0; i < len(sms) && i < len(emails); i++ {
// 			done := make(chan struct{})
// 			s := sms[i]
// 			e := emails[i]
// 			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
// 			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
// 			go func() {
// 				time.Sleep(t1)
// 				chSms <- s
// 				done <- struct{}{}
// 			}()
// 			go func() {
// 				time.Sleep(t2)
// 				chEmails <- e
// 				done <- struct{}{}
// 			}()
// 			<-done
// 			<-done
// 			time.Sleep(10 * time.Millisecond)
// 		}
// 		close(chSms)
// 		close(chEmails)
// 	}()
// 	return chSms, chEmails
// }

// ================================================================ Ranges =============================================================

// package main

// import (
// 	"fmt"
// 	"slices"
// 	"testing"
// )

// func Test(t *testing.T) {
// 	type testCase struct {
// 		n        int
// 		expected []int
// 	}
// 	tests := []testCase{
// 		{5, []int{0, 1, 1, 2, 3}},
// 		{3, []int{0, 1, 1}},
// 	}
// 	if withSubmit {
// 		tests = append(tests, []testCase{
// 			{0, []int{}},
// 			{1, []int{0}},
// 			{7, []int{0, 1, 1, 2, 3, 5, 8}},
// 		}...)
// 	}

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range tests {
// 		actual := concurrentFib(test.n)
// 		if !slices.Equal(actual, test.expected) {
// 			failCount++
// 			t.Errorf(`
// ---------------------------------
// Test Failed:
//   n:        %v
//   expected: %v
//   actual:   %v
// `, test.n, test.expected, actual)
// 		} else {
// 			passCount++
// 			fmt.Printf(`
// ---------------------------------
// Test Passed:
//   n:        %v
//   expected: %v
//   actual:   %v
// `, test.n, test.expected, actual)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// }

// // withSubmit is set at compile time depending on which button is used to run the tests
// var withSubmit = true
