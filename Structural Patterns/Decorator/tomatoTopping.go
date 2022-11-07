package main

type tomatotopping struct {
	pizza pizza
}

func (c *tomatotopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
