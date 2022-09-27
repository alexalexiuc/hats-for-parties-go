package config

import (
	"os"
	"strconv"
)

var ServiceConfig struct {
	// Port is the port the server will listen on
	Port                   int
	MongoConnString        string
	RedisConnString        string
	TotalHats              int
	TotalHatsPerParty      int
	CleaningTimeInSeconds  int
	HatsCollectionName     string
	LockFlagCollectionName string
	DBName                 string
}

func getStringEnvVar(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		panic("FATAL Error: " + key + " environment variable not set")
	}

	return value
}

func getIntEnvVar(key string) int {
	value, exists := os.LookupEnv(key)

	if !exists {
		panic("FATAL Error: " + key + " environment variable not set")
	}

	intValue, err := strconv.Atoi(value)

	if err != nil {
		panic("FATAL Error: " + key + " environment variable is not an integer")
	}
	return intValue
}

func init() {
	ServiceConfig.Port = getIntEnvVar("PORT")
	ServiceConfig.MongoConnString = getStringEnvVar("MONGODB_URI")
	ServiceConfig.RedisConnString = getStringEnvVar("REDIS_URI")
	ServiceConfig.TotalHats = getIntEnvVar("TOTAL_HATS")
	ServiceConfig.TotalHatsPerParty = getIntEnvVar("TOTAL_HATS_PER_PARTY")
	ServiceConfig.CleaningTimeInSeconds = getIntEnvVar("CLEANING_TIME_IN_SECONDS")
	ServiceConfig.HatsCollectionName = getStringEnvVar("HATS_COLLECTION_NAME")
	ServiceConfig.LockFlagCollectionName = getStringEnvVar("LOCK_FLAG_COLLECTION_NAME")
	ServiceConfig.DBName = getStringEnvVar("DB_NAME")
}
