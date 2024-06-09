package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newbill(name string) bill {

	nBill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return nBill
}

func (b *bill) format() string {

	var total float64 = b.tip

	fs := fmt.Sprintf("%v's Bill Breakdown: \n", b.name)
	for key, value := range b.items {

		fs += fmt.Sprintf("%-25v ...Rs %v\n", key+":", value)
		total += value
	}
	fs += fmt.Sprintf("%-25v ...Rs %v\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...Rs %.3f\n", "total:", total)

	return fs
}

func (b *bill) addTip(tip float64) {
	b.tip += tip

}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) saveBill() {
	billData := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", billData, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("The bill have been saved in the name of - ", b.name)
}
