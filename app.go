package main

import (
	"log"      // This package implements a simple logging package.
	"net/http" // This package provides HTTP client and server implementations.
	"os"
	"fmt"
	"strconv"
	"time"
	tele "gopkg.in/telebot.v3" // Telegram bot API implementation
)

var telegramBot *TeleBot

// The main function is the entry point of the program.
func main(){
	// Create a new ServeMux to handle incoming HTTP requests.
	// ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
	mux := http.NewServeMux()

	// Assign the sipHook function to handle requests to the "/siphook" endpoint.
	// This endpoint is specifically designed to process webhook events from Grandstream SIP phones.
	mux.HandleFunc("/siphook", sipHook)
	
	// Assign the indexPageHandler function to handle requests to the "/" endpoint.
	// This endpoint is a catch-all for any requests to endpoints that don't exist, and it will respond with a 404 status.
	mux.HandleFunc("/", indexPageHandler)

	// Start the HTTP server on port 80 and pass the ServeMux as the main handler.
	// ListenAndServe listens on the TCP network address and then calls Serve with handler to handle requests on incoming connections.
	http.ListenAndServe("0.0.0.0:80", mux)
}

// Function to handle requests to the "/" endpoint.
// This function is a catch-all for any requests to endpoints that don't exist. It responds with a 404 status to indicate that the requested resource could not be found on the server.
func indexPageHandler(w http.ResponseWriter , r *http.Request) {
	w.WriteHeader(http.StatusNotFound) // Respond with a 404 status
}

// Function to handle requests to the "/siphook" endpoint.
// This function processes webhook events from Grandstream SIP phones. It logs the incoming call number from the "remote" query parameter in the request URL.
func sipHook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Respond with a 200 status to indicate that the request has succeeded.

	// Get the "remote" query parameter from the request URL.
	// This parameter contains the number of the incoming call.
	remote := r.URL.Query().Get("remote")

	// Log the incoming call number.
	// This information can be useful for debugging and monitoring purposes.
	msg := fmt.Sprintf("- Incoming call, number: %v", remote)
	log.Println(msg)

	//Send the incoming call to the admin user in telegram
	telegramBot.sendMsg(msg)
}

type TeleBot struct {
	Bot *tele.Bot
	User *tele.User
}

//Sends message to the designated bot admin user
func (t *TeleBot)sendMsg(msg string) error{
	// Send a message to the user
	_, err := t.Bot.Send(t.User, msg)
	if err != nil {
		return err
	}
	return nil
}

func init(){
	pref := tele.Settings{
		Token: os.Getenv("TELEBOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil{
		log.Fatal(err)
	}
	userID, err := strconv.Atoi(os.Getenv("TELEBOT_USER"))
	if err != nil{
		log.Fatal(err)
	}
	
	telegramBot = &TeleBot{
		Bot: b,
		User: &tele.User{ID: int64(userID)} ,
	}
}