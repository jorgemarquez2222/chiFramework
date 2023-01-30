package factorymethod

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {
}

type Cash struct {
}

type TDC struct {
}

func (p *Paypal) Pay() {
	fmt.Println("Pagando con Paypal")
}
func (p *Cash) Pay() {
	fmt.Println("Pagando con Efectivo")
}
func (p *TDC) Pay() {
	fmt.Println("Pagando con Tarjeta de Crédito")
}

func GetMethod(method int) PayMethod {
	switch method {
	case 1:
		return &Paypal{}
	case 2:
		return &Cash{}
	case 3:
		return &TDC{}
	default:
		return &Cash{}
	}
}
