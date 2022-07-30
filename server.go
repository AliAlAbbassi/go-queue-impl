package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	q := NewQueue()

	r.POST("/json/:timeInQueue", func(c *gin.Context) {
		q.emptyQueue()
		timeInQueue := c.Param("timeInQueue")
		var json Message
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tiq, err := strconv.Atoi(timeInQueue)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		node := newNode(QueueMessage{
			time:    makeTimestamp() + int64(tiq),
			message: json,
		})

		q.enqueue(node)
		c.String(200, fmt.Sprintf("time in queue will be %v", tiq))
	})

	r.GET("/status", func(ctx *gin.Context) {
		q.emptyQueue()
		ctx.String(200, strconv.Itoa(q.length))
	})

	r.Run("0.0.0.0:8080")
}
