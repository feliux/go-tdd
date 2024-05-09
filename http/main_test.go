package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
// iter 1
func TestGetPlayerScore(t *testing.T) {
	t.Run("return edd score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/edd", nil)
		responseWriter := httptest.NewRecorder()

		PlayerServer(responseWriter, request)
		want := "20"
		got := responseWriter.Body.String()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
*/

/*
// iter 2: happy path and mock
type PlayerStoreMock struct {
	// mock to control player score
	scores map[string]int
}

func (s *PlayerStoreMock) GetScores(name string) int {
	return s.scores[name]
}

func TestGetPlayerScore(t *testing.T) {
	storeMock := &PlayerStoreMock{
		map[string]int{
			"edd":  10,
			"pepe": 20,
		}}
	server := &PlayerServer{store: storeMock}
	t.Run("return edd score", func(t *testing.T) {
		request := newGetScoreRequest("edd")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertResponseBody(t, responseWriter.Body.String(), "10")
	})

	t.Run("return pepe score", func(t *testing.T) {
		request := newGetScoreRequest("pepe")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertResponseBody(t, responseWriter.Body.String(), "20")
	})
}

func newGetScoreRequest(playerName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+playerName, nil)
	return req
}

func assertResponseBody(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
*/

/*
// iter 3: errors (what if a player does not exists)
type PlayerStoreMock struct {
	// mock to control player score
	scores map[string]int
}

func (s *PlayerStoreMock) GetScores(name string) int {
	return s.scores[name]
}

func TestGetPlayerScore(t *testing.T) {
	storeMock := &PlayerStoreMock{
		map[string]int{ // apollo is missing
			"edd":  10,
			"pepe": 20,
		}}
	server := &PlayerServer{store: storeMock}
	t.Run("return edd score", func(t *testing.T) {
		request := newGetScoreRequest("edd")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertStatus(t, responseWriter.Code, http.StatusOK)
		assertResponseBody(t, responseWriter.Body.String(), "10")
	})

	t.Run("return pepe score", func(t *testing.T) {
		request := newGetScoreRequest("pepe")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertStatus(t, responseWriter.Code, http.StatusOK)
		assertResponseBody(t, responseWriter.Body.String(), "20")
	})
	t.Run("404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("apollo")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertStatus(t, responseWriter.Code, http.StatusNotFound)
		assertResponseBody(t, responseWriter.Body.String(), "player not found")
	})
}

func newGetScoreRequest(playerName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+playerName, nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) { // iter 3
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
*/

// iter 4: post endpoint
type PlayerStoreMock struct {
	// mock to control player score
	scores map[string]int
}

func (s *PlayerStoreMock) GetScores(name string) int {
	return s.scores[name]
}

func (s *PlayerStoreMock) RecordWin(name string) {
	s.scores[name]++
}

func TestStoreWins(t *testing.T) {
	storeMock := &PlayerStoreMock{
		map[string]int{ // apollo is missing
			"edd":  10,
			"pepe": 20,
		}}
	server := &PlayerServer{store: storeMock}
	t.Run("records wins when POST", func(t *testing.T) {
		request := newPostWinRequest("edd")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertStatus(t, responseWriter.Code, http.StatusAccepted)
		want := 11
		got := server.store.GetScores("edd")
		if got != want {
			t.Errorf("did not get correct score, got %d , want %d", got, want)
		}
	})
}

func newPostWinRequest(playerName string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/"+playerName, nil)
	return req
}

func TestGetPlayerScore(t *testing.T) {
	storeMock := &PlayerStoreMock{
		map[string]int{ // apollo is missing
			"edd":  10,
			"pepe": 20,
		}}
	server := &PlayerServer{store: storeMock}
	t.Run("return edd score", func(t *testing.T) {
		request := newGetScoreRequest("edd")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertStatus(t, responseWriter.Code, http.StatusOK)
		assertResponseBody(t, responseWriter.Body.String(), "10")
	})

	t.Run("return pepe score", func(t *testing.T) {
		request := newGetScoreRequest("pepe")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertStatus(t, responseWriter.Code, http.StatusOK)
		assertResponseBody(t, responseWriter.Body.String(), "20")
	})
	t.Run("404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("apollo")
		responseWriter := httptest.NewRecorder()
		server.ServeHTTP(responseWriter, request)
		assertStatus(t, responseWriter.Code, http.StatusNotFound)
		assertResponseBody(t, responseWriter.Body.String(), "player not found")
	})
}

func newGetScoreRequest(playerName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+playerName, nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) { // iter 3
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
