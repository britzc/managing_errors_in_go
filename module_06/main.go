package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func main() {
	heading("Type Assertion Output")
	TypeAssertion()

	heading("errors.As Output")
	ErrorsAs()

	heading("Switch Statement Output")
	SwitchStatement()

	heading("Direct Comparison Output")
	DirectComparison()
}

type PaymentError struct {
	Reference string
	Amount    float64
	Message   string
	Timestamp time.Time
}

func NewPaymentError(ref string, amt float64, msg string) *PaymentError {
	return &PaymentError{
		Reference: ref,
		Amount:    amt,
		Message:   msg,
		Timestamp: time.Now(),
	}
}

func (e *PaymentError) Error() string {
	ts := e.Timestamp.Format("2006-01-02 15:04:05")

	return fmt.Sprintf("Payment Error Ref: %s, Amt: %.2f, Msg: %s, Time: %s", e.Reference, e.Amount, e.Message, ts)
}

var ErrInvalidPaymentType = errors.New("Invalid payment reference")

func ProcessPayment(ref string, amt float64) (err error) {
	if ref == "" {
		return ErrInvalidPaymentType
	}

	if amt > 100.0 {
		return NewPaymentError(ref, amt, "Insufficient Funds")
	}

	return nil
}

func TypeAssertion() {
	err := ProcessPayment("ABC", 120.00)

	if err, ok := err.(*PaymentError); ok {
		fmt.Println("Is a payment error")
		fmt.Println(err)
	} else {
		fmt.Println("Is not a payment error")
		fmt.Println(err)
	}
}

func ErrorsAs() {
	err := ProcessPayment("ABC", 120.00)

	var pmtErr *PaymentError
	if errors.As(err, &pmtErr) {
		fmt.Println("Is a payment error")
		fmt.Println(pmtErr)
	} else {
		fmt.Println("Is not a payment error")
		fmt.Println(err)
	}
}

func SwitchStatement() {
	if err := ProcessPayment("ABC", 120.00); err != nil {
		switch e := err.(type) {
		case *PaymentError:
			fmt.Println("Is a payment error")
			fmt.Println(e)
		default:
			fmt.Println("Is not a payment error")
			fmt.Println(e)
		}
	}
}

func DirectComparison() {
	if err := ProcessPayment("", 120.00); err != nil {
		switch err {
		case ErrInvalidPaymentType:
			fmt.Println("Is an invalid payment reference error")
			fmt.Println(err)
		default:
			fmt.Println("Is not an invalid payment reference error")
			fmt.Println(err)
		}
	}
}

func heading(val string) {
	output := fmt.Sprintf("***** %s *****", val)
	line := strings.Repeat("-", len(output))

	fmt.Println(line)
	fmt.Println(output)
}
