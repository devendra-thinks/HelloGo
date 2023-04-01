package main

import "fmt"

type printer interface {
	get() string
}

type goodprinter struct {
}

type badprinter struct {
}

func (goodprinter) get() string {
	return "good"
}

func (badprinter) get() string {
	return "bad"
}

func print(p printer) error {
	fmt.Println("printing .... ", p.get())
	return nil
}
