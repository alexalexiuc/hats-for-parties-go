
## How To

### Install
Install docker(https://docs.docker.com/get-docker/) and docker-compose (https://docs.docker.com/compose/install/)

## Requirements
Create a .env file in the root directory with the following variables:
```
PORT=9090
```

## Scripts
* start application: `docker compose up -d --build`
* stop application: `docker compose down`


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
