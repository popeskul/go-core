package pingpong

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	ch     chan string
	scores map[string]int
	wg     *sync.WaitGroup
}

var finalScore = 11

func New() *Game {
	return &Game{
		ch:     make(chan string),
		scores: make(map[string]int),
		wg:     &sync.WaitGroup{},
	}
}

func (g *Game) Start() {
	fmt.Println("Starting game...")

	g.wg.Add(2)

	go g.play("Pasha")
	go g.play("Masha")

	g.ch <- "begin"

	g.wg.Wait()

	fmt.Println("Game finished:", g.scores)
}

func (g *Game) play(name string) {
	defer g.wg.Done()

	for step := range g.ch {
		if g.isWinner() {
			close(g.ch)
			return
		}

		if randomWinner() {
			g.scores[name]++
			fmt.Println(name, "won", g.scores[name])
			time.Sleep(200 * time.Millisecond)
			g.ch <- "stop"
		} else {
			g.ch <- kindOfStep(step)
		}
	}
}

func (g *Game) isWinner() bool {
	for _, score := range g.scores {
		if score == finalScore {
			return true
		}
	}

	return false
}

func randomWinner() bool {
	// return rand of 20% chance to win the game
	return rand.Intn(5) == 0
}

func kindOfStep(step string) (nextStep string) {
	switch step {
	case "begin", "stop", "pong":
		nextStep = "ping"
	case "ping":
		nextStep = "pong"
	}

	return
}
