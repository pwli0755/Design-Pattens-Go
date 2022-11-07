package main

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Printf("Cashier getting money from patient %v", p.name)
}

func (c *cashier) setNext(next department) {
	c.next = next
}
