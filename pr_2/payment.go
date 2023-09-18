package practice_2

import "fmt"

// Strategy Pattern

type PaymentStrategy interface {
	Pay(amount int) error
}

type CreditCardPayment struct {
}

func (c *CreditCardPayment) Pay(amount int) error {
	fmt.Printf("Credit card payment was successful, amount paid: %d₸\n", amount)
	return nil
}

type KaspiQrPayment struct {
}

func (k *KaspiQrPayment) Pay(amount int) error {
	fmt.Printf("Kaspi QR payment was successful, amount paid: %d₸\n", amount)
	return nil
}
