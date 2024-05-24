package main

import "github.com/gin-gonic/gin"

func main() {
	c := gin.Default()

	c.Run(":9527")
}
