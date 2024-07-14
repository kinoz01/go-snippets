package main

import "fmt"

type PizzaOption func(*Pizza)

type Pizza struct {
	Size     string
	Toppings []string
	Crusty   bool
	Price    float64
}

func PizzaSize(size string) PizzaOption {
	return func(p *Pizza) {
		p.Size = size
		switch p.Size {
		case "Small":
			p.Price = 8.99
		case "Medium":
			p.Price = 10.99
		case "Large":
			p.Price = 12.99
		case "king":
			p.Price = 16.99
		}
	}
}

func PizzaToppings(toppings []string) PizzaOption {
	return func(p *Pizza) {
		p.Toppings = toppings
		p.Price += float64(len(toppings)) * 1.5
	}
}

func PizzaCrust(crust bool) PizzaOption {
	return func(p *Pizza) {
		p.Crusty = crust
		if p.Crusty {
			p.Price += 2
		}
	}
}

func NewPizza(options ...PizzaOption) *Pizza {
	pizza := &Pizza{
		Size:     "Medium",
		Toppings: []string{},
		Crusty:   false,
		Price:    10.99,
	}
	for _, opt := range options {
		opt(pizza)
	}
	return pizza
}

func main() {
	stdPizza := NewPizza()
	fmt.Print("Standard Pizza:\n", *stdPizza,"\n")

	customPizza := NewPizza(PizzaSize("king"), PizzaCrust(true), PizzaToppings([]string{"Cheese", "Mustard"}))
	fmt.Print("Custom Pizza:\n", *customPizza,"\n")
}
