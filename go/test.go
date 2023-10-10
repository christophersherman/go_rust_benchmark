package main

import (
	"github.com/gin-gonic/gin"
)

func fibonacci(n int ) int {
	if n<=1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n+1)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		s := make([]int, 0, 10000)
		for i:=0; i< 10000; i++ {
			s = append(s, i)
		}	
		c.String(200, "Hello World")
	})
	
	r.GET("/fib", func(c *gin.Context) {
			
		c.String(200, "Hello World")
	})


	r.Run() // By default, it serves on :8080 unless a PORT environment variable was defined.
}
