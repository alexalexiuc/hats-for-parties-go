package main

import (
	"hats-for-parties/mongo"
	"hats-for-parties/router"
)

func main() {
	mongo.InitMongoClient()
	defer mongo.CloseMongoClient()

	router.StartRequestHandler()
}
