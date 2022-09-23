package handlers

import (
	"fmt"
	"hats-for-parties/mongo"
	"hats-for-parties/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handlers struct{}

func (h Handlers) HealthCheckHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "pong")
}

func (h Handlers) RouteNotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
}

func (h Handlers) StartParty(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hatsNumber, err := strconv.Atoi(vars["hatsNumber"])
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	partyId := utils.RandString(10)

	err = mongo.RentHats(partyId, hatsNumber)

	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("You rented successfully " + strconv.Itoa(hatsNumber) + " hats for the party\nPlease use this code " + partyId + " it when returning the hats"))
	}
}

func (h Handlers) EndParty(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partyId := vars["partyId"]

	err := mongo.ReturnHats(partyId)

	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("You returned the hats successfully"))
	}
}

var RouteHandlers = Handlers{}
