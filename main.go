package main

func  main(){
	 // Go is a static type lang, cant change variable type later 
	 cards, _ := newDeckFromFile("mycards.txt")
	 cards.shuffle()
	 // same package
	 cards.print()
}