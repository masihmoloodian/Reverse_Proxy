package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const reverseServerAddr = "127.0.0.1:4323"

var serverAddr = []string{
	"http://127.0.0.1:4321",
	"http://127.0.0.1:4322"}

var loadBalancerCount int = 0

func loadBalancerAddr() string {
	addr := serverAddr[loadBalancerCount]
	loadBalancerCount++
	if loadBalancerCount == len(serverAddr) {
		loadBalancerCount = 0
	}
	return addr
}

func main() {
	router := gin.Default()
	router.GET("/:path", func(c *gin.Context) {
		request := c.Request
		proxy, err := url.Parse(loadBalancerAddr())
		if err != nil {
			log.Printf("Error: can't parse server address. %v", err)
			c.String(500, "ERROR")
			return
		}
		request.URL.Scheme = proxy.Scheme
		request.URL.Host = proxy.Host

		transport := http.DefaultTransport
		response, err := transport.RoundTrip(request)
		if err != nil {
			log.Printf("RTT Error: %v", response)
			c.String(500, "ERROR")
			return
		}

		for key, value := range response.Header {
			for _, val := range value {
				c.Header(key, val)
			}
		}

		defer response.Body.Close()
		bufio.NewReader(response.Body).WriteTo(c.Writer)
		return
	})

	if err := router.Run(reverseServerAddr); err != nil {
		log.Printf("ERROR: can't run router. %v", err)
	}

}
