package main

import (
	"github.com/gin-contrib/static"
        "github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	"math/rand"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.Use(static.Serve("/public", static.LocalFile("./public", true)))

	r.GET("/rint", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"rid": rand.Int(),
		})
	})


	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func (s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	}) 

	r.Run(":5000")
}
