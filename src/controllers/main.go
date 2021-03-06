package main

import (
	SC "../conf/server_conf"
	"../models/dial"
	"./auth"
	"./poll"
	"./scrapbook"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"syscall"
	"os"
	"os/signal"
)

func main() {
	
	c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        os.Exit(1)
    }()

	dial.Setup_redis()
	defer dial.Close_redis()

	dial.Setup_sql()
	defer dial.Close_sql()

	// helping-servers set

	router := httprouter.New()
	router.GET("/", auth.Index)
	router.POST("/login", auth.Login)
	router.GET("/paper", poll.Paper)
	router.GET("/thanks", poll.Thank)
	router.POST("/ballot", poll.Vote)
	router.GET("/note", scrapbook.Note)
	router.POST("/book", scrapbook.Paper)
	router.GET("/logout", auth.Logout)
	router.ServeFiles("/res/*filepath", http.Dir(SC.Base_Path+"src/views"))
	log.Fatal(http.ListenAndServe(":"+os.Args[1], router))
}
