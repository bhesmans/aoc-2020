package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type player struct {
	cards []int
}

type game struct {
	players []player
	played  map[string]bool
}

var games map[string]int

func newGame() game {
	return game{played: make(map[string]bool)}
}

func readInput(inputName string) game {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	g := newGame()

	scanner := bufio.NewScanner(file)
	var p player
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			g.players = append(g.players, p)
		} else if strings.Contains(l, "Player") {
			p = player{}
		} else {
			v, _ := strconv.Atoi(l)
			p.cards = append(p.cards, v)
		}
	}
	g.players = append(g.players, p)

	err = scanner.Err()
	panicOnError(err)

	return g
}

func (p *player) getCard() int {
	ret := p.cards[0]
	p.cards = p.cards[1:]
	return ret
}

func (p *player) appendCard(c []int) {
	p.cards = append(p.cards, c...)
}

func (p *player) noCard() bool {
	return len(p.cards) == 0
}

func (g *game) winner() player {
	if g.players[0].noCard() {
		return g.players[1]
	} else {
		return g.players[0]
	}
}

func (g *game) round() bool {
	card0, card1 := g.players[0].getCard(), g.players[1].getCard()
	if card0 > card1 {
		g.players[0].appendCard([]int{card0, card1})
	} else {
		g.players[1].appendCard([]int{card1, card0})
	}

	return !(g.players[0].noCard() || g.players[1].noCard())
}

func (g *game) score() int {
	w := g.winner()

	sum := 0
	i := len(w.cards)
	for _, c := range w.cards {
		sum += (c * i)
		i--
	}

	return sum
}

func (p *player) atLeast(n int) bool {
	return len(p.cards) >= n
}

func (p player) copyPlayer(i int) player {
	ret := player{}
	ret.cards = make([]int, i)
	copy(ret.cards, p.cards[:i])
	return ret
}

func (g *game) round2() (bool, int) {
	gstring := fmt.Sprintf("%v", g.players)
	if g.played[gstring] {
		return false, 0
	}

	roundWinner := -1
	card0, card1 := g.players[0].getCard(), g.players[1].getCard()
	if g.players[0].atLeast(card0) && g.players[1].atLeast(card1) {
		g2 := newGame()
		g2.players = append(g2.players, g.players[0].copyPlayer(card0))
		g2.players = append(g2.players, g.players[1].copyPlayer(card1))
		roundWinner = g2.play2()
	} else if card0 > card1 {
		roundWinner = 0
	} else {
		roundWinner = 1
	}

	if roundWinner == 0 {
		g.players[0].appendCard([]int{card0, card1})
	} else if roundWinner == 1 {
		g.players[1].appendCard([]int{card1, card0})
	} else {
		panic("Haaaa")
	}

	g.played[gstring] = true
	return !(g.players[0].noCard() || g.players[1].noCard()), roundWinner
}

func (g *game) play2() int {
	winner := -1
	nextRound := true
	for nextRound {
		nextRound, winner = g.round2()
	}
	return winner
}

func part1(g game) int {
	for g.round() {
	}
	return g.score()
}

func part2(g game) int {
	g.play2()
	return g.score()
}

func initGames() {
	games = make(map[string]int)
}

func main() {
	fmt.Println("Hello")
	initGames()
	// g := readInput("small_input.txt")
	g := readInput("input.txt")
	fmt.Printf("%#v\n", part1(g))
	g = readInput("input.txt")
	fmt.Printf("%v\n", part2(g))
}
