package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
// iter 1
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, _ = fmt.Fprintf(w, "10")
	}
}
*/

/*
// iter 2
type PlayerStore interface {
	GetScores(name string) int
}
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		playerName := strings.TrimPrefix(r.URL.Path, "/players/")
		_, _ = fmt.Fprintf(w, strconv.Itoa(p.store.GetScores(playerName)))
	}
}

func main() {
	// TDD ensures to test functionality without this main
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":8080", server))
}
*/

/*
// iter 3
type PlayerStore interface {
	GetScores(name string) int
}
type PlayerServer struct {
	store PlayerStore
}

const PlayerNotFoundError = "player not found"

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		playerName := strings.TrimPrefix(r.URL.Path, "/players/")
		playerScore := p.store.GetScores(playerName)
		if playerScore == 0 { // if not find the user apollo return 0
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprintf(w, PlayerNotFoundError)
		} else {
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintf(w, strconv.Itoa(playerScore))
		}
	}
}

func main() {
	// TDD ensures to test functionality without this main
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":8080", server))
}
*/

// iter 4
type PlayerStore interface {
	GetScores(name string) int
	RecordWin(name string)
}
type PlayerServer struct {
	store PlayerStore
}

const PlayerNotFoundError = "player not found"

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	if r.Method == http.MethodPost {
		p.store.RecordWin(playerName)
		w.WriteHeader(http.StatusAccepted)
		return
	}
	if r.Method == http.MethodGet {
		playerScore := p.store.GetScores(playerName)
		if playerScore == 0 { // if not find the user apollo return 0
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprintf(w, PlayerNotFoundError)
		} else {
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintf(w, strconv.Itoa(playerScore))
		}
	}
}

/*
func main() {
	// TDD ensures to test functionality without this main
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":8080", server))
}
*/

// iter 5: make a functional server.
// workaround: copy the PlayerStoreMock as InMemoryStore
type InMemoryStore struct {
	scores map[string]int
}

func (s *InMemoryStore) GetScores(name string) int {
	return s.scores[name]
}

func (s *InMemoryStore) RecordWin(name string) {
	s.scores[name]++
}

func main() {
	// TDD ensures to test functionality without this main
	server := &PlayerServer{store: &InMemoryStore{
		scores: map[string]int{
			"player1": 1,
			"player2": 2,
		}}}
	// curl localhost:8080/players/player1 -> 1
	// curl -X POST localhost:8080/players/player1 -> 2
	// curl localhost:8080/players/player1 -> 1
	log.Fatal(http.ListenAndServe(":8080", server))
}
