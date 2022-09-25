package cache

import (
	"fmt"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

const (
	RedisMutexName = "hats-for-parties-mutex"
)

var redisMutex *redsync.Mutex
var redisClient *goredislib.Client

func InitRedisClient() {
	redisClient = goredislib.NewClient(&goredislib.Options{
		Addr: "cache:6379",
	})
	pool := goredis.NewPool(redisClient) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	redisMutex = rs.NewMutex(RedisMutexName)
}

func CloseRedisClient() {
	redisClient.Close()
}

func SetLockFlag() {
	fmt.Println("Locking")
	if err := redisMutex.Lock(); err != nil {
		panic(err)
	}
	fmt.Println("Locked")
}

func ReleaseLockFlag() {
	fmt.Println("Unlocking")
	if ok, err := redisMutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
	fmt.Println("Unlocked")
}
