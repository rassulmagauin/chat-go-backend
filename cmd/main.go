package main

import (
	"log"

	"github.com/rassulmagauin/chat-go/db"
	"github.com/rassulmagauin/chat-go/internal/user"
	"github.com/rassulmagauin/chat-go/internal/ws"
	"github.com/rassulmagauin/chat-go/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")

}
