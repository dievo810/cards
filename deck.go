package main

// Create a new type of 'deck'
// which is a slice of strings

import (
    "fmt"
    "math/rand"
    "os"
    "strings"
)

type deck []string

func newDeck() deck {
    cards := deck{}

    cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three", "Four"}

    // {"Ace", "King", "Queen", "Jack", "10",
    // "9", "8", "7", "6", "5", "4", "3", "2"}

    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value+" of "+suit)
        }
    }
    return cards
}

func (d deck) print() {
    for i, card := range d {
	fmt.Println(i, card)
    }
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
    // convert type deck into type string
    // cardString := []string(d)
    return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
    return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
    bs, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    s := strings.Split(string(bs), ",")
    return deck(s)
}

func (d deck) shuffleCards() {
    rand.Shuffle(len(d), func(i, j int) {
        d[i], d[j] = d[j], d[i]
	})
}

