package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rassulmagauin/chat-go/internal/user"
	"github.com/rassulmagauin/chat-go/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomID", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getRoom/:roomID", wsHandler.GetClients)

}

func Start(addr string) error {
	return r.Run(addr)
}
