package pattern

import "fmt"

/*
	Паттерн Стратегия (Strategy)
	Поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
	после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
*/

// PaymentStrategy is the interface that all payment strategies must implement
type PaymentStrategy interface {
	Pay(amount float64) error
}

// CreditCardStrategy is a concrete payment strategy that implements the PaymentStrategy interface
type CreditCardStrategy struct {
	cardNumber string
	cvv        string
}

func (c *CreditCardStrategy) Pay(amount float64) error {
	fmt.Printf("Paying %.2f with credit card %s\n", amount, c.cardNumber)
	return nil
}

// PayPalStrategy is another concrete payment strategy that implements the PaymentStrategy interface
type PayPalStrategy struct {
	email    string
	password string
}

func (p *PayPalStrategy) Pay(amount float64) error {
	fmt.Printf("Paying %.2f with PayPal account %s\n", amount, p.email)
	return nil
}

// PaymentContext is the object that will use the selected payment strategy
type PaymentContext struct {
	paymentStrategy PaymentStrategy
}

func (pc *PaymentContext) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
	pc.paymentStrategy = paymentStrategy
}

func (pc *PaymentContext) MakePayment(amount float64) error {
	return pc.paymentStrategy.Pay(amount)
}

func main() {
	creditCardStrategy := &CreditCardStrategy{cardNumber: "1234567890123456", cvv: "123"}
	payPalStrategy := &PayPalStrategy{email: "example@paypal.com", password: "password"}

	paymentContext := &PaymentContext{}

	paymentContext.SetPaymentStrategy(creditCardStrategy)
	paymentContext.MakePayment(100.0)

	paymentContext.SetPaymentStrategy(payPalStrategy)
	paymentContext.MakePayment(50.0)
}
