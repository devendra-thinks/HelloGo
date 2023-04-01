package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// creates a new deck
func newDeck() deck {
	cards := deck{}
	suites := []string{"Chidi", "einth", "paan", "hukum"}
	values := []string{"ekka", "dukki", "tikki", "chauki", "panji", "chakki", "satti", "atthi", "nehli",
		"dehla", "j", "queen", "badshah"}
	for _, suit := range suites {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print a deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal a deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert a deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save to a file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// get from a file
func newDeckFromFile(fileName string) (deck, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error ", err)
		os.Exit(1)
	}

	s := string(bs)
	cards := strings.Split(s, ",")

	return deck(cards), nil
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixMilli())
	r := rand.New(source)
	for i := range d {
		ri := r.Intn(len(d) - 1)
		d[ri], d[i] = d[i], d[ri]
	}
}
