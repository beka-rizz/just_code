package practice_2

import "log"

type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func (sc *ShoppingCart) setPaymentStrategy(strategy PaymentStrategy) {
	sc.paymentStrategy = strategy
}

func (sc *ShoppingCart) Checkout(amount int) {
	err := sc.paymentStrategy.Pay(amount)
	if err != nil {
		log.Println("Payment fail")
	}
}
