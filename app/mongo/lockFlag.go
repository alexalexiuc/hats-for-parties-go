package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func lock() (bool, error) {
	res, err := MongoDbConn.LockFlag.UpdateOne(
		context.Background(),
		bson.M{"isLocked": false},
		bson.M{"$set": bson.M{"isLocked": true}},
	)

	if err != nil {
		return false, err
	}

	return res.ModifiedCount == 1, nil
}

func SetLockFlag() error {
	for {
		locked, err := lock()
		if err != nil {
			return err
		}
		if locked {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func ReleaseLockFlag() error {
	_, err := MongoDbConn.LockFlag.UpdateOne(
		context.Background(),
		bson.M{"isLocked": true},
		bson.M{"$set": bson.M{"isLocked": false}},
	)

	return err
}
