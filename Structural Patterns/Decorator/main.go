package main

import "fmt"

func main() {
	pizza := &veggieMania{}

	// get cheeseTopping
	pizzaWithCheeseTop := &cheeseTopping{pizza: pizza}

	// get tomatoTopping
	pizzaWithTomatoTop := &tomatotopping{pizza: pizzaWithCheeseTop}

	fmt.Printf("Here is the total count: %d", pizzaWithTomatoTop.getPrice())
}
