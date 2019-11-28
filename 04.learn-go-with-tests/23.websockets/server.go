package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/websocket"
	"strconv"
	"io/ioutil"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
	game Game
}

const JsonContentType = "application/json"

const HtmlTemplatePath = "game.html"

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles(HtmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem loading template %s %v", HtmlTemplatePath, err)
	}

	p.game = game
	p.template = tmpl
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.playGame))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))
	p.Handler = router

	return p, nil
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) playGame(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

var wsUpgrader = websocket.Upgrader {
	ReadBufferSize : 1024,
	WriteBufferSize : 1024,
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws := NewPlayerServer(w, r)

	_, numberOfPlayerMsg, _ := ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayerMsg))
	p.game.Start(numberOfPlayers, ioutil.Discard) //todo: Don't discard the blinds messages!

	_, winner, _ := ws.WaitForMsg()
	p.game.Finish(string(winner))
}
