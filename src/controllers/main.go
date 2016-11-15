package main

import (
     SC "../conf/server_conf" 
	"github.com/julienschmidt/httprouter"
    "../models/dial" 
    "./auth" 
    "./poll" 
    "./scrapbook" 
    "log"
    "net/http"
)

func main() {
// helping-servers setup
     auth.Login
    dial.Setup_redis();
    defer dial.Close_redis()

    dial.Setup_sql();
    defer dial.Close_sql()

// helping-servers set

    router := httprouter.New()
    router.GET("/", auth.Index)
    router.POST("/login", auth.Login)
    router.GET("/paper", poll.Paper)
    router.POST("/ballot", poll.Vote)
    router.GET("/note", scrapbook.Note)
    router.POST("/book", scrapbook.Paper)
    router.GET("/logout", auth.Logout)
    router.ServeFiles("/res/*filepath", http.Dir(SC.Base_Path+"src/views"))
    log.Fatal(http.ListenAndServe(":8080", router))
}