package printers

import "fmt"

type defaultResponsePrinter struct {
}

func NewDefaultResponsePrinter() *defaultResponsePrinter {
	return &defaultResponsePrinter{}
}

func (p *defaultResponsePrinter) PrintCity(cityFormatted string) {
	fmt.Println(cityFormatted)
}
