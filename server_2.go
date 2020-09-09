package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const address = "127.0.0.1:4322"

func main() {

	router := gin.Default()
	router.GET("/:path", func(c *gin.Context) {
		request := c.Request
		url := fmt.Sprintf("http://%s%s", address, request.URL.Path)
		ip := fmt.Sprintf("RemoteAddress=%s, Header=%s", request.RemoteAddr, request.Header)
		c.JSON(200, gin.H{
			"path": url,
			"ip":   ip,
		})
	})
	if err := router.Run(address); err != nil {
		log.Printf("Error: %v", err)
	}

}
