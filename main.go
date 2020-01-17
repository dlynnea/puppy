package main

import "github.com/gin-gonic/gin"

func main() {
	// := the go funcky sytnax for creating a new variable
	// capital D means it's exported (public method)
	// only inside a function
	// go likes single letter variables
	router := gin.Default()
	router.GET("/puppy", handlePuppy)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}