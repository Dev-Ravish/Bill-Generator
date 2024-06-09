package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(prompt string, r *bufio.Reader) (string, error) {

	fmt.Print(prompt)

	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}

func creatBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := input("Create bill in the name of: ", reader)

	myBill := newbill(name)

	return myBill
}

func selectOption(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := input("Select any option (s-save, a-add item, t-add tip): ", reader)

	switch option {
	case "a":
		item, _ := input("Item's name : ", reader)
		price, _ := input("Item's price : ", reader)
		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price should only be a number")
			selectOption(b)
		}

		b.addItem(item, p)

		fmt.Println("Item added - ", item, price)
		selectOption(b)

	case "t":
		tip, _ := input("Tip amount : ", reader)

		tipInFloat, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip should only be a number")
			selectOption(b)
		}

		b.addTip(tipInFloat)

		fmt.Println("Tip added - ", tipInFloat)
		selectOption(b)

	case "s":
		b.saveBill()
	default:
		fmt.Println("Sorry the inout you gave is invalid")
		selectOption(b)
	}

}

func main() {
	myBill := creatBill()
	selectOption(&myBill)
}
