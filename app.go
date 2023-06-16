package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var telegramBot *TeleBot
var callCount map[string]int

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/siphook", sipHook)
	mux.HandleFunc("/", indexPageHandler)
	http.ListenAndServe("0.0.0.0:80", mux)
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func sipHook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	remote := r.URL.Query().Get("remote")
	callCount[remote] += 1
	msg := fmt.Sprintf("- Incoming call, number: %v, total calls from this number: %v", remote, callCount[remote])

	log.Println(msg)
	telegramBot.sendMsg(msg)
}

type TeleBot struct {
	Bot  *tele.Bot
	User *tele.User
}

func (t *TeleBot) sendMsg(msg string) error {
	_, err := t.Bot.Send(t.User, msg)
	return err
}

func init() {
	pref := tele.Settings{
		Token:  os.Getenv("TELEBOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	userID, err := strconv.Atoi(os.Getenv("TELEBOT_USER"))
	if err != nil {
		log.Fatal(err)
	}

	callCount = make(map[string]int)

	telegramBot = &TeleBot{
		Bot:  b,
		User: &tele.User{ID: int64(userID)},
	}
}
