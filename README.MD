
## How To

### Install
Install docker(https://docs.docker.com/get-docker/) and docker-compose (https://docs.docker.com/compose/install/)

## Requirements
Create a .env file in the root directory with the following variables:
```
PORT=9090 # Port to run the server on
TOTAL_HATS=100 # Total number of available hats
TOTAL_HATS_PER_PARTY=10 # Total number of hats per party
CLEANING_TIME_IN_SECONDS=20 # Time in seconds when hat is available after returned
HATS_COLLECTION_NAME=hats # Name of the collection in the database
DB_NAME=hats-for-parties # Name of the database
```

## Scripts
* start application: `docker compose up --build`
* stop application: `docker compose down`

## Endpoints
* `POST /start-party/{hatsNumber}` - Rent hatsNumber hats for a party (returns party id to be used to return hats)
* `POST end-party/{partyID}` - Return hats for party with ID partyID

## Technical Implementation
There are 2 routes in the application, to start and end the party. The route that starts the party, returns also a 10 characters string with random capitalized letters that is assigned to the party. This string is used to end the party. The route that ends the party, returns the hats to the pool of available hats. The hats are stored in a MongoDB database in a collection called hats in the following format:
```
{
    "_id": "60a1b1f9b9b1b9b1b9b1b9b1",
    "lastUsage": "2022-09-25T12:04:44.007+00:00", // Date when the hat was returned, it is used to calculate when the hat is available again
    "usedInPartyId": "XVLBZGBAIC", // Party ID that the hat is used in, if the hat is available, this field is null
}
```
When container is build first time, in `./mongo` is an initialize script, that creates ${TOTAL_HATS} documents with hat data.


To ensure concurrency a redis mutex technique is used (https://github.com/go-redsync/redsync). When rent hats handler is called, it locks redis, and if Lock() is called again, it will wait until Unlock have been called.
## Requirements
### Hats For Parties

Imagine you have a service that runs parties and rents hats for them. You have the finite number of hats (**N**). When people have a party, they rent hats from you. Multiple parties can occur at the same time.
For example, a birthday party is started, party holder rents (**M**) hats from you. You give them to him. Then other party is started, that requires (**Q**) hats, that you give if you still have them enough. When a party is over, the hats are returned to you. Returned hats requires some time **T** (for cleaning) and then they can be used again.

#### Write a Golang application for it:
1. Application with any kind of input (it can be CLI, or HTTP, etc). Application's commands (endpoints) are: 
    1. Start a party that requires M hats. (Party ID is returned)
    2. Stop a party P
2. Application needs configuration for storing such values as total number of hats, and maximum number of renting hats per party, time needed for reuse (cleaning time), etc.
3. When selecting hats to be rented: new hats have priority upon hats that were already used. You don't give used hats if you have new ones still. When selecting between used hats: older have priority.
4. Use external database (Mongo is preferable, Postgresql/Mysql is OK as well) for storing current running parties and hats rentals
5. Solution should be conflict-free when two or more parties are being started almost at the same time, they must not rent the same hats. Consider, that the running application will be run as a scalable microservice (so multiple binaries of it can run at the same time)
