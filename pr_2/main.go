package practice_2

func Test() {
	cart := ShoppingCart{}

	cart.setPaymentStrategy(&CreditCardPayment{})
	cart.Checkout(10000)

	cart.setPaymentStrategy(&KaspiQrPayment{})
	cart.Checkout(15000)
}
