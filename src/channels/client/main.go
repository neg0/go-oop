package main

import (
	"fmt"
	"github.com/neg0/go-oop/src/channels/credit_account"
	"time"
)

func main() {
	// Standard Channel
	chargeCh := make(chan float32)
	credit_account.CreateCreditAccount(chargeCh, 12000)
	chargeCh <- 500
	chargeCh <- 7800
	time.Sleep(5 * time.Millisecond)

	/// Buffered Channel
	bufferedCh := make(chan float32, 3)
	bufferedCh <- 22.9
	bufferedCh <- 12.42
	bufferedCh <- 5.81
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
}
