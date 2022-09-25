package main

import (
	"hats-for-parties/cache"
	"hats-for-parties/mongo"
	"hats-for-parties/router"
)

func main() {
	mongo.InitMongoClient()
	defer mongo.CloseMongoClient()

	cache.InitRedisClient()
	defer cache.CloseRedisClient()

	router.StartRequestHandler()
}
