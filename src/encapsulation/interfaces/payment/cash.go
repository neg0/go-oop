package payment

import "fmt"

type Cash struct{}

func CreateCash() *Cash {
	return &Cash{}
}

func (ca Cash) ProcessEncapsulatedPayment(amount float32) bool {
	fmt.Println("Processing a cash payment...")
	return true
}
