package mongo

import (
	"context"
	"errors"
	"fmt"
	"hats-for-parties/cache"
	"hats-for-parties/config"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Hat struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	LastUsage     primitive.DateTime `bson:"lastUsage,omitempty"`
	UsedInPartyId string             `bson:"usedInPartyId,omitempty"`
}

func RentHats(partyId string, hatsNumber int) error {
	if hatsNumber > config.ServiceConfig.TotalHatsPerParty {
		return errors.New("You are not allowed to rent more than " + strconv.Itoa(config.ServiceConfig.TotalHatsPerParty) + " hats")
	}

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.

	cache.SetLockFlag()
	defer cache.ReleaseLockFlag()

	filter := bson.M{
		"usedInPartyId": "",
		"$or": []bson.M{
			{"lastUsage": bson.M{
				"$lt": primitive.NewDateTimeFromTime(time.Now().Add(-time.Duration(config.ServiceConfig.CleaningTimeInSeconds) * time.Second)),
			}},
			{"lastUsage": nil},
		},
	}
	fmt.Printf("Filter: %v\n", filter)

	findOpts := options.Find()
	findOpts.SetLimit(int64(hatsNumber))
	findOpts.SetSort(bson.D{{"lastUsage", 1}})

	cursor, err := MongoDbConn.HatsCollection.Find(context.Background(), filter, findOpts)

	if err != nil {
		return err
	}

	var availableHats []Hat

	if err = cursor.All(context.Background(), &availableHats); err != nil {
		panic(err)
	}

	fmt.Printf("availableHats: %+v\n", availableHats)

	if len(availableHats) < hatsNumber {
		return errors.New("There are not enough hats available, available hats: " + strconv.Itoa(len(availableHats)))
	}

	// For test purposes
	// time.Sleep(5 * time.Second)

	for _, hat := range availableHats {
		hat.UsedInPartyId = partyId
		hat.LastUsage = primitive.NewDateTimeFromTime(time.Now())

		updateResult, err := MongoDbConn.HatsCollection.UpdateOne(
			context.Background(),
			bson.M{"_id": hat.ID},
			bson.D{
				{"$set", bson.D{
					{"usedInPartyId", hat.UsedInPartyId},
				}},
			},
		)

		if err != nil {
			return err
		}

		fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}

	return nil
}

func ReturnHats(partyId string) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	filter := bson.M{"usedInPartyId": partyId}
	update := bson.M{
		"$set": bson.M{
			"usedInPartyId": "",
			"lastUsage":     primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	updateResult, err := MongoDbConn.HatsCollection.UpdateMany(ctx, filter, update)

	if err != nil {
		return err
	}

	if updateResult.MatchedCount == 0 {
		return errors.New("No hats were rented for party " + partyId)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return nil
}
