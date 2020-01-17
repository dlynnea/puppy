package main

import "github.com/gin-gonic/gin"

// structs are objects/classes
type Puppy struct {
	Name string
	Breed string
	Age int
	Female bool
}

func handlePuppy(context *gin.Context) {
	puppy := getPuppy()
	context.JSON(200, puppy)
}

func getPuppy(name string) Puppy {
	return Puppy{
		Name: "Bixby",
		Breed: "retriever",
	}
}