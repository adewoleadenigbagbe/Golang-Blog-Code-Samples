package main

import (
	"errors"
)

// Product represents the complex object
type Pizza struct {
	Dough string
	Sauce string
	Size  string
}

type PizzaBuilder struct {
	pizza Pizza
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

func (pb *PizzaBuilder) SetDough(dough string) *PizzaBuilder {
	pb.pizza.Dough = dough
	return pb
}

func (pb *PizzaBuilder) SetSauce(sauce string) *PizzaBuilder {
	pb.pizza.Sauce = sauce
	return pb
}

func (pb *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	pb.pizza.Size = size
	return pb
}

func (pb *PizzaBuilder) Build() (Pizza, error) {
	if pb.pizza.Size == "" {
		return Pizza{}, errors.New("size is required")
	}
	return pb.pizza, nil
}
