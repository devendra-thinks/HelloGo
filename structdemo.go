package main

type employee struct {
	name    string
	company string
}

func (e *employee) changeCompany(newComp string) {
	(*e).company = newComp
}
