package practice

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJOb(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10,20,30},
		TaxRate : taxRate,
		// TaxIncludedPrices: InputPrices *(1+taxRate),
	}
}

func Multiple(){
	fmt.Println("This is Second  file")
}